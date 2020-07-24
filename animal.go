package main

import "time"

// animal is a composable field.
type animal struct {
	astrology
	color
	gender
	name
	species

	// Birthday is the birth date for the Animal Crossing character.
	Birthday time.Time `json:"birthday"`

	// Games is the Animal Crossing games the Animal Crossing character has appeared in.
	Games []Game `json:"games"`

	// ECard is the E-card information for the Animal Crossing character.
	ECard ECard `json:"ecard"`
}
