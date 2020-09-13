package atsumori

import "time"

// Penelope is an Animal Crossing villager.
var Penelope = villager{
	penelopeAstrology,
	penelopeBirthDay,
	penelopeBirthMonth,
	penelopeBubbleColor,
	penelopeCategory,
	penelopeClothes,
	penelopeClothesColors,
	penelopeFlooring,
	penelopeFurniture,
	penelopeGames,
	penelopeGender,
	penelopeInterest,
	penelopeName,
	penelopeNameColor,
	penelopeMusic,
	penelopePersonality,
	penelopeSpecies,
	penelopeStyle,
	penelopeWallpaper}

var (
	penelopeAstrology     = villagersAstrology{villagerAstrologyAquarius}
	penelopeBirthDay      = villagersBirthDay{5}
	penelopeBirthMonth    = villagersBirthMonth{time.February} // 2
	penelopeBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	penelopeCategory      = villagersCategory{villagerCategoryA}
	penelopeClothes       = villagersClothes{} // 4556
	penelopeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorRed}}
	penelopeFlooring      = villagersFlooring{villagerFlooringRedDotFlooring}
	penelopeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureCuteSofa, villagerFurnitureWoodenLowTable, villagerFurnitureCuteWardrobe, villagerFurnitureCuteWallMountedClock, villagerFurnitureShowerSet, villagerFurnitureClawFootTub, villagerFurnitureCuteFloorLamp, villagerFurnitureCuteVanity, villagerFurnitureCuteBed, villagerFurnitureCuteTeaTable, villagerFurnitureCuteMusicPlayer}}
	penelopeGames         = villagersGames{[]VillagerGame{}} // TBD
	penelopeGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	penelopeInterest      = villagersInterest{villagerInterestFashion}
	penelopeName          = villagersName{villagerNamePenelope}
	penelopeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	penelopeMusic         = villagersMusic{villagerMusicBubblegumKK}
	penelopePersonality   = villagersPersonality{villagerPersonalityPeppy}
	penelopeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	penelopeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleGorgeous}}
	penelopeWallpaper     = villagersWallpaper{villagerWallpaperCuteRedWall}
)
