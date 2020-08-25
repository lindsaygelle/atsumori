package atsumori

import "time"

// Gigi is an Animal Crossing villager
var Gigi = villager{
	gigiAstrology,
	gigiBirthDay,
	gigiBirthMonth,
	gigiBubbleColor,
	gigiCategory,
	gigiClothes,
	gigiClothesColors,
	gigiFlooring,
	gigiFurniture,
	gigiGames,
	gigiGender,
	gigiInterest,
	gigiName,
	gigiNameColor,
	gigiMusic,
	gigiPersonality,
	gigiSpecies,
	gigiStyle,
	gigiWallpaper}

var (
	gigiAstrology     = villagersAstrology{villagerAstrologyLeo}
	gigiBirthDay      = villagersBirthDay{11}
	gigiBirthMonth    = villagersBirthMonth{time.August} // 8
	gigiBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	gigiCategory      = villagersCategory{villagerCategoryB}
	gigiClothes       = villagersClothes{} // 4341
	gigiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorWhite}}
	gigiFlooring      = villagersFlooring{villagerFlooringPurpleDesertTileFlooring}
	gigiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRoseBed, villagerFurniturePlainSink, villagerFurnitureKitchenIsland, villagerFurnitureLongBathtub, villagerFurnitureShowerBooth, villagerFurnitureDeluxeWasher, villagerFurnitureAntiqueConsoleTable, villagerFurnitureRattanTowelBasket, villagerFurniturePhonograph, villagerFurnitureFragranceSticks, villagerFurnitureWallMountedTV50In}}
	gigiGames         = villagersGames{[]VillagerGame{}} // TBD
	gigiGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	gigiInterest      = villagersInterest{villagerInterestFashion}
	gigiName          = villagersName{villagerNameGigi}
	gigiNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	gigiMusic         = villagersMusic{villagerMusicKKSwing}
	gigiPersonality   = villagersPersonality{villagerPersonalitySnooty}
	gigiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	gigiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	gigiWallpaper     = villagersWallpaper{villagerWallpaperPurpleRoseWall}
)
