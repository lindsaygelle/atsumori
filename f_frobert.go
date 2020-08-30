package atsumori

import "time"

// Frobert is an Animal Crossing villager.
var Frobert = villager{
	frobertAstrology,
	frobertBirthDay,
	frobertBirthMonth,
	frobertBubbleColor,
	frobertCategory,
	frobertClothes,
	frobertClothesColors,
	frobertFlooring,
	frobertFurniture,
	frobertGames,
	frobertGender,
	frobertInterest,
	frobertName,
	frobertNameColor,
	frobertMusic,
	frobertPersonality,
	frobertSpecies,
	frobertStyle,
	frobertWallpaper}

var (
	frobertAstrology     = villagersAstrology{villagerAstrologyAquarius}
	frobertBirthDay      = villagersBirthDay{8}
	frobertBirthMonth    = villagersBirthMonth{time.February} // 2
	frobertBubbleColor   = villagersBubbleColor{villagerBubbleColor87E0A9}
	frobertCategory      = villagersCategory{villagerCategoryB}
	frobertClothes       = villagersClothes{villagerClothesStripedTee} // 4275
	frobertClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorColorful}}
	frobertFlooring      = villagersFlooring{villagerFlooringGreenVinylFlooring}
	frobertFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureMiniDIYWorkbench, villagerFurnitureWoodenEndTable, villagerFurnitureDigitalAlarmClock, villagerFurnitureSimpleMediumAvocadoMat, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureWoodenChest, villagerFurnitureYucca, villagerFurnitureToolCart, villagerFurnitureWoodenTable, villagerFurnitureThrowbackWallClock, villagerFurnitureWhiteMessageMat, villagerFurnitureThrowbackSkullRadio, villagerFurnitureThrowbackWrestlingFigure, villagerFurnitureWoodenToolbox}}
	frobertGames         = villagersGames{[]VillagerGame{}} // TBD
	frobertGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	frobertInterest      = villagersInterest{villagerInterestFitness}
	frobertName          = villagersName{villagerNameFrobert}
	frobertNameColor     = villagersNameColor{villagerNameColor219479}
	frobertMusic         = villagersMusic{villagerMusicKKRagtime}
	frobertPersonality   = villagersPersonality{villagerPersonalityJock}
	frobertSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	frobertStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	frobertWallpaper     = villagersWallpaper{villagerWallpaperGreenPlayroomWall}
)
