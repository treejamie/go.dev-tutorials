package main

import "fmt"

// returns a function that returns an int
func fibonacci() func() int{

	// start off with 0 & 1
	a, b := 0, 1

	// this is the bit that does the lifting
	return func() int {
		// a and b are defined, outside, but this inner function can access it
		// so now it's just a case of redefining a and b as the next items in the fib sequence
		a, b = b, a + b 

		// return a, because b exists in relationship to a
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}