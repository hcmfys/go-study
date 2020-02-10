package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done finish")
				return
			default:
				fmt.Println("gorouting run ...")
				time.Sleep(2 * time.Second)
			}

		}
	}(ctx)

	time.Sleep(20 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}
