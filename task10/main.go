package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	fmt.Println(createMap(arr))
}

func createMap(arr []float64) map[int][]float64 {
	m := make(map[int][]float64, len(arr))
	for _, val := range arr {
		key := int(val/10) * 10
		if _, ok := m[key]; !ok {
			m[key] = make([]float64, 0)
		}

		m[key] = append(m[key], val)
	}

	return m
}
