package atsumori

import "time"

// Gladys is an Animal Crossing villager
var Gladys = villager{
	gladysAstrology,
	gladysBirthDay,
	gladysBirthMonth,
	gladysBubbleColor,
	gladysCategory,
	gladysClothes,
	gladysClothesColors,
	gladysFlooring,
	gladysFurniture,
	gladysGames,
	gladysGender,
	gladysInterest,
	gladysName,
	gladysNameColor,
	gladysMusic,
	gladysPersonality,
	gladysSpecies,
	gladysStyle,
	gladysWallpaper}

var (
	gladysAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	gladysBirthDay      = villagersBirthDay{15}
	gladysBirthMonth    = villagersBirthMonth{time.January} // 1
	gladysBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	gladysCategory      = villagersCategory{villagerCategoryB}
	gladysClothes       = villagersClothes{villagerClothesMistyTee} // 8201
	gladysClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorPink}}
	gladysFlooring      = villagersFlooring{villagerFlooringTatami}
	gladysFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureTraditionalTeaSet, villagerFurnitureZenCushion, villagerFurnitureElaborateKimonoStand, villagerFurnitureKimonoStand, villagerFurnitureLowScreen, villagerFurniturePaperLantern, villagerFurnitureFuton, villagerFurniturePileOfZenCushions, villagerFurnitureLoom}}
	gladysGames         = villagersGames{[]VillagerGame{}} // TBD
	gladysGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	gladysInterest      = villagersInterest{villagerInterestEducation}
	gladysName          = villagersName{villagerNameGladys}
	gladysNameColor     = villagersNameColor{villagerNameColor848484}
	gladysMusic         = villagersMusic{villagerMusicKKFolk}
	gladysPersonality   = villagersPersonality{villagerPersonalityNormal}
	gladysSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	gladysStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCute}}
	gladysWallpaper     = villagersWallpaper{villagerWallpaperShojiScreen}
)
