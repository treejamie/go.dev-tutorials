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

	// get a greeting and return a message
	message, err := greetings.Hello("Donny")

	// handle the error if there was one
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
