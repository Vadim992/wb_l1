package main

import "fmt"

func main() {
	arr := make([]int, 0, 10)
	for i := 0; i < 2; i++ {
		arr = append(arr, i)
	}

	i := 0

	copy(arr[i:], arr[i+1:])

	fmt.Println(arr[:len(arr)-1])
}
