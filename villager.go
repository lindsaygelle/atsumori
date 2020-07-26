package main

// Villager is an Animal Crossing villager.
type Villager struct {
	color
	description
	games
	names
	villagerAmiiboCard
	villagerAppearance
	villagerAstrology
	villagerBirthday
	villagerClothes
	villagerECard
	villagerGoal
	villagerGender
	villagerSpecies
	villagerSong
	villagerStyle
	villagerName
	villagerPersonality

	// Phrases is a collection of quotes said by the Animal Crossing villager.
	Phrases []I18N
}
