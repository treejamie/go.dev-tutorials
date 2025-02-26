package main

import "fmt"


type Number interface {
	int64 | float64
}


func main(){
	// initialise a map for integer values
	ints := map[string] int64 {
		"first": 34,
		"second": 12,
	}


	// initialise a map for float values
	floats := map[string] float64 {
		"first": 34.98,
		"second": 12.99,
	}
	fmt.Printf("Non-generic sums: %v and %v\n",
		SumInts(ints),
		sumFloats(floats),
	)

	fmt.Printf("generic sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)

	fmt.Printf("generic sums with constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}


// SumInts together

func SumInts(m map[string] int64) int64{
	var s int64
	for _, v := range m {
		s +=v
	}
	return s
}


// SumFloats together

func sumFloats(m map[string] float64) float64{
	var s float64
	for _, v := range m {
		s +=v
	}
	return s
}


// the generic version
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K] V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// now use the type delcared at the top
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m{
		s += v
	}
	return s
}