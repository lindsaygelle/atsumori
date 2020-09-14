package atsumori

import "time"

// Queenie is an Animal Crossing villager.
var Queenie = villager{
	queenieAstrology,
	queenieBirthDay,
	queenieBirthMonth,
	queenieBubbleColor,
	queenieCategory,
	queenieClothes,
	queenieClothesColors,
	queenieFlooring,
	queenieFurniture,
	queenieGames,
	queenieGender,
	queenieInterest,
	queenieName,
	queenieNameColor,
	queenieMusic,
	queeniePersonality,
	queenieSpecies,
	queenieStyle,
	queenieWallpaper}

var (
	queenieAstrology     = villagersAstrology{villagerAstrologyScorpio}
	queenieBirthDay      = villagersBirthDay{13}
	queenieBirthMonth    = villagersBirthMonth{time.November} // 11
	queenieBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	queenieCategory      = villagersCategory{villagerCategoryB}
	queenieClothes       = villagersClothes{} // 3386
	queenieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	queenieFlooring      = villagersFlooring{villagerFlooringDarkParquetFlooring}
	queenieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanWardrobe, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureRattanVanity, villagerFurnitureRattanLowTable, villagerFurnitureIronwoodKitchenette, villagerFurnitureRattanBed, villagerFurnitureRattanWasteBin, villagerFurnitureRattanEndTable, villagerFurnitureRattanArmchair, villagerFurniturePortableRecordPlayer, villagerFurnitureHangingTerrarium}}
	queenieGames         = villagersGames{[]VillagerGame{}} // TBD
	queenieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	queenieInterest      = villagersInterest{villagerInterestFashion}
	queenieName          = villagersName{villagerNameQueenie}
	queenieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	queenieMusic         = villagersMusic{villagerMusicKKGumbo}
	queeniePersonality   = villagersPersonality{villagerPersonalitySnooty}
	queenieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	queenieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	queenieWallpaper     = villagersWallpaper{villagerWallpaperAbstractWall}
)
