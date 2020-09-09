package atsumori

import "time"

// Opal is an Animal Crossing villager.
var Opal = villager{
	opalAstrology,
	opalBirthDay,
	opalBirthMonth,
	opalBubbleColor,
	opalCategory,
	opalClothes,
	opalClothesColors,
	opalFlooring,
	opalFurniture,
	opalGames,
	opalGender,
	opalInterest,
	opalName,
	opalNameColor,
	opalMusic,
	opalPersonality,
	opalSpecies,
	opalStyle,
	opalWallpaper}

var (
	opalAstrology     = villagersAstrology{villagerAstrologyAquarius}
	opalBirthDay      = villagersBirthDay{20}
	opalBirthMonth    = villagersBirthMonth{time.January} // 1
	opalBubbleColor   = villagersBubbleColor{villagerBubbleColor87E0A9}
	opalCategory      = villagersCategory{villagerCategoryB}
	opalClothes       = villagersClothes{villagerClothesFrontTieTee} // 4320
	opalClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorOrange}}
	opalFlooring      = villagersFlooring{villagerFlooringOliveDesertTileFlooring}
	opalFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWhirlpoolBath, villagerFurnitureImperialPartition, villagerFurnitureImperialBed, villagerFurnitureImperialDecorativeShelves, villagerFurnitureBambooStool, villagerFurnitureIronwoodKitchenette, villagerFurnitureRattanTableLamp, villagerFurnitureShowerSet, villagerFurnitureBathroomTowelRack, villagerFurnitureImperialLowTable, villagerFurnitureFragranceSticks, villagerFurnitureImperialChest, villagerFurniturePortableRecordPlayer}}
	opalGames         = villagersGames{[]VillagerGame{}} // TBD
	opalGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	opalInterest      = villagersInterest{villagerInterestFashion}
	opalName          = villagersName{villagerNameOpal}
	opalNameColor     = villagersNameColor{villagerNameColor219479}
	opalMusic         = villagersMusic{villagerMusicKKCruisin}
	opalPersonality   = villagersPersonality{villagerPersonalitySnooty}
	opalSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	opalStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	opalWallpaper     = villagersWallpaper{villagerWallpaperBrownBotanicalTileWall}
)
