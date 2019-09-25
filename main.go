package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("main: starting")

	d := NewDaemon()

	var wg sync.WaitGroup
	defer func() {
		fmt.Println("main: terminating")
		d.Stop()
		wg.Wait()
		fmt.Println("main: shutdown")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		q := NewQueue()
		prod := NewProducer(q)
		pool := NewWorkerPool(5, q)
		d.Start(prod, pool)
	}()

	time.Sleep(time.Second * 3)
}
