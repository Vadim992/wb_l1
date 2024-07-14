package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func closeByChannel(ch <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ch:
			fmt.Println("End goroutine by channel")
			return
		default:
			fmt.Println("Hi from closeByChannel")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func closeByContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("End goroutine by context")
			return
		default:
			fmt.Println("Hi from closeByContext")
			time.Sleep(time.Second)
		}
	}
}

func closeBySomeStmt(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		fmt.Println("Hi from closeBySomeStm")
		if val == 1 {
			fmt.Println("End goroutine by stmt")
			return
		}
	}
}

func main() {
	byChan := make(chan struct{})
	bySomeStmt := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)

	go closeByChannel(byChan, &wg)
	go closeByContext(ctx, &wg)
	go closeBySomeStmt(bySomeStmt, &wg)

	time.Sleep(time.Second)
	byChan <- struct{}{}
	bySomeStmt <- 2
	bySomeStmt <- 1

	wg.Wait()

}
