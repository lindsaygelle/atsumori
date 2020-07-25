package main

// VillagerAmiiboCard is the Animal Crossing villagers Amiibo card.
type VillagerAmiiboCard struct {
	birthday
	card
	handsign

	Request   string
	RollValue int
}
