package atsumori

import "time"

// Merengue is an Animal Crossing villager.
var Merengue = villager{
	merengueAstrology,
	merengueBirthDay,
	merengueBirthMonth,
	merengueBubbleColor,
	merengueCategory,
	merengueClothes,
	merengueClothesColors,
	merengueFlooring,
	merengueFurniture,
	merengueGames,
	merengueGender,
	merengueInterest,
	merengueName,
	merengueNameColor,
	merengueMusic,
	merenguePersonality,
	merengueSpecies,
	merengueStyle,
	merengueWallpaper}

var (
	merengueAstrology     = villagersAstrology{villagerAstrologyPisces}
	merengueBirthDay      = villagersBirthDay{19}
	merengueBirthMonth    = villagersBirthMonth{time.March} // 3
	merengueBubbleColor   = villagersBubbleColor{villagerBubbleColorF2BDC7}
	merengueCategory      = villagersCategory{villagerCategoryA}
	merengueClothes       = villagersClothes{villagerClothesChefsOutfit} // 3177
	merengueClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorRed}}
	merengueFlooring      = villagersFlooring{villagerFlooringMintDotFlooring}
	merengueFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureMenuChalkboard, villagerFurnitureCuteSofa, villagerFurnitureKitchenIsland, villagerFurnitureCuteTeaTable, villagerFurniturePortableRadio, villagerFurniturePottedIvy, villagerFurnitureWoodenLowTable, villagerFurnitureCreamAndSugar, villagerFurnitureTeaSet, villagerFurnitureDinerCounterTable, villagerFurnitureMicrowave, villagerFurnitureStandMixer, villagerFurnitureSoftServeLamp}}
	merengueGames         = villagersGames{[]VillagerGame{}} // TBD
	merengueGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	merengueInterest      = villagersInterest{villagerInterestNature}
	merengueName          = villagersName{villagerNameMerengue}
	merengueNameColor     = villagersNameColor{villagerNameColor634B4B}
	merengueMusic         = villagersMusic{villagerMusicKKBallad}
	merenguePersonality   = villagersPersonality{villagerPersonalityNormal}
	merengueSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRhino}}
	merengueStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	merengueWallpaper     = villagersWallpaper{villagerWallpaperPastelDottedWall}
)
