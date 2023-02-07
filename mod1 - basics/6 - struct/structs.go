package main

import "fmt"

type user struct {
	name string
	age  uint
	addr address
}

type address struct {
	street string
	number uint
}

func main() {
	var u user
	fmt.Println(u)
	u.name = "Tyrone"
	u.age = 26
	fmt.Println(u)

	a := address{"ruazite", 10}

	u2 := user{"ze", 20, a}
	fmt.Println(u2)

	u3 := user{age: 55}
	fmt.Println(u3)
}
