package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 7, 8, 9}
	ch, chSquare := make(chan int), make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go readData(arr, ch, &wg)
	go squareData(ch, chSquare, &wg)
	go writeData(chSquare, &wg)

	wg.Wait()
}

func readData(arr []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	for _, val := range arr {
		ch <- val
	}
}

func squareData(chRead <-chan int, chWrite chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(chWrite)

	for val := range chRead {
		chWrite <- val * val
	}
}

func writeData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		fmt.Println(val)
	}
}
