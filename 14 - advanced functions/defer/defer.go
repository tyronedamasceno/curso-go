package main

import "fmt"

func func1() {
	fmt.Println("func1")
}

func func2() {
	fmt.Println("func2")
}

func isApproved(n1, n2 float32) bool {
	defer fmt.Println("avg has been calculated")
	fmt.Println("get in the func")

	avg := (n1 + n2) / 2

	if avg >= 7 {
		return true
	}

	return false
}

func main() {

	// defer func1()
	// func2()

	fmt.Println(isApproved(7, 10))
	fmt.Println(isApproved(7, 3))

}
