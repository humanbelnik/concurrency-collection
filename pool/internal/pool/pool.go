package pool

import (
	"sync"

	"github.com/humanbelnik/q/internal/executor"
	"github.com/humanbelnik/q/internal/model"
)

type Executor interface {
	Execute(job *model.Job) error
	ID() int
}

type Result struct {
	J *model.Job
	R error
}

type Pool struct {
	executors chan Executor
	jobs      chan *model.Job

	// Used to signal that execution must be stopped.
	// close(done) in order to unlock reading from it in a select.
	done chan struct{}

	// Since Execute(job) is an async operation, we're unable to syncronically return.
	results chan Result

	wg sync.WaitGroup
}

func New(nexecutors int) *Pool {
	executors := make(chan Executor, nexecutors)

	for i := range nexecutors {
		executors <- executor.New(i)
	}
	return &Pool{
		executors: executors,
		jobs:      make(chan *model.Job, 100),
		done:      make(chan struct{}),
		results:   make(chan Result, 1),
	}
}

func (p *Pool) Go() {
	go func() {
		for {
			select {
			case j := <-p.jobs:
				// Job must be executed so that operation is done in a blocking way:
				// wait untill some executor is not ready.
				p.execute(<-p.executors, j)
			case <-p.done:
				// Gracefully execute all remaining jobs
				for j := range p.jobs {
					p.execute(<-p.executors, j)
				}
				return
			}
		}
	}()
}

func (p *Pool) execute(exe Executor, j *model.Job) {
	go func(exe Executor, j *model.Job) {
		defer func() {
			p.executors <- exe
			p.wg.Done()
		}()
		p.results <- Result{j, exe.Execute(j)}
	}(exe, j)
}

func (p *Pool) Submit(j *model.Job) {
	select {
	case <-p.done:
		return
	default:
		p.jobs <- j
		p.wg.Add(1)

	}
}

func (p *Pool) Result() <-chan Result {
	return p.results
}

func (p *Pool) Stop() {
	close(p.done)
	close(p.jobs)
	p.wg.Wait()

	close(p.executors)
	close(p.results)
}
