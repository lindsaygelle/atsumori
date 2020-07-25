package main

import "image"

// card is a composable field.
type card struct {
	villagerAstrology
	villagerName
	villagerSpecies

	Back  image.Image
	Front image.Image
	ID    int
}
