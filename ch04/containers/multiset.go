package main

import (
	"fmt"

	"github.com/soniakeys/multiset"
)

func main() {

	x := multiset.Multiset{"foo": 1, "bar": 2, "baz": 3}
	fmt.Println("x: ", x)
	// Create a scaled version of x
	y := multiset.Scale(x, 2)
	fmt.Println("y: ", y)
	fmt.Print("x is a subset of y: ")
	fmt.Println(multiset.Subset(x, y))

	fmt.Print("Cardinality of x: ")
	fmt.Println(x.Cardinality())

}
