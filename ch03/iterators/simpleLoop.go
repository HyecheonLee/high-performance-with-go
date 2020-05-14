package iterators

var sumLoop int

func simpleLoop(n int) int {
	for i := 0; i < n; i++ {
		sumLoop += i
	}
	return sumLoop
}
