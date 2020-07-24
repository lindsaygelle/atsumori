package main

import "time"

// Villager is an Animal Crossing villager.
type Villager struct {
	astrology
	color
	goal
	gender
	hobby
	name
	personality
	species
	style

	// Aliases is a collection of alternative names of the Animal Crossing villager.
	Aliases []Alias

	// Appearance is a description of the Animal Crossing villagers appearance.
	Appearance string

	// Birthday is the birth date of the Animal Crossing villager.
	Birthday time.Time

	// Description is the verbose description of the Animal Crossing villager.
	Description string

	// ECard is the Animal Crossing villager E-Card.
	ECard ECard

	// Games is a collection of Animal Crossing games the villager has appeared in.
	Games []Game

	// Phrases is a collection of phrases said by the Animal Crossing villager.
	Phrases []Phrase

	// Quotes is a collection of quotes said by the Animal Crossing villager.
	Quotes []string
}
