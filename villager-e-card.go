package main

// VillagerECard is the Animal Crossing villagers e-card.
type VillagerECard struct {
	description
	villagerCard
	villagerClothes
	villagerGender

	Clothes  string
	Letter   string
	Password string
	Phrase   string
}

// villagerECard is a composable struct.
type villagerECard struct {
	ECard VillagerECard `json:"e_card"`
}
