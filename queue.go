package main

type queue struct {
	work chan int
}

func NewQueue() *queue {
	return &queue{
		work: make(chan int),
	}
}

func (q *queue) Send() chan<- int {
	return q.work
}

func (q *queue) Receive() <-chan int {
	return q.work
}
