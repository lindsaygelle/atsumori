package main

// Villager is an Animal Crossing villager.
type Villager struct {
	astrology
	birthday
	color
	game
	gender
	name
	species

	// Amiibo is the Nintendo Amiibo card.
	Amiibo Amiibo `json:"amiibo"`

	// ECard is the E-card information for the Animal Crossing character.
	ECard ECard `json:"ecard"`

	// Aliases is a collection of alternative names of the Animal Crossing villager.
	Aliases []Alias

	// Appearance is a description of the Animal Crossing villagers appearance.
	Appearance string

	// Description is the verbose description of the Animal Crossing villager.
	Description string

	// Phrases is a collection of phrases said by the Animal Crossing villager.
	Phrases []Phrase

	// Quotes is a collection of quotes said by the Animal Crossing villager.
	Quotes []string
}
