package main

import "time"

// Villager is an Animal Crossing villager.
type Villager struct {
	astrology
	gender
	species

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

	// Hobbies is a collection of interests for the Animal Crossing villager.
	Hobbies []Hobby

	// Name is the name of the Animal Crossing villager.
	Name Name

	// Personality is the personality of the Animal Crossing villager.
	Personality Personality

	// Phrases is a collection of phrases said by the Animal Crossing villager.
	Phrases []Phrase

	// Quotes is a collection of quotes said by the Animal Crossing villager.
	Quotes []string
}
