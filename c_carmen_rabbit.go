package atsumori

import "time"

// CarmenRabbit is an Animal Crossing villager.
var CarmenRabbit = villager{
	carmenRabbitAstrology,
	carmenRabbitBirthDay,
	carmenRabbitBirthMonth,
	carmenRabbitBubbleColor,
	carmenRabbitCategory,
	carmenRabbitClothes,
	carmenRabbitClothesColors,
	carmenRabbitFlooring,
	carmenRabbitFurniture,
	carmenRabbitGames,
	carmenRabbitGender,
	carmenRabbitInterest,
	carmenRabbitName,
	carmenRabbitNameColor,
	carmenRabbitMusic,
	carmenRabbitPersonality,
	carmenRabbitSpecies,
	carmenRabbitStyle,
	carmenRabbitWallpaper}

var (
	carmenRabbitAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	carmenRabbitBirthDay      = villagersBirthDay{6}
	carmenRabbitBirthMonth    = villagersBirthMonth{time.January} // 1
	carmenRabbitBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	carmenRabbitCategory      = villagersCategory{villagerCategoryB}
	carmenRabbitClothes       = villagersClothes{} // 9421
	carmenRabbitClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBeige}}
	carmenRabbitFlooring      = villagersFlooring{villagerFlooringMintDotFlooring}
	carmenRabbitFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanArmchair, villagerFurnitureWoodenBlockBed, villagerFurnitureCuteDIYTable, villagerFurnitureMonstera, villagerFurnitureRattanWasteBin, villagerFurnitureIroningBoard, villagerFurnitureUprightVacuum, villagerFurnitureFloralSwag, villagerFurnitureStandardUmbrellaStand, villagerFurnitureWoodenBlockstool, villagerFurnitureDeluxeWasher, villagerFurnitureWoodenBlockstool, villagerFurnitureRattanLowTable, villagerFurniturePortableRecordPlayer, villagerFurnitureRattanTowelBasket, villagerFurnitureBathroomTowelRack, villagerFurnitureTableLamp, villagerFurnitureLCDTV50In}}
	carmenRabbitGames         = villagersGames{[]VillagerGame{}} // TBD
	carmenRabbitGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	carmenRabbitInterest      = villagersInterest{villagerInterestFashion}
	carmenRabbitName          = villagersName{villagerNameCarmen}
	carmenRabbitNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	carmenRabbitMusic         = villagersMusic{villagerMusicKKStroll}
	carmenRabbitPersonality   = villagersPersonality{villagerPersonalityPeppy}
	carmenRabbitSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	carmenRabbitStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCute}}
	carmenRabbitWallpaper     = villagersWallpaper{villagerWallpaperCuteWhiteWall}
)
