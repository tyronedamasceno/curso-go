package main

import "fmt"

func main() {
	user := map[string]string{
		"name":   "tyrone",
		"school": "ufrn",
	}
	fmt.Println(user)
	fmt.Println(user["name"])

	delete(user, "school")
	fmt.Println(user)

	user["xpto"] = "banana"
	fmt.Println(user)
}
