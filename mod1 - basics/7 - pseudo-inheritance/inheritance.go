package main

import "fmt"

type person struct {
	name   string
	age    uint
	height int
}

type student struct {
	person
	course string
	year   uint8
}

func main() {

	p1 := person{"Tyrone", 26, 182}
	fmt.Println(p1, p1.name)

	s1 := student{p1, "computer science", 3}
	fmt.Println(s1, s1.name, s1.course)
}
