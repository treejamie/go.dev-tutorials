/*
Range continued
You can skip the index or value by assigning to _.

for i, _ := range pow
for _, value := range pow
If you only want the index, you can omit the second variable.

for i := range pow
*/
package main

import "fmt"


func main(){
	// make a slice with 10 default values
	pow := make([]int, 10)
	fmt.Println(pow)

	// iterate over it and re-assign the value to a power	
	for i := range pow{
		pow[i] = 1 << uint(i) // == 2**i
	}

	// now iterate over it again, but this time to print it out
	for _, value := range pow{
		fmt.Printf("%d\n", value)
	}
}