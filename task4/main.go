package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	fmt.Println("Enter number of workers (minimum number is 0)")
	var num int
	fmt.Scanln(&num)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ch := make(chan int)
	defer close(ch)

	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go worker(ctx, ch, i, &wg)
	}

	wg.Add(1)
	go producer(ctx, ch, &wg)

	select {
	case <-ctx.Done():
		fmt.Println("program done")
		wg.Wait()
		stop()
	}
}

func worker(ctx context.Context, ch <-chan int, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d done\n", num)
			return
		case val := <-ch:
			fmt.Printf("read data %d from worker %d\n", val, num)
		}
	}
}

func producer(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("producer done\n")
			return
		default:
			ch <- i
		}
	}
}
