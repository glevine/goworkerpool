package main

import (
	"context"
	"fmt"
	"sync"
)

type daemon struct {
	producer *producer
	pool     *workerPool
}

func NewDaemon() *daemon {
	return &daemon{}
}

func (d *daemon) Start(ctx context.Context, producer *producer, pool *workerPool) {
	fmt.Println("daemon: starting")
	defer fmt.Println("daemon: shutdown")

	d.producer = producer
	d.pool = pool

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		d.pool.StartWorkers(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		d.producer.Start(ctx)
	}()

	wg.Wait()
}