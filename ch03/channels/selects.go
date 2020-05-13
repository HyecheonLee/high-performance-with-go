package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan bool)
	ch3 := make(chan rune)

	go func() {
		ch1 <- "channels are fun"
	}()

	go func() {
		ch2 <- true
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- 'r'
	}()
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Channel 1 message: ", msg1)
			fmt.Println("")
		case msg2 := <-ch2:
			fmt.Println("Channel 2 message: ", msg2)
			fmt.Println("")
		case msg3 := <-ch3:
			fmt.Println("Channel 3 message: ", msg3)
			fmt.Println("")
		}
	}
}
