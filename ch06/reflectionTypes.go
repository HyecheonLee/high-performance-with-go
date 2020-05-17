package main

import (
	"fmt"
	"reflect"
)

func main() {
	foo := "Hi Go!"
	fooType := reflect.TypeOf(foo)
	fmt.Print("Foo type: ", fooType)
}
