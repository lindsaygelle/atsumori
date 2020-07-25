package main

// Astrology is the astrological information for Animal Crossing villagers.
type Astrology struct {

	// Name is the name of the Astrological star sign.
	Name VillagerAstrology `json:"name"`

	// Villagers is a collection of Animal Crossing villagers.
	Villagers []VillagerName `json:"villagers"`
}
