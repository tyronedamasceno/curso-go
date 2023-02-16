package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("hello world", channel)

	// for {
	// 	msg, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("the end")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel) // if not close: deadlock
}
