package main

import "fmt"

func recoverExecution() {
	fmt.Println("try recover")
	if r := recover(); r != nil {
		fmt.Println("execution recovered")
	}
}

func isApproved(n1, n2 float32) bool {
	defer recoverExecution()
	avg := (n1 + n2) / 2

	if avg > 6 {
		return true
	} else if avg < 6 {
		return false
	}

	panic("AVG IS EXACTLY 6!")
}

func main() {
	fmt.Println(isApproved(6, 6))
	fmt.Println("post execution")

}
