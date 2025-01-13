package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	go foo(ctx)
	time.Sleep(time.Second * 5)
}

func foo(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("done")
	}
}
