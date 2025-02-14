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
