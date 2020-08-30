package atsumori

import "time"

// Kevin is an Animal Crossing villager.
var Kevin = villager{
	kevinAstrology,
	kevinBirthDay,
	kevinBirthMonth,
	kevinBubbleColor,
	kevinCategory,
	kevinClothes,
	kevinClothesColors,
	kevinFlooring,
	kevinFurniture,
	kevinGames,
	kevinGender,
	kevinInterest,
	kevinName,
	kevinNameColor,
	kevinMusic,
	kevinPersonality,
	kevinSpecies,
	kevinStyle,
	kevinWallpaper}

var (
	kevinAstrology     = villagersAstrology{villagerAstrologyTaurus}
	kevinBirthDay      = villagersBirthDay{26}
	kevinBirthMonth    = villagersBirthMonth{time.April} // 4
	kevinBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	kevinCategory      = villagersCategory{villagerCategoryA}
	kevinClothes       = villagersClothes{villagerClothesAfterSchoolJacket} // 4366
	kevinClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorRed}}
	kevinFlooring      = villagersFlooring{villagerFlooringConcreteFlooring}
	kevinFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurnitureUtilitySink, villagerFurniturePunchingBag, villagerFurniturePullUpBarStand, villagerFurnitureMiniFridge, villagerFurnitureProteinShakerBottle, villagerFurnitureRubberMudMat, villagerFurnitureFoldingChair, villagerFurnitureNeutralCorner, villagerFurnitureCardboardBox, villagerFurnitureRedCorner, villagerFurnitureTapeDeck, villagerFurnitureWallFan}}
	kevinGames         = villagersGames{[]VillagerGame{}} // TBD
	kevinGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	kevinInterest      = villagersInterest{villagerInterestPlay}
	kevinName          = villagersName{villagerNameKevin}
	kevinNameColor     = villagersNameColor{villagerNameColor874C25}
	kevinMusic         = villagersMusic{villagerMusicKKRockabilly}
	kevinPersonality   = villagersPersonality{villagerPersonalityJock}
	kevinSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	kevinStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	kevinWallpaper     = villagersWallpaper{villagerWallpaperGrayShantyWall}
)
