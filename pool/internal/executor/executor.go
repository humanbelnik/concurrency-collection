package executor

import (
	"time"

	"github.com/humanbelnik/q/internal/model"
)

type Executor struct {
	id int
}

func New(id int) *Executor {
	return &Executor{id: id}
}

func (e *Executor) ID() int {
	return e.id
}

func (e *Executor) Execute(j *model.Job) error {
	time.Sleep(1 * time.Second)
	return nil
}
