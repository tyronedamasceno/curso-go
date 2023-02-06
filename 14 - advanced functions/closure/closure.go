package main

import "fmt"

func closure() func() {
	txt := "inside closure func"

	f := func() {
		fmt.Println(txt)
	}

	return f
}

func main() {
	txt := "inside main func"

	fmt.Println(txt)

	newfunc := closure()
	newfunc()
}
