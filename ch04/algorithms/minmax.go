package main

import "fmt"

func findMinInt(a []int) int {
	var minInt = a[0]
	for _, i := range a {
		if minInt > i {
			minInt = i
		}
	}
	return minInt
}

func findMaxInt(a []int) int {
	var maxInt = a[0]
	for _, i := range a {
		if maxInt < i {
			maxInt = i
		}
	}
	return maxInt
}
func main() {
	intData := []int{3, 1, 2, 5, 6, 4}
	minResult := findMinInt(intData)
	maxResult := findMaxInt(intData)
	fmt.Println("Minimum value in array: ", minResult)
	fmt.Println("Maximum value in array: ", maxResult)
}
