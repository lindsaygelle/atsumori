package atsumori

import "time"

// Celeste is an Animal Crossing villager.
var Celeste = villager{
	celesteAstrology,
	celesteBirthDay,
	celesteBirthMonth,
	celesteBubbleColor,
	celesteCategory,
	celesteClothes,
	celesteClothesColors,
	celesteFlooring,
	celesteFurniture,
	celesteGames,
	celesteGender,
	celesteInterest,
	celesteName,
	celesteNameColor,
	celesteMusic,
	celestePersonality,
	celesteSpecies,
	celesteStyle,
	celesteWallpaper}

var (
	celesteAstrology     = villagersAstrology{villagerAstrologyVirgo} // villagerAstrology
	celesteBirthDay      = villagersBirthDay{7}
	celesteBirthMonth    = villagersBirthMonth{time.September} // 2
	celesteBubbleColor   = villagersBubbleColor{villagerBubbleColor81280A}
	celesteCategory      = villagersCategory{}
	celesteClothes       = villagersClothes{} //
	celesteClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	celesteFlooring      = villagersFlooring{}
	celesteFurniture     = villagersFurniture{[]VillagerFurniture{}}
	celesteGames         = villagersGames{[]VillagerGame{}} // TBD
	celesteGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	celesteInterest      = villagersInterest{}
	celesteName          = villagersName{villagerNameCeleste}
	celesteNameColor     = villagersNameColor{villagerNameColorF7F0C2}
	celesteMusic         = villagersMusic{} //
	celestePersonality   = villagersPersonality{}
	celesteSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOwl}}
	celesteStyle         = villagersStyle{[2]VillagerStyle{}}
	celesteWallpaper     = villagersWallpaper{}
)
