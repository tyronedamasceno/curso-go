package main

import "fmt"

func main() {
	func() {
		fmt.Println("I have no name")
	}()

	func(txt string) {
		fmt.Println(txt)
	}("banana")

	x := func(txt string) string {
		return fmt.Sprintf("Received -> %s", txt)
	}("banana")

	fmt.Println(x)
}
