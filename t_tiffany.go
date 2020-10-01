package atsumori

import "time"

// Tiffany is an Animal Crossing villager.
var Tiffany = villager{
	tiffanyAstrology,
	tiffanyBirthDay,
	tiffanyBirthMonth,
	tiffanyBubbleColor,
	tiffanyCategory,
	tiffanyClothes,
	tiffanyClothesColors,
	tiffanyFlooring,
	tiffanyFurniture,
	tiffanyGames,
	tiffanyGender,
	tiffanyInterest,
	tiffanyName,
	tiffanyNameColor,
	tiffanyMusic,
	tiffanyPersonality,
	tiffanySpecies,
	tiffanyStyle,
	tiffanyWallpaper}

var (
	tiffanyAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	tiffanyBirthDay      = villagersBirthDay{9}
	tiffanyBirthMonth    = villagersBirthMonth{time.January} // 1
	tiffanyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	tiffanyCategory      = villagersCategory{villagerCategoryB}
	tiffanyClothes       = villagersClothes{} // 4347
	tiffanyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorRed}}
	tiffanyFlooring      = villagersFlooring{villagerFlooringRedAndBlackVinylFlooring}
	tiffanyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBilliardTable, villagerFurnitureDartboard, villagerFurnitureDartboard, villagerFurniturePinballMachine, villagerFurnitureDoubleSofa, villagerFurnitureJukebox, villagerFurnitureDinerNeonSign, villagerFurnitureRedCarpet, villagerFurnitureDinerCounterTable, villagerFurniturePinballMachine, villagerFurnitureWallMountedTV50In, villagerFurnitureExitSign, villagerFurnitureWallMountedTV50In}}
	tiffanyGames         = villagersGames{[]VillagerGame{}} // TBD
	tiffanyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tiffanyInterest      = villagersInterest{villagerInterestFashion}
	tiffanyName          = villagersName{villagerNameTiffany}
	tiffanyNameColor     = villagersNameColor{villagerNameColor848484}
	tiffanyMusic         = villagersMusic{villagerMusicKKRagtime}
	tiffanyPersonality   = villagersPersonality{villagerPersonalitySnooty}
	tiffanySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	tiffanyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCool}}
	tiffanyWallpaper     = villagersWallpaper{villagerWallpaperBlackCrownWall}
)
