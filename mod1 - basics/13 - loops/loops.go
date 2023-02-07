package main

import (
	"fmt"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	// time.Sleep(time.Second)
	// 	i++
	// 	fmt.Println(i)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// 	// time.Sleep(time.Second)
	// }

	numbers := [5]int{1, 2, 3, 4, 5}

	for idx, n := range numbers {
		fmt.Println(idx, n)
	}

	for idx, l := range "tyrone" {
		fmt.Println(idx, l, string(l))
	}
}
