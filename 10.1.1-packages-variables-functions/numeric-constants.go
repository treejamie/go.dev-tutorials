package main

import "fmt"


const (
	// Create a huge number by shifting a bit left by 100 places.
	// aka take 1 in binary (actually 1) and add one hundred zeros
	Big = 1 << 100 // I cannot even think of what the value of that is

	// However, shift it back 99 places and I can understand "2"
	Smol = Big >> 99
)

func needInt(x int) int {return x*10 + 1} // operator precedence informally done by omitting spaces?
func needFloat(x float64) float64{
	return x * 0.1
}


func main() {
	fmt.Println(needInt(Smol))
	// fmt.Println(needInt(Big)) // cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to needInt (overflows)
	fmt.Println(needFloat(Smol))
	fmt.Println(needFloat(Big))
}