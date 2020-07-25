package main

// Villager is an Animal Crossing villager.
type Villager struct {
	color
	description
	games
	names
	villagerAmiiboCard
	villagerAstrology
	villagerBirthday
	villagerECard
	villagerGoal
	villagerGender
	villagerSpecies
	villagerSong
	villagerStyle
	villagerName
	villagerPersonality

	// Appearance is a description of the Animal Crossing villagers appearance.
	Appearance string

	// Phrases is a collection of quotes said by the Animal Crossing villager.
	Phrases []I18N
}
