package main

import "time"

// animal is a composable field.
type animal struct {
	astrology
	color
	game
	gender
	name
	species

	// ACard is the Amiibo-Card information for the Animal Crossing character.
	ACard ACard `json:"acard"`

	// Birthday is the birth date for the Animal Crossing character.
	Birthday time.Time `json:"birthday"`

	// ECard is the E-card information for the Animal Crossing character.
	ECard ECard `json:"ecard"`
}
