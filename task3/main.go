package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{2, 4, 6, 8, 10}
	result := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, num := range data {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			result += num * num
		}(num)
	}

	wg.Wait()

	fmt.Printf("the sum is %d\n", result)
}
