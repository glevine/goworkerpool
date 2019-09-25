package main

import (
	"fmt"
	"time"
)

type worker struct {
	queue *queue
}

func newWorker(q *queue) *worker {
	return &worker{
		queue: q,
	}
}

func (w *worker) start() {
	fmt.Println("worker: starting")
	defer fmt.Println("worker: shutdown")

	for i := range w.queue.Receive() {
		time.Sleep(time.Second)
		fmt.Println("worker: ", i)
	}
}
