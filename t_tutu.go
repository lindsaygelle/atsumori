package atsumori

import "time"

// Tutu is an Animal Crossing villager.
var Tutu = villager{
	tutuAstrology,
	tutuBirthDay,
	tutuBirthMonth,
	tutuBubbleColor,
	tutuCategory,
	tutuClothes,
	tutuClothesColors,
	tutuFlooring,
	tutuFurniture,
	tutuGames,
	tutuGender,
	tutuInterest,
	tutuName,
	tutuNameColor,
	tutuMusic,
	tutuPersonality,
	tutuSpecies,
	tutuStyle,
	tutuWallpaper}

var (
	tutuAstrology     = villagersAstrology{villagerAstrologyVirgo}
	tutuBirthDay      = villagersBirthDay{15}
	tutuBirthMonth    = villagersBirthMonth{time.September} // 9
	tutuBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	tutuCategory      = villagersCategory{villagerCategoryB}
	tutuClothes       = villagersClothes{villagerClothesHeartSweater} // 8193
	tutuClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorRed}}
	tutuFlooring      = villagersFlooring{villagerFlooringRedDotFlooring}
	tutuFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureCuteSofa, villagerFurnitureCuteChair, villagerFurnitureCuteBed, villagerFurnitureAirConditioner, villagerFurnitureCuteWallMountedClock, villagerFurnitureCuteWardrobe, villagerFurnitureCuteFloorLamp, villagerFurnitureCuteTeaTable, villagerFurnitureDeluxeWasher, villagerFurnitureCuteMusicPlayer}}
	tutuGames         = villagersGames{[]VillagerGame{}} // TBD
	tutuGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tutuInterest      = villagersInterest{villagerInterestFashion}
	tutuName          = villagersName{villagerNameTutu}
	tutuNameColor     = villagersNameColor{villagerNameColor848484}
	tutuMusic         = villagersMusic{villagerMusicBubblegumKK}
	tutuPersonality   = villagersPersonality{villagerPersonalityPeppy}
	tutuSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	tutuStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	tutuWallpaper     = villagersWallpaper{villagerWallpaperPinkStripedWall}
)
