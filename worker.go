package main

import (
	"context"
	"fmt"
)

type worker struct {
	queue *queue
}

func NewWorker(q *queue) *worker {
	return &worker{
		queue: q,
	}
}

func (w *worker) Start(ctx context.Context) {
	fmt.Println("worker: starting")
	defer fmt.Println("worker: shutdown")

	for {
		select {
		case <- ctx.Done():
			return
		case i := <-w.queue.Receive():
			fmt.Println("worker: ", i)
		}
	}
}