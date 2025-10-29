package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func slow(x int) (int, error) {
	if x == 3 {
		return 777, nil
	}
	return 0, errors.New("oops")
}

func wrapWithTimeoutAndRetry(ctx context.Context, x int) (int, error) {
	type res struct {
		v   int
		err error
	}
	resCh := make(chan res)
	doneCh := make(chan struct{})
	defer func() {
		close(doneCh)
		close(resCh)
	}()

	for {
		go func() {
			v, err := slow(x)
			resCh <- res{v: v, err: err}
			select {
			case <-doneCh:
			default:
				select {
				case <-doneCh:
				case resCh <- res{v: v, err: err}:
				}
			}
		}()

		select {
		// Timeout happens or...
		case <-ctx.Done():
			return 0, ctx.Err()
		// Function is done.
		case res := <-resCh:
			v, err := res.v, res.err
			if err != nil {
				select {
				case <-ctx.Done():
					return 0, ctx.Err()
				default:
					x++
					fmt.Println("error happened but I have time! retrying...")
				}
			} else {
				return v, nil
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
	defer cancel()

	v, err := wrapWithTimeoutAndRetry(ctx, 0)
	fmt.Println(v, err)

}
