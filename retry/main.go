package main

import (
	"context"
	"errors"
	"fmt"
)

type Executor func(ctx context.Context, v *int) error

func worker(ctx context.Context, v *int) error {
	if *v == 0 {
		return nil
	}
	(*v)--
	return errors.New("oops")
}

func retry(e Executor) Executor {
	

	
}
