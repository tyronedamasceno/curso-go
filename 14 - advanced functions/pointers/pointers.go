package main

import "fmt"

func invertSignal(n int) int {
	return n * -1
}

func invertSignalPointer(n *int) {
	*n = *n * -1
}

func main() {
	n := 20
	nInverted := invertSignal(n)
	fmt.Println(nInverted)
	fmt.Println(n)

	invertSignalPointer(&n)
	fmt.Println(n)
}
