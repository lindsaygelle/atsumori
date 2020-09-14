package atsumori

import "time"

// Pinky is an Animal Crossing villager.
var Pinky = villager{
	pinkyAstrology,
	pinkyBirthDay,
	pinkyBirthMonth,
	pinkyBubbleColor,
	pinkyCategory,
	pinkyClothes,
	pinkyClothesColors,
	pinkyFlooring,
	pinkyFurniture,
	pinkyGames,
	pinkyGender,
	pinkyInterest,
	pinkyName,
	pinkyNameColor,
	pinkyMusic,
	pinkyPersonality,
	pinkySpecies,
	pinkyStyle,
	pinkyWallpaper}

var (
	pinkyAstrology     = villagersAstrology{villagerAstrologyVirgo}
	pinkyBirthDay      = villagersBirthDay{9}
	pinkyBirthMonth    = villagersBirthMonth{time.September} // 9
	pinkyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	pinkyCategory      = villagersCategory{villagerCategoryB}
	pinkyClothes       = villagersClothes{villagerClothesSilkFloralPrintShirt} // 3248
	pinkyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorPink}}
	pinkyFlooring      = villagersFlooring{villagerFlooringFloralRushMatFlooring}
	pinkyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureImperialBed, villagerFurnitureImperialPartition, villagerFurnitureBambooBasket329, villagerFurnitureImperialChest, villagerFurnitureImperialLowTable, villagerFurnitureCassettePlayer, villagerFurnitureTraditionalTeaSet, villagerFurnitureWoodenStool, villagerFurnitureFloatingBiotopePlanter}}
	pinkyGames         = villagersGames{[]VillagerGame{}} // TBD
	pinkyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pinkyInterest      = villagersInterest{villagerInterestFashion}
	pinkyName          = villagersName{villagerNamePinky}
	pinkyNameColor     = villagersNameColor{villagerNameColor879B96}
	pinkyMusic         = villagersMusic{villagerMusicImperialKK}
	pinkyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	pinkySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	pinkyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	pinkyWallpaper     = villagersWallpaper{villagerWallpaperBambooWall}
)
