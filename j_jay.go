package atsumori

import "time"

// Jay is an Animal Crossing villager.
var Jay = villager{
	jayAstrology,
	jayBirthDay,
	jayBirthMonth,
	jayBubbleColor,
	jayCategory,
	jayClothes,
	jayClothesColors,
	jayFlooring,
	jayFurniture,
	jayGames,
	jayGender,
	jayInterest,
	jayName,
	jayNameColor,
	jayMusic,
	jayPersonality,
	jaySpecies,
	jayStyle,
	jayWallpaper}

var (
	jayAstrology     = villagersAstrology{villagerAstrologyCancer}
	jayBirthDay      = villagersBirthDay{17}
	jayBirthMonth    = villagersBirthMonth{time.July} // 7
	jayBubbleColor   = villagersBubbleColor{villagerBubbleColor6B75CE}
	jayCategory      = villagersCategory{villagerCategoryB}
	jayClothes       = villagersClothes{villagerClothesSixBallTee} // 3323
	jayClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	jayFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	jayFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLogWallMountedClock, villagerFurnitureFloatingBiotopePlanter, villagerFurnitureLogStool, villagerFurnitureLogRoundTable, villagerFurnitureHayBed, villagerFurnitureBirdbath, villagerFurnitureLogBench, villagerFurnitureSimpleDIYWorkbench, villagerFurniturePortableRadio, villagerFurnitureCoconutWallPlanter}}
	jayGames         = villagersGames{[]VillagerGame{}} // TBD
	jayGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	jayInterest      = villagersInterest{villagerInterestFitness}
	jayName          = villagersName{villagerNameJay}
	jayNameColor     = villagersNameColor{villagerNameColor9AE8DF}
	jayMusic         = villagersMusic{villagerMusicKKCountry}
	jayPersonality   = villagersPersonality{villagerPersonalityJock}
	jaySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	jayStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleActive}}
	jayWallpaper     = villagersWallpaper{villagerWallpaperStackedWoodWall}
)
