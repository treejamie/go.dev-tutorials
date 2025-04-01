/*
Maps
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string] Vertex


func main(){
	m = make(map[string]Vertex)
	m["bell labs"] = Vertex{
		40.68433,
		-74.39967,
	}
	fmt.Println(m["bell labs"])
}