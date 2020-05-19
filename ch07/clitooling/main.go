package main

import (
	"fmt"
	"hyecheonlee/high-performance-with-go/ch07/clitooling/cmd"
	"os"
)

func main() {
	if err := cmd.DateCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
