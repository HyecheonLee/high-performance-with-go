package main

import (
	"fmt"
	"sort"
)

func main() {
	intpu := []string{"foo", "bar", "baz"}
	var result []string

	func() {
		result = append(intpu, "abc")
		result = append(result, "def")
		sort.Sort(sort.StringSlice(result))
	}()
	fmt.Println(result)
}
