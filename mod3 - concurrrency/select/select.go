package main

import (
	"fmt"
	"time"
)

func main() {
	chan1, chan2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			chan1 <- "chan1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			chan2 <- "chan2"
		}
	}()

	for {

		select {
		case msgChan1 := <-chan1:
			fmt.Println(msgChan1)
		case msgChan2 := <-chan2:
			fmt.Println(msgChan2)
		}
		// msgChan1 := <-chan1
		// fmt.Println(msgChan1)
		// msgChan2 := <-chan2
		// fmt.Println(msgChan2)
	}
}
