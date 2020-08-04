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
	villagersCategory
	villagersColors
	villagersGender
	villagersName
	villagersMusic
	villagersPersonality
	villagersSpecies
	villagersStyle
}
