package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker is cancelled")
			return
		default:
			fmt.Println("worker is running")
			time.Sleep(1 * time.Second)
		}
	}
}

// ref: https://ithelp.ithome.com.tw/articles/10303330?sc=iThelpR
// ref: https://pjchender.dev/golang/pkg-context
// ref:

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("main is cancelled")
}
