package main

import "time"

// Villager is a representation of an Animal Crossing villager.
type Villager struct {

	// Birthday is the birth date of the Animal Crossing villager.
	Birthday time.Time

	// Description is the verbose description of the Animal Crossing villager.
	Description string

	// Gender is the biological gender of the Animal Crossing villager.
	Gender Gender

	// Hobbies is a collection of interests for the Animal Crossing villager.
	Hobbies []Hobby

	// Name is the name of the Animal Crossing villager.
	Name string

	// Personality is the personality of the Animal Crossing villager.
	Personality Personality

	// Phrases is a collection of phrases said by the Animal Crossing villager.
	Phrases []Phrase

	// Quotes is a collection of quotes said by the Animal Crossing villager.
	Quotes []string

	// Species is the species of the Animal Crossing villager.
	Species Species
}
