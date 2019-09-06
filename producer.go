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

	for {
		select {
		case <- ctx.Done():
			fmt.Println("producer: graceful shutdown")
			return
		default:
			p.queue.Send() <- 2
			time.Sleep(time.Second)
		}
	}
}