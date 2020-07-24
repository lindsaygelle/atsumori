package main

// ECard is a Animal Crossing villagers E-Card.
type ECard struct {
	astrology
	gender
	species

	Clothes     string
	Description string
	ID          int
	Letter      string
	Password    string
	Phrase      string
}
