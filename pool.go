package main

import (
	"context"
	"fmt"
	"sync"
)

type workerPool struct {
	size  int
	queue *queue
}

func NewWorkerPool(size int, q *queue) *workerPool {
	return &workerPool{
		size:  size,
		queue: q,
	}
}

func (wp *workerPool) StartWorkers(ctx context.Context) {
	fmt.Println("worker pool: starting")

	var wg sync.WaitGroup

	for i := 0; i < wp.size; i++ {
		wg.Add(1)
		go NewWorker(wp.queue).Start(ctx, &wg)
	}

	wg.Wait()

	fmt.Println("worker pool: graceful shutdown")
}