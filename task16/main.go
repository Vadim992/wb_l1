package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{10, 5, 6, 9, 7, 8}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(arr)
}
