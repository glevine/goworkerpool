package main

type queue struct {
	work chan int
}

func NewQueue() *queue {
	return &queue{
		work: make(chan int),
	}
}

func (q *queue) send() chan<- int {
	return q.work
}

func (q *queue) receive() <-chan int {
	return q.work
}

func (q *queue) close() {
	close(q.work)
}
