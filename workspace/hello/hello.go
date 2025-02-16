package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

// This is a comment. It has a purpose.
func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
