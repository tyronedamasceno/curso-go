package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello world")

	for msg := range channel {
		fmt.Println(msg)
	}
}

func write(text string) <-chan string { // generator returns goroutine
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
