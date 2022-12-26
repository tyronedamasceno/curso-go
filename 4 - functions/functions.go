package main

import "fmt"

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func mathCalc(n1, n2 int) (int, int) {
	sum := n1 + n2
	diff := n1 - n2
	return sum, diff
}

func main() {
	a := sum(10, 20)
	fmt.Println(a)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	k := f("banana")
	fmt.Println(k)

	rsum, rdiff := mathCalc(10, 5)
	fmt.Println(rsum, rdiff)

	rsum2, _ := mathCalc(10, 5)
	fmt.Println(rsum2)
}
