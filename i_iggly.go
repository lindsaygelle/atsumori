package atsumori

import "time"

// Iggly is an Animal Crossing villager.
var Iggly = villager{
	igglyAstrology,
	igglyBirthDay,
	igglyBirthMonth,
	igglyBubbleColor,
	igglyCategory,
	igglyClothes,
	igglyClothesColors,
	igglyFlooring,
	igglyFurniture,
	igglyGames,
	igglyGender,
	igglyInterest,
	igglyName,
	igglyNameColor,
	igglyMusic,
	igglyPersonality,
	igglySpecies,
	igglyStyle,
	igglyWallpaper}

var (
	igglyAstrology     = villagersAstrology{villagerAstrologyScorpio}
	igglyBirthDay      = villagersBirthDay{2}
	igglyBirthMonth    = villagersBirthMonth{time.November} // 11
	igglyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	igglyCategory      = villagersCategory{villagerCategoryB}
	igglyClothes       = villagersClothes{} // 4586
	igglyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlue}}
	igglyFlooring      = villagersFlooring{villagerFlooringIceFlooring}
	igglyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureMenuChalkboard, villagerFurnitureSoftServeLamp, villagerFurnitureIlluminatedTree, villagerFurnitureIlluminatedTree, villagerFurnitureFrozenCounter, villagerFurnitureFrozenTable, villagerFurnitureFrozenTreatSet, villagerFurnitureFrozenTreatSet, villagerFurnitureFrozenTreatSet, villagerFurnitureFrozenTable, villagerFurnitureFrozenTreatSet, villagerFurnitureFrozenTreatSet, villagerFurnitureFrozenTreatSet, villagerFurnitureCuteMusicPlayer, villagerFurnitureShavedIceMaker}}
	igglyGames         = villagersGames{[]VillagerGame{}} // TBD
	igglyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	igglyInterest      = villagersInterest{villagerInterestFitness}
	igglyName          = villagersName{villagerNameIggly}
	igglyNameColor     = villagersNameColor{villagerNameColor848484}
	igglyMusic         = villagersMusic{villagerMusicILoveYou}
	igglyPersonality   = villagersPersonality{villagerPersonalityJock}
	igglySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	igglyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	igglyWallpaper     = villagersWallpaper{villagerWallpaperIceWall}
)
