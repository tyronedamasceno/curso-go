package main

import "fmt"

func weekday(number int) string {
	switch number {
	case 1:
		return "sunday"
	case 2:
		return "monday"
	case 3:
		return "tuesday"
	case 4:
		return "wednesday"
	case 5:
		return "thursday"
	case 6:
		return "friday"
	case 7:
		return "saturday"
	default:
		return "invalid"
	}
}

func weekday2(number int) string {
	switch {
	case number == 1:
		return "sunday"
	case number == 2:
		return "monday"
	case number == 3:
		return "tuesday"
	case number == 4:
		return "wednesday"
	case number == 5:
		return "thursday"
	case number == 6:
		return "friday"
	case number == 7:
		return "saturday"
	default:
		return "invalid"
	}
}

func weekday3(number int) string {
	var ans string

	switch number {
	case 1:
		ans = "sunday"
	case 2:
		ans = "monday"
	case 3:
		ans = "tuesday"
	case 4:
		ans = "wednesday"
	case 5:
		ans = "thursday"
	case 6:
		ans = "friday"
	case 7:
		ans = "saturday"
	default:
		ans = "invalid"
	}

	return ans
}

func main() {
	fmt.Println(weekday(1))
	fmt.Println(weekday(3))
	fmt.Println(weekday(10))

	fmt.Println(weekday2(1))
	fmt.Println(weekday2(3))
	fmt.Println(weekday2(10))
}
