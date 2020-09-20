package atsumori

import "time"

// Ricky is an Animal Crossing villager.
var Ricky = villager{
	rickyAstrology,
	rickyBirthDay,
	rickyBirthMonth,
	rickyBubbleColor,
	rickyCategory,
	rickyClothes,
	rickyClothesColors,
	rickyFlooring,
	rickyFurniture,
	rickyGames,
	rickyGender,
	rickyInterest,
	rickyName,
	rickyNameColor,
	rickyMusic,
	rickyPersonality,
	rickySpecies,
	rickyStyle,
	rickyWallpaper}

var (
	rickyAstrology     = villagersAstrology{villagerAstrologyVirgo}
	rickyBirthDay      = villagersBirthDay{14}
	rickyBirthMonth    = villagersBirthMonth{time.September} // 9
	rickyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	rickyCategory      = villagersCategory{villagerCategoryB}
	rickyClothes       = villagersClothes{villagerClothesThreeBallTee} // 2498
	rickyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorRed}}
	rickyFlooring      = villagersFlooring{villagerFlooringSwampFlooring}
	rickyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurniturePondStone, villagerFurnitureOldFashionedWashtub, villagerFurnitureTreeStandee, villagerFurnitureHayBed, villagerFurnitureTreeStandee, villagerFurnitureShantyMat, villagerFurnitureLogStool, villagerFurnitureLantern, villagerFurnitureTreeStandee, villagerFurnitureTreeStandee, villagerFurnitureLogBench, villagerFurniturePortableRadio, villagerFurnitureCampStove}}
	rickyGames         = villagersGames{[]VillagerGame{}} // TBD
	rickyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rickyInterest      = villagersInterest{villagerInterestEducation}
	rickyName          = villagersName{villagerNameRicky}
	rickyNameColor     = villagersNameColor{villagerNameColor9B553A}
	rickyMusic         = villagersMusic{villagerMusicKKSafari}
	rickyPersonality   = villagersPersonality{villagerPersonalityCranky}
	rickySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	rickyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	rickyWallpaper     = villagersWallpaper{villagerWallpaperWoodlandWall}
)
