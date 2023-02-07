package main

import "fmt"

func mathCalc(n1, n2 int) (sum int, diff int) {
	sum = n1 + n2
	diff = n1 - n2
	return
}

func main() {
	s, d := mathCalc(20, 10)
	fmt.Println(s, d)
}
