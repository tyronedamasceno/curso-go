package main

import (
	"errors"
	"fmt"
)

func main() {
	// var num int16 = 10000000000
	var num int64 = 10000000000
	fmt.Println(num)

	var num2 uint32 = 500
	fmt.Println(num2)

	var num3 float32 = 123.45
	fmt.Println(num3)

	var str string = "banana"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)

	var b1 bool = true
	fmt.Println(b1)

	var err error
	fmt.Println(err)

	var err1 error = errors.New("internal error")
	fmt.Println(err1)
}
