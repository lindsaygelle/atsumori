package main

import "image"

// card is a composable field.
type card struct {
	villagerAstrology
	villagerName
	villagerSpecies

	BackImage  image.Image `json:"back_image"`
	FrontImage image.Image `json:"front_image"`
	ID         int         `json:"id"`
}
