package main

import "fmt"

func fib(p uint) uint {
	if p == 0 || p == 1 {
		return p
	}
	return fib(p-1) + fib(p-2)
}

func main() {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(10))
}
