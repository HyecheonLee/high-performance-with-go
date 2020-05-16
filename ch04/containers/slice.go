package main

import "fmt"

func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func main() {
	slice := []string{"foo", "bar", "baz"}
	slice = append(slice, "tri")
	fmt.Println("Appended Slice: ", slice)
	slice = remove(slice, 2)
	fmt.Println("After Removed Item: ", slice)
}
