package main

import (
	"fmt"
	"math"
)

func main() {
	a := NewPoint(1, 1)
	b := NewPoint(4, 5)

	fmt.Println(getDistance(a, b))
}

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func getDistance(a, b *Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}
