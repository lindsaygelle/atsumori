package main

// VillagerCard is the Animal Crossing villagers character card.
type VillagerCard struct {
	card
	description
	villagerGender

	Clothes  string
	Letter   string
	Password string
	Phrase   string
}
