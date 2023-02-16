package main

import (
	"fmt"
	"time"
)

func main() {
	go write("hello world") // goroutine
	write("coding go")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
