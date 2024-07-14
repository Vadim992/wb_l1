package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var num, bit, bitValue int64

	err := getData(in, &num, &bit, &bitValue)
	if err != nil {
		log.Fatal(err)
	}

	res := changeBit(num, bit, bitValue)
	fmt.Printf("%08b, %d\n", res, res)
}

func getData(in *bufio.Reader, num, bit, bitValue *int64) error {
	fmt.Println("Enter int64 number")
	_, err := fmt.Fscanln(in, num)
	if err != nil {
		return err
	}

	fmt.Println("Enter bit: number from 0 to 63")
	_, err = fmt.Fscanln(in, bit)
	if err != nil {
		return err
	}

	fmt.Println("Enter value of bit: number 0 or 1")
	_, err = fmt.Fscanln(in, bitValue)
	if err != nil {
		return err
	}

	return nil
}

func changeBit(num, bit, bitValue int64) int64 {
	uNum := uint64(num)

	bitNumWithOne := uint64(1 << bit)

	if bitValue == 1 {
		return int64(uNum | bitNumWithOne)
	}
	bitNumWithZero := uint64(int64(-bitNumWithOne) - 1)

	return int64(uNum & bitNumWithZero)
}
