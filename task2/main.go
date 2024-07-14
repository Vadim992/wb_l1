package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{2, 4, 6, 8, 10}
	result := make([]int, len(data))

	var wg sync.WaitGroup
	for i, num := range data {
		wg.Add(1)

		go func(i, num int) {
			defer wg.Done()

			result[i] = num * num
		}(i, num)
	}

	wg.Wait()

	for i := range data {
		fmt.Printf("the sqare of %d is %d\n", data[i], result[i])
	}
}
