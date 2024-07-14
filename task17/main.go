package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 6, 7, 8, 9}

	ind := binarySearch(arr, 5)
	fmt.Println(ind)
	if ind != -1 {
		fmt.Println(arr[ind])
	}

}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
