package atsumori

import "time"

// Greta is an Animal Crossing villager.
var Greta = villager{
	gretaAstrology,
	gretaBirthDay,
	gretaBirthMonth,
	gretaBubbleColor,
	gretaCategory,
	gretaClothes,
	gretaClothesColors,
	gretaFlooring,
	gretaFurniture,
	gretaGames,
	gretaGender,
	gretaInterest,
	gretaName,
	gretaNameColor,
	gretaMusic,
	gretaPersonality,
	gretaSpecies,
	gretaStyle,
	gretaWallpaper}

var (
	gretaAstrology     = villagersAstrology{villagerAstrologyVirgo}
	gretaBirthDay      = villagersBirthDay{5}
	gretaBirthMonth    = villagersBirthMonth{time.September} // 9
	gretaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	gretaCategory      = villagersCategory{villagerCategoryA}
	gretaClothes       = villagersClothes{} // 3536
	gretaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorPurple}}
	gretaFlooring      = villagersFlooring{villagerFlooringCommonFlooring}
	gretaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFirewood, villagerFurnitureClayFurnace, villagerFurnitureRetroFan, villagerFurnitureTapeDeck, villagerFurnitureFuton, villagerFurniturePendulumClock, villagerFurnitureZenCushion, villagerFurnitureHearth, villagerFurniturePaperLantern, villagerFurnitureBambooBasket, villagerFurnitureTatamiMat, villagerFurnitureTatamiMat, villagerFurnitureHangingScroll, villagerFurnitureTatamiMat, villagerFurnitureTatamiMat}}
	gretaGames         = villagersGames{[]VillagerGame{}} // TBD
	gretaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	gretaInterest      = villagersInterest{villagerInterestEducation}
	gretaName          = villagersName{villagerNameGreta}
	gretaNameColor     = villagersNameColor{villagerNameColor848484}
	gretaMusic         = villagersMusic{villagerMusicComradeKK}
	gretaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	gretaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	gretaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleSimple}}
	gretaWallpaper     = villagersWallpaper{villagerWallpaperModernShojiScreenWall}
)
