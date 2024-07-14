package main

import "fmt"

func main() {
	set1 := []int{3, 5, 4, 9, 1, 2, 0, 6, 8, 7}
	set2 := []int{5, 8, 1, 10, 18, 9}

	fmt.Println(createIntersection(set1, set2))
}

func createIntersection(set1, set2 []int) map[int]struct{} {
	m := make(map[int]struct{}, len(set1))
	for _, val := range set1 {
		m[val] = struct{}{}
	}

	minLen := len(set1)
	if len(set2) < minLen {
		minLen = len(set2)
	}

	res := make(map[int]struct{}, minLen)
	for _, val := range set2 {
		if _, ok := m[val]; ok {
			res[val] = struct{}{}
		}
	}

	return res
}
