package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting, for some named meat
func Hello(name string) (string, error) {
	// if no name was given return and error
	if name == "" {
		return "", errors.New("empty name")
	}

	// return a greeting
	message := fmt.Sprintf(randomformat(), name)
	return message, nil
}

// Returns a map that assiciates people with a greeting
func Hellos(names []string) (map[string]string, error) {
	// make a map
	messages := make(map[string]string)

	// loop over the received slice of names,
	// calling hello to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// map message to the name
		messages[name] = message
	}

	return messages, nil

}

// Decides on a random format for a string
func randomformat() string {
	// a slice of formats
	formats := []string{
		"Hi, %v. Welcome",
		"Whats happening, %v!",
		"Hail, %v! Well met",
	}

	// return a random one
	return formats[rand.Intn(len(formats))]

}
