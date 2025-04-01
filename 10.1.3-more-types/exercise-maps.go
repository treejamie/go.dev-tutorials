/*

Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// define a map to hold the answers
	foo := map[string]int{}


	// iterate over the strings in the sentence
	for _, word := range strings.Fields(s) {

		// I guess O(n2) is acceptable for this situation
		// but iterate over it again and this time count
		// the amount of times an element matches word.
		hits := 0 
		for _, el := range strings.Fields(s){
			if el == word {
				hits += 1
			}

		}
		// we are now done
		foo[word] = hits
	}

	return foo
}

func main(){
	wc.Test(WordCount)
}