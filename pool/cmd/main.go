package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/humanbelnik/q/internal/model"
	"github.com/humanbelnik/q/internal/pool"
	"github.com/humanbelnik/q/internal/scheduler"
)

func main() {
	pool := pool.New(10)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	sched := scheduler.New(pool)
	wg := sync.WaitGroup{}
	wg.Go(func() {
		sched.Go(ctx, func(ctx context.Context) <-chan *model.Job {
			source := make(chan *model.Job, 500)
			id := 0
			go func() {
				defer func() {
					close(source)
					fmt.Printf("Send %d jobs to a source\n", id)
				}()
				for {
					select {
					case <-ctx.Done():
						fmt.Printf("At the moment of context cancelling %d jobs were sent\n", id)
						return
					default:
						job := &model.Job{ID: id}
						source <- job
						id++
					}
				}
			}()
			return source
		}(ctx),
		)
	})
	wg.Wait()
}
