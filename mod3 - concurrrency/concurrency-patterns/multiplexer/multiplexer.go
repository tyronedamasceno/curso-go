package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplex(write("hello world"), write("coding go"))

	for msg := range channel {
		fmt.Println(msg)
	}
}

func multiplex(inChan1, inChan2 <-chan string) <-chan string {
	outChan := make(chan string)

	go func() {
		for {
			select {
			case msg := <-inChan1:
				outChan <- msg
			case msg := <-inChan2:
				outChan <- msg
			}
		}
	}()

	return outChan
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
