package atsumori

import "time"

// Axel is an Animal Crossing villager.
var Axel = villager{
	axelAstrology,
	axelBirthDay,
	axelBirthMonth,
	axelBubbleColor,
	axelCategory,
	axelClothes,
	axelClothesColors,
	axelFlooring,
	axelFurniture,
	axelGames,
	axelGender,
	axelInterest,
	axelName,
	axelNameColor,
	axelMusic,
	axelPersonality,
	axelSpecies,
	axelStyle,
	axelWallpaper}

var (
	axelAstrology     = villagersAstrology{villagerAstrologyAries}
	axelBirthDay      = villagersBirthDay{23}
	axelBirthMonth    = villagersBirthMonth{time.March} // 3
	axelBubbleColor   = villagersBubbleColor{villagerBubbleColor459ABA}
	axelCategory      = villagersCategory{villagerCategoryB}
	axelClothes       = villagersClothes{villagerClothesKanjiTee} // 3575
	axelClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorWhite}}
	axelFlooring      = villagersFlooring{villagerFlooringBlueRubberFlooring}
	axelFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBlockStereo, villagerFurnitureTennisTable, villagerFurnitureThrowbackWallClock, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureCandyMachine, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureWoodenBlockstool, villagerFurnitureThrowbackRocket, villagerFurnitureDigitalAlarmClock, villagerFurnitureWoodenBlockBench, villagerFurnitureThrowbackContainer, villagerFurnitureToyBox}}
	axelGames         = villagersGames{[]VillagerGame{}} // TBD
	axelGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	axelInterest      = villagersInterest{villagerInterestFitness}
	axelName          = villagersName{villagerNameAxel}
	axelNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	axelMusic         = villagersMusic{villagerMusicKKDixie}
	axelPersonality   = villagersPersonality{villagerPersonalityJock}
	axelSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	axelStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	axelWallpaper     = villagersWallpaper{villagerWallpaperBluePlayroomWall}
)
