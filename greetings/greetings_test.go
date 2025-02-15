package greetings

import (
	"regexp"
	"testing"
)

// ensure hello returns a valid name
func TestHelloName(t *testing.T) {
	name := "Clarence"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Clarence")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Clarence") = %q, %v, want match for %#q nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
