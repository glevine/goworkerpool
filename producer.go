package main

import (
	"fmt"
)

type producer struct {
	queue *queue
	quit  bool
}

func NewProducer(q *queue) *producer {
	return &producer{
		queue: q,
		quit:  false,
	}
}

func (p *producer) start() {
	fmt.Println("producer: starting")
	defer p.queue.close()
	defer fmt.Println("producer: shutdown")

	counter := 0
	for !p.quit {
		counter++
		fmt.Println("producer: ", counter)
		p.queue.send() <- counter
	}
}

func (p *producer) stop() {
	p.quit = true
}
