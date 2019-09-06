package main

import (
	"context"
	"fmt"
	"sync"
)

type worker struct {
	queue *queue
}

func NewWorker(q *queue) *worker {
	return &worker{
		queue: q,
	}
}

func (w *worker) Start(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("worker: starting")
	defer wg.Done()

	for {
		select {
		case <- ctx.Done():
			fmt.Println("worker: graceful shutdown")
			return
		case i := <-w.queue.Receive():
			fmt.Println("worker: ", i)
		}
	}
}