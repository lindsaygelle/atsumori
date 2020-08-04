package atsumori

var _ Villager = villager{}

type Villager interface {
	VillagerCategorizer
	VillagerColorer
	VillagerGamer
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
	villagersGames
	villagersGender
	villagersInterest
	villagersName
	villagersMusic
	villagersPersonality
	villagersSpecies
	villagersStyle
}
