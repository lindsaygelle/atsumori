package main

// ECard is a Animal Crossing villagers E-Card.
type ECard struct {
	astrology
	gender

	Clothes     string
	Description string
	ID          int
	Letter      string
	Password    string
	Phrase      string
	Species     Species
}
