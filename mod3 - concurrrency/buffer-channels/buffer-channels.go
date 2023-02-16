package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "hello world"
	channel <- "coding go"

	msg := <-channel
	fmt.Println(msg)

	msg2 := <-channel
	fmt.Println(msg2)

	channel <- "channels 3"

	msg3 := <-channel
	fmt.Println(msg3)
}
