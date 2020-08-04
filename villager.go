package atsumori

var _ Villager = villager{}

type Villager interface {
	VillagerCategorizer
	VillagerColorer
	VillagerGenderer
	VillagerNamer
	VillagerMusicer
	VillagerPersonalitier
	VillagerSpecieser
	VillagerStyler
}

type villager struct {
	villagerCategory
	villagerColors
	villagerGender
	villagerName
	villagerMusic
	villagerPersonality
	villagerSpecies
	villagerStyle
}
