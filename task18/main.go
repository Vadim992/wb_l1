package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	c := NewCounter()

	wg.Add(1)
	go getCounter(c, &wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(c, &wg)
	}

	wg.Wait()
}

type Counter struct {
	count *int64
}

func NewCounter() *Counter {
	count := int64(0)
	return &Counter{
		count: &count,
	}
}

func (c *Counter) Increment() {
	atomic.AddInt64(c.count, 1)
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	for {
		select {
		case <-exit:
			return
		default:
			c.Increment()
		}
	}
}

func getCounter(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	select {
	case <-exit:
		fmt.Println(*c.count)
	}
}
