package atsumori

import "time"

// Purrl is an Animal Crossing villager.
var Purrl = villager{
	purrlAstrology,
	purrlBirthDay,
	purrlBirthMonth,
	purrlBubbleColor,
	purrlCategory,
	purrlClothes,
	purrlClothesColors,
	purrlFlooring,
	purrlFurniture,
	purrlGames,
	purrlGender,
	purrlInterest,
	purrlName,
	purrlNameColor,
	purrlMusic,
	purrlPersonality,
	purrlSpecies,
	purrlStyle,
	purrlWallpaper}

var (
	purrlAstrology     = villagersAstrology{villagerAstrologyGemini}
	purrlBirthDay      = villagersBirthDay{29}
	purrlBirthMonth    = villagersBirthMonth{time.May} // 5
	purrlBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	purrlCategory      = villagersCategory{villagerCategoryB}
	purrlClothes       = villagersClothes{villagerClothesKungFuTee} // 2781
	purrlClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorBlue}}
	purrlFlooring      = villagersFlooring{villagerFlooringDarkHerringboneFlooring}
	purrlFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBed, villagerFurnitureAntiqueClock, villagerFurnitureAntiqueConsoleTable, villagerFurnitureAntiqueBureau, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueMiniTable, villagerFurnitureAntiqueWardrobe, villagerFurnitureFireplace, villagerFurniturePortableRecordPlayer}}
	purrlGames         = villagersGames{[]VillagerGame{}} // TBD
	purrlGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	purrlInterest      = villagersInterest{villagerInterestFashion}
	purrlName          = villagersName{villagerNamePurrl}
	purrlNameColor     = villagersNameColor{villagerNameColor848484}
	purrlMusic         = villagersMusic{villagerMusicKKChorale}
	purrlPersonality   = villagersPersonality{villagerPersonalitySnooty}
	purrlSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	purrlStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleElegant}}
	purrlWallpaper     = villagersWallpaper{villagerWallpaperArchedWindowWall}
)
