package main

import "image"

// card is a composable field.
type card struct {
	villagerName

	Back  image.Image
	Front image.Image
	ID    int
}
