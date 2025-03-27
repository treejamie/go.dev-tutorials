package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 { // note however the statements don't need to be in parenthesis.
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}