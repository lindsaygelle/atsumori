package atsumori

import "time"

// Poncho is an Animal Crossing villager.
var Poncho = villager{
	ponchoAstrology,
	ponchoBirthDay,
	ponchoBirthMonth,
	ponchoBubbleColor,
	ponchoCategory,
	ponchoClothes,
	ponchoClothesColors,
	ponchoFlooring,
	ponchoFurniture,
	ponchoGames,
	ponchoGender,
	ponchoInterest,
	ponchoName,
	ponchoNameColor,
	ponchoMusic,
	ponchoPersonality,
	ponchoSpecies,
	ponchoStyle,
	ponchoWallpaper}

var (
	ponchoAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	ponchoBirthDay      = villagersBirthDay{2}
	ponchoBirthMonth    = villagersBirthMonth{time.January} // 1
	ponchoBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	ponchoCategory      = villagersCategory{villagerCategoryB}
	ponchoClothes       = villagersClothes{villagerClothesEnergeticSweater} // 5262
	ponchoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorYellow}}
	ponchoFlooring      = villagersFlooring{villagerFlooringBluePaintFlooring}
	ponchoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBathroomTowelRack, villagerFurnitureSimpleSmallBlueMat, villagerFurnitureSimpleSmallBlueMat, villagerFurnitureBunkBed, villagerFurnitureDeluxeWasher, villagerFurnitureCorkboard, villagerFurnitureRattanTowelBasket, villagerFurniturePlainSink, villagerFurnitureCacaoTree, villagerFurnitureWoodenEndTable, villagerFurniturePortableRecordPlayer, villagerFurnitureSimpleSmallBlueMat, villagerFurnitureShowerBooth, villagerFurnitureBunkBed, villagerFurnitureSimpleNavyBathMat, villagerFurnitureBunkBed, villagerFurnitureAirConditioner, villagerFurnitureWoodenBookshelf, villagerFurnitureToilet}}
	ponchoGames         = villagersGames{[]VillagerGame{}} // TBD
	ponchoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	ponchoInterest      = villagersInterest{villagerInterestFitness}
	ponchoName          = villagersName{villagerNamePoncho}
	ponchoNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	ponchoMusic         = villagersMusic{villagerMusicKKBlues}
	ponchoPersonality   = villagersPersonality{villagerPersonalityJock}
	ponchoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	ponchoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	ponchoWallpaper     = villagersWallpaper{villagerWallpaperBlueSimpleClothWall}
)
