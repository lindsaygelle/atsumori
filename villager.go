package atsumori

var _ Villager = villager{}

// Villager is a representation of an Animal Crossing villager as code.
type Villager interface {
	villagerAstrology
	villagerBubbleColor
	villagerCategory
	villagerClothes
	villagerClothesColor
	villagerFlooring
	villagerFurniture
	villagerGame
	villagerGender
	villagerInterest
	villagerNameColor
	villagerName
	villagerMusic
	villagerPersonality
	villagerSpecies
	villagerStyle
	villagerWallpaper
}

type villager struct {
	villagersAstrology
	villagersBubbleColor
	villagersCategory
	villagersClothes
	villagersClothesColors
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

var (
	// Villagers is a collection of Animal Crossing villagers.
	Villagers = [...]Villager{
		Ace,
		Admiral}
)
