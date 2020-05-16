package main

import (
	"fmt"
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

func main() {
	d := deque.New()
	elements := []string{"foo", "bar", "baz"}
	for i := range elements {
		d.PushLeft(elements[i])
	}
	fmt.Println(d.PopLeft())
	fmt.Println(d.PopRight())
	fmt.Println(d.PopLeft())
}
