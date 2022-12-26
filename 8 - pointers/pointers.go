package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1 // copy

	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	var v3 int
	var p *int

	fmt.Println(v3, p)

	v3 = 100
	p = &v3

	fmt.Println(v3, p, *p)

	v3 = 150

	fmt.Println(v3, p, *p)
}
