package main

// VillagerAmiiboCard is the Animal Crossing villagers Amiibo card.
type VillagerAmiiboCard struct {
	villagerBirthday
	villagerCard
	villagerHandSign
	villagerHobby
	villagerHome
	villagerRoll

	Request string `json:"request"`
}

// villagerAmiiboCard is a composable struct.
type villagerAmiiboCard struct {

	// AmiiboCard is an Animal Crossing villagers Amiibo character card.
	AmiiboCard VillagerAmiiboCard `json:"amiibo_card"`
}
