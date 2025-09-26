package scheduler

import (
	"context"
	"fmt"
	"sync"

	"github.com/humanbelnik/q/internal/model"
	"github.com/humanbelnik/q/internal/pool"
)

type Scheduler struct {
	wg   sync.WaitGroup
	pool *pool.Pool
}

func New(p *pool.Pool) *Scheduler {
	return &Scheduler{
		pool: p,
	}
}

func (s *Scheduler) Go(ctx context.Context, source <-chan *model.Job) {
	s.pool.Go()
	s.wg.Add(1)
	s.handleResults()
	for {
		select {
		case <-ctx.Done():
			// All source jobs must be executed
			for job := range source {
				s.pool.Submit(job)
			}
			s.pool.Stop()
			s.wg.Wait()
			return
		case job := <-source:
			s.pool.Submit(job)
		}
	}
}

func (s *Scheduler) handleResults() {
	counter := 0
	go func() {
		defer s.wg.Done()
		for range s.pool.Result() {
			counter++
		}
		fmt.Printf("Made %d executions\n", counter)
	}()
}
