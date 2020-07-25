package main

// Villager is an Animal Crossing villager.
type Villager struct {
	birthday
	color
	description
	games
	villagerAstrology
	villagerGender
	villagerSpecies
	villagerSong
	villagerStyle
	villagerName
	villagerPersonality

	// Amiibo is the Nintendo Amiibo card.
	Amiibo VillagerAmiiboCard `json:"amiibo"`

	// Card is the Animal Crossing character card.
	Card VillagerCard `json:"card"`

	// Aliases is a collection of alternative names of the Animal Crossing villager.
	Aliases []ISO

	// Appearance is a description of the Animal Crossing villagers appearance.
	Appearance string

	// Phrases is a collection of quotes said by the Animal Crossing villager.
	Phrases []ISO
}
