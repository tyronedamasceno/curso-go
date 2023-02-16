package main

import "fmt"

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		r := <-results
		fmt.Println(r)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for t := range tasks {
		results <- fib(t)
	}
}

func fib(p int) int {
	if p == 0 || p == 1 {
		return p
	}
	return fib(p-1) + fib(p-2)
}
