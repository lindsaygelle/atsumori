package atsumori

import "time"

// Moe is an Animal Crossing villager.
var Moe = villager{
	moeAstrology,
	moeBirthDay,
	moeBirthMonth,
	moeBubbleColor,
	moeCategory,
	moeClothes,
	moeClothesColors,
	moeFlooring,
	moeFurniture,
	moeGames,
	moeGender,
	moeInterest,
	moeName,
	moeNameColor,
	moeMusic,
	moePersonality,
	moeSpecies,
	moeStyle,
	moeWallpaper}

var (
	moeAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	moeBirthDay      = villagersBirthDay{12}
	moeBirthMonth    = villagersBirthMonth{time.January} // 1
	moeBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	moeCategory      = villagersCategory{villagerCategoryB}
	moeClothes       = villagersClothes{villagerClothesSimpleParka} // 7975
	moeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	moeFlooring      = villagersFlooring{villagerFlooringStripeFlooring}
	moeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureWoodenLowTable, villagerFurnitureWoodenBlockChest, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenBlockWallClock, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureThrowbackContainer, villagerFurnitureToyBox, villagerFurnitureWoodenBlockBench, villagerFurnitureThrowbackRocket}}
	moeGames         = villagersGames{[]VillagerGame{}} // TBD
	moeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	moeInterest      = villagersInterest{villagerInterestPlay}
	moeName          = villagersName{villagerNameMoe}
	moeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	moeMusic         = villagersMusic{villagerMusicTwoDaysAgo}
	moePersonality   = villagersPersonality{villagerPersonalityLazy}
	moeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	moeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	moeWallpaper     = villagersWallpaper{villagerWallpaperBlueSubwayTileWall}
)
