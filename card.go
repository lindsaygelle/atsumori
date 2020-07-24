package main

// Card is the Animal Crossing villagers character Card.
type Card struct {
	astrology
	gender
	name
	species

	Clothes     string
	Description string
	ID          int
	Letter      string
	Password    string
	Phrase      string
}
