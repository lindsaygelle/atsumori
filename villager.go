package atsumori

var _ Villager = villager{}

type Villager interface {
	VillagerGenderer
	VillagerNamer
	VillagerPersonalitier
	VillagerSpecieser
}

type villager struct {
	villagerGender
	villagerName
	villagerPersonality
	villagerSpecies
}
