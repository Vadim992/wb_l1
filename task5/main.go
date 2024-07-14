package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Enter time in second")
	var N int64
	fmt.Scanln(&N)

	ctx, stop := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer stop()

	ch := make(chan int)
	defer close(ch)

	var wg sync.WaitGroup
	wg.Add(2)
	go worker(ctx, ch, &wg)
	go producer(ctx, ch, &wg)

	wg.Wait()
}

func worker(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		select {
		case <-ctx.Done():
			fmt.Println("worker done")
			return
		default:
			fmt.Printf("read data %d\n", val)
		}
	}
}

func producer(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("producer done")
			return
		default:
			ch <- i
		}
	}
}
