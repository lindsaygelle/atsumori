package main

import "time"

// animal is a composable field.
type animal struct {
	astrology
	color
	gender
	name
	species

	// Birthday is the date of birth of the Animal Crossing character.
	Birthday time.Time `json:"birthday"`

	// ECard is the E-Card for the Animal Crossing character.
	ECard ECard `json:"ecard"`
}
