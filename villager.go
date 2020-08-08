package atsumori

var _ Villager = villager{}

type Villager interface {
	VillagerAstrologizer
	VillagerBubbleColorer
	VillagerCategorizer
	VillagerClotheser
	VillagerColorer
	VillagerFlooringer
	VillagerFurniturer
	VillagerGamer
	VillagerGenderer
	VillagerInterester
	VillagerNameColorer
	VillagerNamer
	VillagerMusicer
	VillagerPersonalitier
	VillagerSpecieser
	VillagerStyler
	VillagerWallpaperer
}

type villager struct {
	villagersAstrology
	villagersBubbleColors
	villagersCategory
	villagersClothes
	villagersColors
	villagersFlooring
	villagersFurniture
	villagersGames
	villagersGender
	villagersInterest
	villagersName
	villagersNameColors
	villagersMusic
	villagersPersonality
	villagersSpecies
	villagersStyle
	villagersWallpaper
}
