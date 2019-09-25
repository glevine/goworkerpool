package main

import (
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

func (wp *workerPool) start() {
	fmt.Println("worker pool: starting")
	defer fmt.Println("worker pool: shutdown")

	var wg sync.WaitGroup

	for i := 0; i < wp.size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newWorker(wp.queue).start()
		}()
	}

	wg.Wait()
}
