/*
I'd love to say I figured this out straight away, but I didn't.

It took a bit of googling and some assiance from stack overflow
and good old ChatGPT.

This would be a fun bit of code to extend, so that it would become
a thing for generating images. That's one for later though
*/
package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	// dx and dy are integers and this function expects
	// a slice of slices that contains unsigned 8 bit integers (0-255)

	img := make([][]uint8, dy) // this is the slice of slices

	// now iterate over dy
	for y := 0; y < dy; y++ {

		// same for dx
		for x := 0; x < dx; x ++ {

			// img is a slice of slices
			img[y] = make([]uint8, dx)

			for x := range img[y]{
				img[y][x] = uint8(blue(x, y))
			}
		}
	}

	return img
}

func blue(x, y int) int{
	return x*1 + y*x
}
func main() {
	pic.Show(Pic)
}
