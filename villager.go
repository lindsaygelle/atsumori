package atsumori

var _ Villager = villager{}

type Villager interface {
	VillagerCategorizer
	VillagerColorer
	VillagerGenderer
	VillagerNamer
	VillagerPersonalitier
	VillagerSpecieser
	VillagerStyler
}

type villager struct {
	villagerCategory
	villagerColors
	villagerGender
	villagerName
	villagerPersonality
	villagerSpecies
	villagerStyle
}
