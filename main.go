package main

import (
	"fmt"
	"context"
	"sync"
	"time"
)

func main() {
	fmt.Println("main: starting")

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	defer func() {
		fmt.Println("main: terminating")
		cancel()
		wg.Wait()
		fmt.Println("main: shutdown")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		d := NewDaemon()
		q := NewQueue()
		prod := NewProducer(q)
		pool := NewWorkerPool(5, q)
		d.Start(ctx, prod, pool)
	}()

	time.Sleep(time.Second * 15)
}
