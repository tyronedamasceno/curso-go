package main

import (
	"fmt"
	"modu/temp"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println(("Writing from main file"))
	temp.Write()
	err := checkmail.ValidateFormat("tyronedamasceno@gmail.com")
	fmt.Println(err)
}
