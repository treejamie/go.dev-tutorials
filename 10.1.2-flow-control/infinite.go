package main

import (
	"fmt"
	"time"
)

func main() {
	// timer one
	timer := time.NewTimer( 2 * time.Second)

	for {
		// infinite loop

		// but prevent stupidity
		<-timer.C
		fmt.Println("Apologies, but you're being silly. Someone has to be the adult here.")
		return
	}
}