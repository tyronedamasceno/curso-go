package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	d := dog{"Leucocito", "westie", 1}
	fmt.Println(d)

	jsonDog, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonDog)

	fmt.Println(bytes.NewBuffer(jsonDog))

	d2 := map[string]string{
		"name":  "Tekilla",
		"breed": "fox paulistinha",
	}

	jsonDog2, err := json.Marshal(d2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonDog2)

	fmt.Println(bytes.NewBuffer(jsonDog2))

	jsonDog3 := `{"name":"Leucocito","breed":"westie","age":1}`

	var d3 dog

	if err := json.Unmarshal([]byte(jsonDog3), &d3); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d3)

	jsonDog4 := `{"breed":"fox paulistinha","name":"Tekilla"}`
	d4 := make(map[string]string)

	if err := json.Unmarshal([]byte(jsonDog4), &d4); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d4)
}
