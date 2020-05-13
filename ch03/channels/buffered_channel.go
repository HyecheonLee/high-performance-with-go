package main

import "fmt"

func main() {
	bufferedChannel := make(chan string, 2)
	defer close(bufferedChannel)
	bufferedChannel <- "foo"
	bufferedChannel <- "bar"

	// Length of channel is 2 because both elements added to channel
	fmt.Println("Channel Length After Add:", len(bufferedChannel))

	// Pop foo and bar off the stack
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	fmt.Println("Channel Length After Pop: ", len(bufferedChannel))

	bufferedChannel <- "baz"

	out := <-bufferedChannel
	fmt.Println(out)
}
