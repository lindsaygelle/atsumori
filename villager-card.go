package main

// VillagerCard is the Animal Crossing villagers character card.
type VillagerCard struct {
	astrology
	description
	villagerGender
	villagerSpecies
	villagerName

	Clothes  string
	ID       int
	Letter   string
	Password string
	Phrase   string
}
