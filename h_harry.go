package atsumori

import "time"

// Harry is an Animal Crossing villager.
var Harry = villager{
	harryAstrology,
	harryBirthDay,
	harryBirthMonth,
	harryBubbleColor,
	harryCategory,
	harryClothes,
	harryClothesColors,
	harryFlooring,
	harryFurniture,
	harryGames,
	harryGender,
	harryInterest,
	harryName,
	harryNameColor,
	harryMusic,
	harryPersonality,
	harrySpecies,
	harryStyle,
	harryWallpaper}

var (
	harryAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	harryBirthDay      = villagersBirthDay{7}
	harryBirthMonth    = villagersBirthMonth{time.January} // 1
	harryBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	harryCategory      = villagersCategory{villagerCategoryB}
	harryClothes       = villagersClothes{villagerClothesCamoTee} // 8420
	harryClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorOrange}}
	harryFlooring      = villagersFlooring{villagerFlooringWhiteMosaicTileFlooring}
	harryFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureShowerSet, villagerFurnitureCypressBathtub, villagerFurnitureSaunaHeater, villagerFurnitureBambooPartition, villagerFurnitureOutdoorBath}}
	harryGames         = villagersGames{[]VillagerGame{}} // TBD
	harryGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	harryInterest      = villagersInterest{villagerInterestEducation}
	harryName          = villagersName{villagerNameHarry}
	harryNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	harryMusic         = villagersMusic{villagerMusicKKRally}
	harryPersonality   = villagersPersonality{villagerPersonalityCranky}
	harrySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHippo}}
	harryStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	harryWallpaper     = villagersWallpaper{villagerWallpaperAquaTileWall}
)
