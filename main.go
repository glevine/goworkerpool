package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("main: starting")

	d := NewDaemon()
	q := NewQueue()
	prod := NewProducer(q)
	pool := NewWorkerPool(5, q)

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
		d.Start(prod, pool)
	}()

	time.Sleep(time.Second * 3)
}
