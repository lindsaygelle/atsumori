package atsumori

import "time"

// Curly is an Animal Crossing villager
var Curly = villager{
	curlyAstrology,
	curlyBirthDay,
	curlyBirthMonth,
	curlyBubbleColor,
	curlyCategory,
	curlyClothes,
	curlyClothesColors,
	curlyFlooring,
	curlyFurniture,
	curlyGames,
	curlyGender,
	curlyInterest,
	curlyName,
	curlyNameColor,
	curlyMusic,
	curlyPersonality,
	curlySpecies,
	curlyStyle,
	curlyWallpaper}

var (
	curlyAstrology     = villagersAstrology{villagerAstrologyLeo}
	curlyBirthDay      = villagersBirthDay{26}
	curlyBirthMonth    = villagersBirthMonth{time.July} // 7
	curlyBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	curlyCategory      = villagersCategory{villagerCategoryB}
	curlyClothes       = villagersClothes{villagerClothesEnergeticSweater} // 5262
	curlyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorColorful}}
	curlyFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	curlyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBathroomTowelRack, villagerFurnitureDoubleSofa, villagerFurnitureGasRange, villagerFurnitureRefrigerator, villagerFurnitureWaterCooler, villagerFurnitureBunkBed, villagerFurnitureBathroomSink, villagerFurnitureWoodenBlockstool, villagerFurnitureWoodenTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureMagneticKnifeRack, villagerFurnitureYellowArgyleRug, villagerFurniturePopUpToaster}}
	curlyGames         = villagersGames{[]VillagerGame{}} // TBD
	curlyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	curlyInterest      = villagersInterest{villagerInterestFitness}
	curlyName          = villagersName{villagerNameCurly}
	curlyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	curlyMusic         = villagersMusic{} // K.K. Reggae
	curlyPersonality   = villagersPersonality{villagerPersonalityJock}
	curlySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	curlyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	curlyWallpaper     = villagersWallpaper{villagerWallpaperCommonWall}
)
