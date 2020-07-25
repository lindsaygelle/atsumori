package main

import "image"

// card is a composable field.
type card struct {
	villagerAstrology
	villagerName
	villagerSpecies

	// BackImage is the cards back-side image.
	BackImage image.Image `json:"back_image"`

	// FrontImage is the cards front-side image.
	FrontImage image.Image `json:"front_image"`

	// ID is the numeric ID for the card.
	ID int `json:"id"`
}
