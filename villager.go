package atsumori

var _ Villager = villager{}

type Villager interface {
	VillagerCategorizer
	VillagerColorer
	VillagerGenderer
	VillagerInterester
	VillagerNamer
	VillagerMusicer
	VillagerPersonalitier
	VillagerSpecieser
	VillagerStyler
}

type villager struct {
	villagersCategory
	villagersColors
	villagersGender
	villagersInterest
	villagersName
	villagersMusic
	villagersPersonality
	villagersSpecies
	villagersStyle
}
