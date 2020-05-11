package main

import "fmt"

func helloGo() {
	fmt.Println("Hello Go from a Function")
}

func main() {
	helloGo()
	func() {
		fmt.Println("Hello GO from an Anonymous Frunction")
	}()
	var hello func() = func() { fmt.Println("Hello Go from an Anonymous Function Variable") }
	hello()
}
