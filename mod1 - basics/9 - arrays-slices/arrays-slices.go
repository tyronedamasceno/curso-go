package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int

	fmt.Println(array1)
	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]string{"p1", "p2", "p3", "p4", "p5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 12, 15}

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 20)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)
	fmt.Println(reflect.TypeOf(slice2))

	array2[2] = "banana"
	fmt.Println(slice2)

	// internal arrays
	fmt.Println("--------------")

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 30)
	slice3 = append(slice3, 31)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
