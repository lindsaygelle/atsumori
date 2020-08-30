package atsumori

import "time"

// Chow is an Animal Crossing villager.
var Chow = villager{
	chowAstrology,
	chowBirthDay,
	chowBirthMonth,
	chowBubbleColor,
	chowCategory,
	chowClothes,
	chowClothesColors,
	chowFlooring,
	chowFurniture,
	chowGames,
	chowGender,
	chowInterest,
	chowName,
	chowNameColor,
	chowMusic,
	chowPersonality,
	chowSpecies,
	chowStyle,
	chowWallpaper}

var (
	chowAstrology     = villagersAstrology{villagerAstrologyCancer}
	chowBirthDay      = villagersBirthDay{22}
	chowBirthMonth    = villagersBirthMonth{time.July} // 7
	chowBubbleColor   = villagersBubbleColor{villagerBubbleColorF2BDC7}
	chowCategory      = villagersCategory{villagerCategoryB}
	chowClothes       = villagersClothes{} // 4327
	chowClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorWhite}}
	chowFlooring      = villagersFlooring{villagerFlooringBambooFlooring}
	chowFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBambooPartition, villagerFurnitureBambooPartition, villagerFurnitureBambooPartition, villagerFurnitureBambooFloorLamp, villagerFurnitureBambooFloorLamp, villagerFurnitureBambooBasket, villagerFurnitureBambooNoodleSlide, villagerFurnitureBambooDoll, villagerFurnitureBambooDoll, villagerFurnitureBambooBench, villagerFurniturePhonograph}}
	chowGames         = villagersGames{[]VillagerGame{}} // TBD
	chowGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	chowInterest      = villagersInterest{villagerInterestFitness}
	chowName          = villagersName{villagerNameChow}
	chowNameColor     = villagersNameColor{villagerNameColor634B4B}
	chowMusic         = villagersMusic{villagerMusicImperialKK}
	chowPersonality   = villagersPersonality{villagerPersonalityCranky}
	chowSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	chowStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	chowWallpaper     = villagersWallpaper{villagerWallpaperBambooGroveWall}
)
