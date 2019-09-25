package main

import (
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

func (d *daemon) Start(producer *producer, pool *workerPool) {
	fmt.Println("daemon: starting")
	defer fmt.Println("daemon: shutdown")

	d.producer = producer
	d.pool = pool

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		d.pool.start()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		d.producer.start()
	}()

	wg.Wait()
}

func (d *daemon) Stop() {
	d.producer.stop()
}
