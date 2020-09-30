package atsumori

import "time"

// Tammi is an Animal Crossing villager.
var Tammi = villager{
	tammiAstrology,
	tammiBirthDay,
	tammiBirthMonth,
	tammiBubbleColor,
	tammiCategory,
	tammiClothes,
	tammiClothesColors,
	tammiFlooring,
	tammiFurniture,
	tammiGames,
	tammiGender,
	tammiInterest,
	tammiName,
	tammiNameColor,
	tammiMusic,
	tammiPersonality,
	tammiSpecies,
	tammiStyle,
	tammiWallpaper}

var (
	tammiAstrology     = villagersAstrology{villagerAstrologyAries}
	tammiBirthDay      = villagersBirthDay{2}
	tammiBirthMonth    = villagersBirthMonth{time.April} // 4
	tammiBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	tammiCategory      = villagersCategory{villagerCategoryB}
	tammiClothes       = villagersClothes{villagerClothesSilkFloralPrintShirt} // 9138
	tammiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorGreen}}
	tammiFlooring      = villagersFlooring{villagerFlooringBambooFlooring}
	tammiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCherryBlossomBranches, villagerFurnitureImperialDiningTable, villagerFurnitureImperialDiningChair, villagerFurnitureBambooSpeaker, villagerFurnitureHangingScroll, villagerFurnitureImperialPartition, villagerFurnitureImperialLowTable, villagerFurnitureClayFurnace, villagerFurnitureSteamerBasketSet, villagerFurnitureSteamerBasketSet, villagerFurnitureServingCart, villagerFurnitureImperialDecorativeShelves, villagerFurnitureTraditionalTeaSet}}
	tammiGames         = villagersGames{[]VillagerGame{}} // TBD
	tammiGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tammiInterest      = villagersInterest{villagerInterestFashion}
	tammiName          = villagersName{villagerNameTammi}
	tammiNameColor     = villagersNameColor{villagerNameColor9B553A}
	tammiMusic         = villagersMusic{villagerMusicImperialKK}
	tammiPersonality   = villagersPersonality{villagerPersonalityPeppy}
	tammiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMonkey}}
	tammiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleActive}}
	tammiWallpaper     = villagersWallpaper{villagerWallpaperBambooWall}
)
