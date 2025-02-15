package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set the log up
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// make a slice of names - make sure clarence is last
	names := []string{"Big lez", "Sassy", "Mike Nolan", "Donny", "Wayne", "Norton", "Quinton", "Clarence"}

	// get the greetings for the slice of names
	message, err := greetings.Hellos(names)

	// handle the error if there was one
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
