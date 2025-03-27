package main

import "fmt"

func main(){
	var i, j int = 1, 2 // shortcut the type as they are the same
	k := 3 // the shortcut for assignment when inside a function
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}