package main

import (
	"fmt"
	"math/big"
)

func main() {
	z := big.NewInt(0)

	a := big.NewInt(1248576789)
	b := big.NewInt(2345678922)

	fmt.Println(z.Add(a, b))
	fmt.Println(z.Div(b, a))
	fmt.Println(z.Mul(b, a))
	fmt.Println(z.Sub(b, a))
}
