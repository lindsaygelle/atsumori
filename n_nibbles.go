package atsumori

import "time"

// Nibbles is an Animal Crossing villager.
var Nibbles = villager{
	nibblesAstrology,
	nibblesBirthDay,
	nibblesBirthMonth,
	nibblesBubbleColor,
	nibblesCategory,
	nibblesClothes,
	nibblesClothesColors,
	nibblesFlooring,
	nibblesFurniture,
	nibblesGames,
	nibblesGender,
	nibblesInterest,
	nibblesName,
	nibblesNameColor,
	nibblesMusic,
	nibblesPersonality,
	nibblesSpecies,
	nibblesStyle,
	nibblesWallpaper}

var (
	nibblesAstrology     = villagersAstrology{villagerAstrologyCancer}
	nibblesBirthDay      = villagersBirthDay{19}
	nibblesBirthMonth    = villagersBirthMonth{time.July} // 7
	nibblesBubbleColor   = villagersBubbleColor{villagerBubbleColor00D1BD}
	nibblesCategory      = villagersCategory{villagerCategoryB}
	nibblesClothes       = villagersClothes{} // 3287
	nibblesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorYellow}}
	nibblesFlooring      = villagersFlooring{villagerFlooringOliveDesertTileFlooring}
	nibblesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJuicyAppleTV, villagerFurnitureYucca, villagerFurnitureCorkboard, villagerFurnitureDoubleSofa, villagerFurnitureTennisTable, villagerFurniturePearBed, villagerFurnitureNaturalGardenTable, villagerFurnitureUnfinishedPuzzle, villagerFurnitureOrangeEndTable, villagerFurnitureNaturalGardenChair, villagerFurnitureElectricKickScooter, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureCuteMusicPlayer}}
	nibblesGames         = villagersGames{[]VillagerGame{}} // TBD
	nibblesGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	nibblesInterest      = villagersInterest{villagerInterestFashion}
	nibblesName          = villagersName{villagerNameNibbles}
	nibblesNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	nibblesMusic         = villagersMusic{villagerMusicKKDisco}
	nibblesPersonality   = villagersPersonality{villagerPersonalityPeppy}
	nibblesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	nibblesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleActive}}
	nibblesWallpaper     = villagersWallpaper{villagerWallpaperPearWall}
)
