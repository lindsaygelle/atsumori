package atsumori

import "time"

// Annalise is an Animal Crossing villager.
var Annalise = villager{
	annaliseAstrology,
	annaliseBirthDay,
	annaliseBirthMonth,
	annaliseBubbleColor,
	annaliseCategory,
	annaliseClothes,
	annaliseClothesColors,
	annaliseFlooring,
	annaliseFurniture,
	annaliseGames,
	annaliseGender,
	annaliseInterest,
	annaliseName,
	annaliseNameColor,
	annaliseMusic,
	annalisePersonality,
	annaliseSpecies,
	annaliseStyle,
	annaliseWallpaper}

var (
	annaliseAstrology     = villagersAstrology{villagerAstrologySagittarius}
	annaliseBirthDay      = villagersBirthDay{2}
	annaliseBirthMonth    = villagersBirthMonth{time.December}
	annaliseBubbleColor   = villagersBubbleColor{villagerBubbleColor7C6559}
	annaliseCategory      = villagersCategory{villagerCategoryB}
	annaliseClothes       = villagersClothes{} // 8190
	annaliseClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorPurple}}
	annaliseFlooring      = villagersFlooring{villagerFlooringYellowFloralFlooring}
	annaliseFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePalmTreeLamp, villagerFurnitureShellBed, villagerFurnitureSurfboard, villagerFurnitureShellSpeaker, villagerFurnitureLifeRing, villagerFurnitureCoconutWallPlanter, villagerFurnitureFanPalm, villagerFurnitureKitchenIsland, villagerFurnitureCoconutJuice, villagerFurnitureWoodenLowTable, villagerFurnitureFruitBasket}}
	annaliseGames         = villagersGames{[]VillagerGame{}} // TBD
	annaliseGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	annaliseInterest      = villagersInterest{villagerInterestFashion}
	annaliseName          = villagersName{villagerNameAnnalise}
	annaliseNameColor     = villagersNameColor{villagerNameColorEAC113}
	annaliseMusic         = villagersMusic{villagerMusicAlohaKK}
	annalisePersonality   = villagersPersonality{villagerPersonalitySnooty}
	annaliseSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	annaliseStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleActive}}
	annaliseWallpaper     = villagersWallpaper{villagerWallpaperTropicalVista}
)
