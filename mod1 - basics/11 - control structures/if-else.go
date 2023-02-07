package main

import "fmt"

func main() {
	number := 10
	if number > 15 {
		fmt.Println("greater than 15")
	} else {
		fmt.Println("smaller or equal to 15")
	}

	if otherNumber := number; otherNumber > 0 {
		fmt.Println("other number is greater than 0")
	}

	// fmt.Println(otherNumber)  otherNumber doesn't exists here
}
