package main

// Card is the Animal Crossing villagers character Card.
type Card struct {
	astrology
	description
	VillagerGender
	species
	villagerName

	Clothes  string
	ID       int
	Letter   string
	Password string
	Phrase   string
}
