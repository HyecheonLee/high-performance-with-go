package main

import (
	"container/list"
	"fmt"
)

func main() {

	ll := list.New()
	three := ll.PushBack(3)           //[3]
	four := ll.InsertBefore(4, three) // [4,3]
	ll.InsertBefore(2, three)         // [4,2,3]
	ll.MoveToBack(four)               // [2,3,4]
	ll.PushFront(1)                   // [1,2,3,4]
	listLength := ll.Len()

	fmt.Printf("ll tpye: %T\n", ll)
	fmt.Println("ll length: ", listLength)

	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
