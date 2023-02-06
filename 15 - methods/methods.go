package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving %s...\n", u.name)
}

func main() {

	user1 := user{"User 1", 20}
	fmt.Println(user1)
	user1.save()

	user2 := user{"User 2", 15}
	fmt.Println(user2)
	user2.save()
}
