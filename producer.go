package main

import (
	"context"
	"fmt"
	"time"
)

type producer struct {
	queue *queue
}

func NewProducer(q *queue) *producer {
	return &producer{
		queue: q,
	}
}

func (p *producer) Start(ctx context.Context) {
	fmt.Println("producer: starting")
	defer fmt.Println("producer: shutdown")

	for {
		select {
		case <- ctx.Done():
			return
		case p.queue.Send() <- 2:
			time.Sleep(time.Second)
		default:
			fmt.Println("producer: unable to send work")
		}
	}
}