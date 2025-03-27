package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//z = sum - x .  <<< error: ./named-results.go:7:2: undefined: z
	return
}

func main(){
	fmt.Println(split(17))
}