package atsumori

var _ Villager = villager{}

type Villager interface {
	villagerAstrology
	villagerBubbleColor
	villagerCategory
	villagerClothes
	villagerColor
	villagerFlooring
	villagerFurniture
	villagerGame
	villagerGender
	villagerInterest
	villagerNameColor
	VillagerNamer
	villagerMusic
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
