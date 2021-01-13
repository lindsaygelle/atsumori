package atsumori

import "time"

// Blathers is an Animal Crossing villager.
var Blathers = villager{
	blathersAstrology,
	blathersBirthDay,
	blathersBirthMonth,
	blathersBubbleColor,
	blathersCategory,
	blathersClothes,
	blathersClothesColors,
	blathersFlooring,
	blathersFurniture,
	blathersGames,
	blathersGender,
	blathersInterest,
	blathersName,
	blathersNameColor,
	blathersMusic,
	blathersPersonality,
	blathersSpecies,
	blathersStyle,
	blathersWallpaper}

var (
	blathersAstrology     = villagersAstrology{villagerAstrologyGemini} // villagerAstrology
	blathersBirthDay      = villagersBirthDay{24}
	blathersBirthMonth    = villagersBirthMonth{time.September} // 9
	blathersBubbleColor   = villagersBubbleColor{villagerBubbleColor7D593C}
	blathersCategory      = villagersCategory{}
	blathersClothes       = villagersClothes{} //
	blathersClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	blathersFlooring      = villagersFlooring{}
	blathersFurniture     = villagersFurniture{[]VillagerFurniture{}}
	blathersGames         = villagersGames{[]VillagerGame{}} // TBD
	blathersGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	blathersInterest      = villagersInterest{}
	blathersName          = villagersName{villagerNameBlathers}
	blathersNameColor     = villagersNameColor{villagerNameColorF7F0C2}
	blathersMusic         = villagersMusic{} //
	blathersPersonality   = villagersPersonality{}
	blathersSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOwl}}
	blathersStyle         = villagersStyle{[2]VillagerStyle{}}
	blathersWallpaper     = villagersWallpaper{}
)
