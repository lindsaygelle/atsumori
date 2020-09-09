package atsumori

import "time"

// Paolo is an Animal Crossing villager.
var Paolo = villager{
	paoloAstrology,
	paoloBirthDay,
	paoloBirthMonth,
	paoloBubbleColor,
	paoloCategory,
	paoloClothes,
	paoloClothesColors,
	paoloFlooring,
	paoloFurniture,
	paoloGames,
	paoloGender,
	paoloInterest,
	paoloName,
	paoloNameColor,
	paoloMusic,
	paoloPersonality,
	paoloSpecies,
	paoloStyle,
	paoloWallpaper}

var (
	paoloAstrology     = villagersAstrology{villagerAstrologyTaurus}
	paoloBirthDay      = villagersBirthDay{5}
	paoloBirthMonth    = villagersBirthMonth{time.May} // 5
	paoloBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	paoloCategory      = villagersCategory{villagerCategoryA}
	paoloClothes       = villagersClothes{villagerClothesSimpleParka} // 3056
	paoloClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorLightBlue}}
	paoloFlooring      = villagersFlooring{villagerFlooringBlueMosaicTileFlooring}
	paoloFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureWhirlpoolBath, villagerFurnitureCoconutWallPlanter, villagerFurnitureShowerSet, villagerFurnitureShowerSet, villagerFurnitureBathroomTowelRack, villagerFurnitureCypressBathtub, villagerFurniturePlainSink, villagerFurnitureToiletCleaningSet, villagerFurnitureTanklessToilet, villagerFurnitureIvorySimpleBathMat, villagerFurnitureRattanEndTable, villagerFurnitureRattanTowelBasket}}
	paoloGames         = villagersGames{[]VillagerGame{}} // TBD
	paoloGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	paoloInterest      = villagersInterest{villagerInterestNature}
	paoloName          = villagersName{villagerNamePaolo}
	paoloNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	paoloMusic         = villagersMusic{villagerMusicToTheEdge}
	paoloPersonality   = villagersPersonality{villagerPersonalityLazy}
	paoloSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	paoloStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	paoloWallpaper     = villagersWallpaper{villagerWallpaperBlueDesertTileWall}
)
