package main

// VillagerECard is the Animal Crossing villagers e-card.
type VillagerECard struct {
	description
	villagerCard
	villagerClothes
	villagerGender

	Letter   string
	Password string
	Phrase   string
}

// villagerECard is a composable struct.
type villagerECard struct {

	// ECard is the Animal Crossing villagers e-card information.
	ECard VillagerECard `json:"e_card"`
}
