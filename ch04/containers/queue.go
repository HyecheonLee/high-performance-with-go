package main

import "fmt"

func main() {

	var simpleQueue []string
	simpleQueue = append(simpleQueue, "Performance ")
	simpleQueue = append(simpleQueue, "Go")

	for len(simpleQueue) > 0 {
		fmt.Println(simpleQueue[0])
		simpleQueue = simpleQueue[1:]
	}
	fmt.Println(simpleQueue)
}
