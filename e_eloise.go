package atsumori

import "time"

// Eloise is an Animal Crossing villager.
var Eloise = villager{
	eloiseAstrology,
	eloiseBirthDay,
	eloiseBirthMonth,
	eloiseBubbleColor,
	eloiseCategory,
	eloiseClothes,
	eloiseClothesColors,
	eloiseFlooring,
	eloiseFurniture,
	eloiseGames,
	eloiseGender,
	eloiseInterest,
	eloiseName,
	eloiseNameColor,
	eloiseMusic,
	eloisePersonality,
	eloiseSpecies,
	eloiseStyle,
	eloiseWallpaper}

var (
	eloiseAstrology     = villagersAstrology{villagerAstrologySagittarius}
	eloiseBirthDay      = villagersBirthDay{8}
	eloiseBirthMonth    = villagersBirthMonth{time.December} // 12
	eloiseBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	eloiseCategory      = villagersCategory{villagerCategoryB}
	eloiseClothes       = villagersClothes{} // 8818
	eloiseClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorOrange}}
	eloiseFlooring      = villagersFlooring{villagerFlooringBrownFloralFlooring}
	eloiseFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAirConditioner, villagerFurnitureRattanBed, villagerFurnitureRattanWardrobe, villagerFurnitureRattanArmchair, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureIronwoodClock, villagerFurnitureRattanLowTable, villagerFurnitureNailArtSet, villagerFurnitureHangingTerrarium, villagerFurnitureKitchenIsland, villagerFurnitureMixer, villagerFurnitureOrangeEndTable, villagerFurnitureRetroStereo}}
	eloiseGames         = villagersGames{[]VillagerGame{}} // TBD
	eloiseGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	eloiseInterest      = villagersInterest{villagerInterestFashion}
	eloiseName          = villagersName{villagerNameEloise}
	eloiseNameColor     = villagersNameColor{villagerNameColor9B553A}
	eloiseMusic         = villagersMusic{villagerMusicKKSoul}
	eloisePersonality   = villagersPersonality{villagerPersonalitySnooty}
	eloiseSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	eloiseStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleSimple}}
	eloiseWallpaper     = villagersWallpaper{villagerWallpaperYellowIntricateWall}
)
