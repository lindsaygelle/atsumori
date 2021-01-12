package atsumori

import "time"

// Saharah is an Animal Crossing villager.
var Saharah = villager{
	saharahAstrology,
	saharahBirthDay,
	saharahBirthMonth,
	saharahBubbleColor,
	saharahCategory,
	saharahClothes,
	saharahClothesColors,
	saharahFlooring,
	saharahFurniture,
	saharahGames,
	saharahGender,
	saharahInterest,
	saharahName,
	saharahNameColor,
	saharahMusic,
	saharahPersonality,
	saharahSpecies,
	saharahStyle,
	saharahWallpaper}

var (
	saharahAstrology     = villagersAstrology{villagerAstrologyGemini} // villagerAstrology
	saharahBirthDay      = villagersBirthDay{10}
	saharahBirthMonth    = villagersBirthMonth{time.November} // 11
	saharahBubbleColor   = villagersBubbleColor{villagerBubbleColorFCBB3A}
	saharahCategory      = villagersCategory{}
	saharahClothes       = villagersClothes{} // 
	saharahClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	saharahFlooring      = villagersFlooring{}
	saharahFurniture     = villagersFurniture{[]VillagerFurniture{}}
	saharahGames         = villagersGames{[]VillagerGame{}} // TBD
	saharahGender        = villagersGender{[2]VillagerGender{villagerGenderMale, villagerGenderFemale}}
	saharahInterest      = villagersInterest{}
	saharahName          = villagersName{villagerNameSaharah}
	saharahNameColor     = villagersNameColor{villagerNameColor954F07}
	saharahMusic         = villagersMusic{} // 
	saharahPersonality   = villagersPersonality{}
	saharahSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCamel}}
	saharahStyle         = villagersStyle{[2]VillagerStyle{}}
	saharahWallpaper     = villagersWallpaper{}
)
