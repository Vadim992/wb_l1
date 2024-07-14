package main

import "fmt"

func main() {
	x, y := 10, 20
	fmt.Println(x, y)

	SwapTwoNum(&x, &y)
	fmt.Println(x, y)

	SwapWithMath(&x, &y)
	fmt.Println(x, y)
}

func SwapTwoNum(x, y *int) {
	*x, *y = *y, *x
}

func SwapWithMath(x, y *int) {
	*x += *y
	*y = *x - *y
	*x -= *y
}
