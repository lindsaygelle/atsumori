package atsumori

import "time"

// Tia is an Animal Crossing villager.
var Tia = villager{
	tiaAstrology,
	tiaBirthDay,
	tiaBirthMonth,
	tiaBubbleColor,
	tiaCategory,
	tiaClothes,
	tiaClothesColors,
	tiaFlooring,
	tiaFurniture,
	tiaGames,
	tiaGender,
	tiaInterest,
	tiaName,
	tiaNameColor,
	tiaMusic,
	tiaPersonality,
	tiaSpecies,
	tiaStyle,
	tiaWallpaper}

var (
	tiaAstrology     = villagersAstrology{villagerAstrologyScorpio}
	tiaBirthDay      = villagersBirthDay{18}
	tiaBirthMonth    = villagersBirthMonth{time.November} // 11
	tiaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	tiaCategory      = villagersCategory{villagerCategoryA}
	tiaClothes       = villagersClothes{} // 3138
	tiaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorWhite}}
	tiaFlooring      = villagersFlooring{villagerFlooringBrownIronParquetFlooring}
	tiaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronGardenTable, villagerFurnitureDoubleSofa, villagerFurnitureMenuChalkboard, villagerFurnitureIronGardenChair, villagerFurnitureIronwoodLowTable, villagerFurnitureTeaSet, villagerFurnitureCuckooClock, villagerFurnitureYellowKitchenMat, villagerFurnitureRefrigerator, villagerFurnitureGasRange, villagerFurnitureLilyRecordPlayer, villagerFurnitureSimpleKettle, villagerFurnitureIronwoodKitchenette, villagerFurnitureCuttingBoard}}
	tiaGames         = villagersGames{[]VillagerGame{}} // TBD
	tiaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tiaInterest      = villagersInterest{villagerInterestNature}
	tiaName          = villagersName{villagerNameTia}
	tiaNameColor     = villagersNameColor{villagerNameColor848484}
	tiaMusic         = villagersMusic{villagerMusicKKBossa}
	tiaPersonality   = villagersPersonality{villagerPersonalityNormal}
	tiaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	tiaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	tiaWallpaper     = villagersWallpaper{villagerWallpaperCafeCurtainWall}
)
