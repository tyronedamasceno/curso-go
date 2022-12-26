package main

import "fmt"

func main() {
	// arithmetic
	fmt.Println("sum 3 + 2 == ", 3+2)
	fmt.Println("diff 3 - 2 == ", 3-2)
	fmt.Println("quocient 3 / 2 == ", 3/2)
	fmt.Println("product 3 * 2 == ", 3*2)
	fmt.Println("rest 3 % 2 == ", 3%2)

	fmt.Println("---------------------")

	// attribution
	var x string = "banana"
	y := "banana2"

	fmt.Println(x, y)

	fmt.Println("---------------------")

	// relational
	fmt.Println(2 > 3)
	fmt.Println(2 <= 3)
	fmt.Println(2 == 3)
	fmt.Println(2 != 3)

	fmt.Println("---------------------")

	// logical
	t, f := true, false
	fmt.Println(t && f)
	fmt.Println(t || f)
	fmt.Println(!t)
	fmt.Println(!f)

	fmt.Println("---------------------")

	// inplace

	num := 10
	num++
	fmt.Println(num)
	num += 10
	fmt.Println(num)
	num--
	fmt.Println(num)
	num *= 3
	fmt.Println(num)
	num /= 5
	fmt.Println(num)

	fmt.Println("---------------------")

	if num > 10 {
		fmt.Println("entrou no if")
	}
}
