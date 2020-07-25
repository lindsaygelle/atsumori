package main

// VillagerECard is the Animal Crossing villagers e-card.
type VillagerECard struct {
	description
	villagerCard
	villagerGender

	Clothes  string
	Letter   string
	Password string
	Phrase   string
}

type villagerECard struct {
	ECard VillagerECard `json:"e_card"`
}
