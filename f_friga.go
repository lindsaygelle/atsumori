package atsumori

import "time"

// Friga is an Animal Crossing villager.
var Friga = villager{
	frigaAstrology,
	frigaBirthDay,
	frigaBirthMonth,
	frigaBubbleColor,
	frigaCategory,
	frigaClothes,
	frigaClothesColors,
	frigaFlooring,
	frigaFurniture,
	frigaGames,
	frigaGender,
	frigaInterest,
	frigaName,
	frigaNameColor,
	frigaMusic,
	frigaPersonality,
	frigaSpecies,
	frigaStyle,
	frigaWallpaper}

var (
	frigaAstrology     = villagersAstrology{villagerAstrologyLibra}
	frigaBirthDay      = villagersBirthDay{16}
	frigaBirthMonth    = villagersBirthMonth{time.October} // 10
	frigaBubbleColor   = villagersBubbleColor{villagerBubbleColorF2BDC7}
	frigaCategory      = villagersCategory{villagerCategoryB}
	frigaClothes       = villagersClothes{} // 8509
	frigaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorBlack}}
	frigaFlooring      = villagersFlooring{villagerFlooringStoneTile}
	frigaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanWardrobe, villagerFurnitureRattanStool, villagerFurnitureRattanArmchair, villagerFurnitureRattanTowelBasket, villagerFurnitureTanklessToilet, villagerFurnitureRattanVanity, villagerFurniturePlainSink, villagerFurnitureFrozenPartition, villagerFurnitureClawFootTub, villagerFurnitureRattanWasteBin, villagerFurnitureRattanEndTable, villagerFurnitureRattanLowTable, villagerFurnitureRattanTableLamp, villagerFurnitureAirConditioner, villagerFurnitureBathroomTowelRack, villagerFurnitureShowerSet, villagerFurnitureCuteMusicPlayer}}
	frigaGames         = villagersGames{[]VillagerGame{}} // TBD
	frigaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	frigaInterest      = villagersInterest{villagerInterestFashion}
	frigaName          = villagersName{villagerNameFriga}
	frigaNameColor     = villagersNameColor{villagerNameColor634B4B}
	frigaMusic         = villagersMusic{villagerMusicFarewell}
	frigaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	frigaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	frigaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	frigaWallpaper     = villagersWallpaper{villagerWallpaperWhiteSubwayTileWall}
)
