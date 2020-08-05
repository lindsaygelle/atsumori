package atsumori

var _ Villager = villager{}

type Villager interface {
	VillagerAstrologizer
	VillagerCategorizer
	VillagerClotheser
	VillagerColorer
	VillagerGamer
	VillagerGenderer
	VillagerInterester
	VillagerNameColorer
	VillagerNamer
	VillagerMusicer
	VillagerPersonalitier
	VillagerSpecieser
	VillagerStyler
}

type villager struct {
	villagersAstrology
	villagersCategory
	villagersClothes
	villagersColors
	villagersGames
	villagersGender
	villagersInterest
	villagersName
	villagersNameColors
	villagersMusic
	villagersPersonality
	villagersSpecies
	villagersStyle
}
