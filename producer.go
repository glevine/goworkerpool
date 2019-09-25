package main

import (
	"fmt"
)

type producer struct {
	queue *queue
	quit  chan struct{}
}

func NewProducer(q *queue) *producer {
	return &producer{
		queue: q,
		quit:  make(chan struct{}, 1),
	}
}

func (p *producer) Start() {
	fmt.Println("producer: starting")
	defer p.queue.Close()
	defer fmt.Println("producer: shutdown")

	counter := 0
	for {
		select {
		case <-p.quit:
			return
		default:
			counter++
			fmt.Println("producer: ", counter)
			p.queue.Send() <- counter
		}
	}
}

func (p *producer) Stop() {
	p.quit <- struct{}{}
}
