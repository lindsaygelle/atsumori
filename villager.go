package atsumori

var _ Villager = villager{}

type Villager interface {
	VillagerColorer
	VillagerGenderer
	VillagerNamer
	VillagerPersonalitier
	VillagerSpecieser
	VillagerStyler
}

type villager struct {
	villagerColors
	villagerGender
	villagerName
	villagerPersonality
	villagerSpecies
	villagerStyle
}
