package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func txtAndNums(txt string, nums ...int) {
	for _, num := range nums {
		fmt.Println(txt, num)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 10, 20, 30))

	txtAndNums("banana", 1, 2, 3, 4)
}
