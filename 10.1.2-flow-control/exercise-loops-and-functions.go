package main

/*
This is the fourth time I've came accross this problem.

This is the newton-raphson method.

Take a guess?
- less than, add half the value and try again
- more than, remove half the valye and try again
*/

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {

	// base case, don't go negative
	if x < 0 {
		return x, errors.New("Cannot compute square of a negative number")
	}

	// base case, x is 0
	if x == 0 {
		return x, nil
	}

	// take a guesss - any guess will do
	guess := x / 2
	for math.Abs(guess*guess-x) > 0.000001 { // that's a tolerance and can be adjusted
		guess = (guess + x/guess) / 2
	}

	// and we're done
	return guess, nil
}

func main() {
	fmt.Println(Sqrt(9))
}
