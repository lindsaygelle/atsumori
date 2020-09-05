package atsumori

import "time"

// Molly is an Animal Crossing villager.
var Molly = villager{
	mollyAstrology,
	mollyBirthDay,
	mollyBirthMonth,
	mollyBubbleColor,
	mollyCategory,
	mollyClothes,
	mollyClothesColors,
	mollyFlooring,
	mollyFurniture,
	mollyGames,
	mollyGender,
	mollyInterest,
	mollyName,
	mollyNameColor,
	mollyMusic,
	mollyPersonality,
	mollySpecies,
	mollyStyle,
	mollyWallpaper}

var (
	mollyAstrology     = villagersAstrology{villagerAstrologyPisces}
	mollyBirthDay      = villagersBirthDay{7}
	mollyBirthMonth    = villagersBirthMonth{time.March} // 3
	mollyBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	mollyCategory      = villagersCategory{villagerCategoryA}
	mollyClothes       = villagersClothes{} // 8885
	mollyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorPink}}
	mollyFlooring      = villagersFlooring{villagerFlooringLightHerringboneFlooring}
	mollyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBureau, villagerFurnitureWoodenSimpleBed, villagerFurnitureFancyViolin, villagerFurnitureDoubleSofa, villagerFurnitureMusicStand, villagerFurnitureWoodenBookshelf, villagerFurniturePottedIvy, villagerFurnitureWoodenLowTable, villagerFurnitureWoodenChest, villagerFurnitureTeaSet, villagerFurnitureIvorySmallRoundMat, villagerFurnitureRedDottedRug, villagerFurnitureCuteMusicPlayer, villagerFurnitureYucca}}
	mollyGames         = villagersGames{[]VillagerGame{}} // TBD
	mollyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	mollyInterest      = villagersInterest{villagerInterestNature}
	mollyName          = villagersName{villagerNameMolly}
	mollyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	mollyMusic         = villagersMusic{villagerMusicKKBallad}
	mollyPersonality   = villagersPersonality{villagerPersonalityNormal}
	mollySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	mollyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	mollyWallpaper     = villagersWallpaper{villagerWallpaperLatticeWall}
)
