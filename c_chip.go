package atsumori

import "time"

// Chip is an Animal Crossing villager.
var Chip = villager{
	chipAstrology,
	chipBirthDay,
	chipBirthMonth,
	chipBubbleColor,
	chipCategory,
	chipClothes,
	chipClothesColors,
	chipFlooring,
	chipFurniture,
	chipGames,
	chipGender,
	chipInterest,
	chipName,
	chipNameColor,
	chipMusic,
	chipPersonality,
	chipSpecies,
	chipStyle,
	chipWallpaper}

var (
	chipAstrology     = villagersAstrology{villagerAstrologySagittarius} // villagerAstrology
	chipBirthDay      = villagersBirthDay{9}
	chipBirthMonth    = villagersBirthMonth{time.December} // 2
	chipBubbleColor   = villagersBubbleColor{villagerBubbleColor77614C}
	chipCategory      = villagersCategory{}
	chipClothes       = villagersClothes{} //
	chipClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	chipFlooring      = villagersFlooring{}
	chipFurniture     = villagersFurniture{[]VillagerFurniture{}}
	chipGames         = villagersGames{[]VillagerGame{}} // TBD
	chipGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	chipInterest      = villagersInterest{}
	chipName          = villagersName{villagerNameChip}
	chipNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	chipMusic         = villagersMusic{} //
	chipPersonality   = villagersPersonality{}
	chipSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBeaver}}
	chipStyle         = villagersStyle{[2]VillagerStyle{}}
	chipWallpaper     = villagersWallpaper{}
)
