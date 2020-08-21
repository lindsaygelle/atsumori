package atsumori

import "time"

// Alli is an Animal Crossing villager
var Alli = villager{
	alliAstrology,
	alliBirthDay,
	alliBirthMonth,
	alliBubbleColor,
	alliCategory,
	alliClothes,
	alliClothesColors,
	alliFlooring,
	alliFurniture,
	alliGames,
	alliGender,
	alliInterest,
	alliName,
	alliNameColor,
	alliMusic,
	alliPersonality,
	alliSpecies,
	alliStyle,
	alliWallpaper}

var (
	alliAstrology     = villagersAstrology{villagerAstrologyScorpio}
	alliBirthDay      = villagersBirthDay{8}
	alliBirthMonth    = villagersBirthMonth{time.November} // 11
	alliBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	alliCategory      = villagersCategory{villagerCategoryB}
	alliClothes       = villagersClothes{villagerClothesLeopardTee} // 8199
	alliClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorBrown}}
	alliFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	alliFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanVanity, villagerFurnitureRattanLowTable, villagerFurnitureHiFiStereo, villagerFurnitureRattanEndTable, villagerFurnitureRattanWardrobe, villagerFurnitureRattanArmchair, villagerFurnitureHumidifier, villagerFurnitureRattanBed, villagerFurnitureFoldingFloorLamp, villagerFurnitureHangingTerrarium}}
	alliGames         = villagersGames{[]VillagerGame{}} // TBD
	alliGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	alliInterest      = villagersInterest{villagerInterestFashion}
	alliName          = villagersName{villagerNameAlli}
	alliNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	alliMusic         = villagersMusic{villagerMusicKKSoul}
	alliPersonality   = villagersPersonality{villagerPersonalitySnooty}
	alliSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAlligator}}
	alliStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	alliWallpaper     = villagersWallpaper{villagerWallpaperWoodlandWall}
)
