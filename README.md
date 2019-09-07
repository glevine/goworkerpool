# workerpool
A worker pool in Golang with a completely graceful shutdown based on a canceled context

`go run .`

## alternative shutdown technique

The `Producer` closes the `Queue`, which closes the channel. Once the channel is closed, it begins to send zero values to the consumers. The `Worker` tests for the zero value and initiates its shutdown sequence when a zero value is received. The `Worker` wouldn't need to know about the context in this case.

```golang
# queue.go
func (q *queue) Close() {
	defer fmt.Println("queue: closed")
	close(q.work)
}

# producer.go
func (p *producer) Start(ctx context.Context) {
	fmt.Println("producer: starting")
	defer fmt.Println("producer: shutdown")
	defer p.queue.Close()

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

# worker.go
func (w *worker) Start(ctx context.Context) {
	fmt.Println("worker: starting")
	defer fmt.Println("worker: shutdown")

	for {
		select {
		case i := <-w.queue.Receive():
			if i > 0 {
				fmt.Println("worker: ", i)
			} else {
				return
			}
		}
	}
}
```