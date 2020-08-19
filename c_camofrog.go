package atsumori

import "time"

// Camofrog is an Animal Crossing villager
var Camofrog = villager{
	camofrogAstrology,
	camofrogBirthDay,
	camofrogBirthMonth,
	camofrogBubbleColor,
	camofrogCategory,
	camofrogClothes,
	camofrogClothesColors,
	camofrogFlooring,
	camofrogFurniture,
	camofrogGames,
	camofrogGender,
	camofrogInterest,
	camofrogName,
	camofrogNameColor,
	camofrogMusic,
	camofrogPersonality,
	camofrogSpecies,
	camofrogStyle,
	camofrogWallpaper}

var (
	camofrogAstrology     = villagersAstrology{villagerAstrologyGemini}
	camofrogBirthDay      = villagersBirthDay{5}
	camofrogBirthMonth    = villagersBirthMonth{time.June} // 6
	camofrogBubbleColor   = villagersBubbleColor{villagerBubbleColor76B788}
	camofrogCategory      = villagersCategory{villagerCategoryB}
	camofrogClothes       = villagersClothes{villagerClothesMVPTee} // 6835
	camofrogClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	camofrogFlooring      = villagersFlooring{villagerFlooringCamoFlooring}
	camofrogFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureProTapeRecorder, villagerFurnitureCampingCot, villagerFurnitureTreeStandee, villagerFurnitureSmoker, villagerFurnitureBirdcage, villagerFurnitureLogStool, villagerFurnitureBirdbath, villagerFurnitureLogRoundTable, villagerFurnitureCampStove, villagerFurnitureLantern}}
	camofrogGames         = villagersGames{[]VillagerGame{}} // TBD
	camofrogGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	camofrogInterest      = villagersInterest{villagerInterestMusic}
	camofrogName          = villagersName{villagerNameCamofrog}
	camofrogNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	camofrogMusic         = villagersMusic{} // K.K. Safari
	camofrogPersonality   = villagersPersonality{villagerPersonalityCranky}
	camofrogSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	camofrogStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	camofrogWallpaper     = villagersWallpaper{villagerWallpaperJungleWall}
)
