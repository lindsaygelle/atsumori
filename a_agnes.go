package atsumori

import "time"

// Agnes is an Animal Crossing villager.
var Agnes Villager = villager{
	agnesAstrology,
	agnesBirthDay,
	agnesBirthMonth,
	agnesBubbleColor,
	agnesCategory,
	agnesClothes,
	agnesClothesColors,
	agnesFlooring,
	agnesFurniture,
	agnesGames,
	agnesGender,
	agnesInterest,
	agnesName,
	agnesNameColor,
	agnesMusic,
	agnesPersonality,
	agnesSpecies,
	agnesStyle,
	agnesWallpaper}

var (
	agnesAstrology     = villagersAstrology{villagerAstrologyTaurus}
	agnesBirthDay      = villagersBirthDay{21}
	agnesBirthMonth    = villagersBirthMonth{time.April}
	agnesBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	agnesCategory      = villagersCategory{villagerCategoryA}
	agnesClothes       = villagersClothes{} // TBD
	agnesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorWhite}}
	agnesFlooring      = villagersFlooring{villagerFlooringArabesqueFlooring}
	agnesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueConsoleTable, villagerFurnitureAntiqueMiniTable, villagerFurnitureCatGrass, villagerFurnitureCatTower, villagerFurnitureFloralSwag, villagerFurnitureIvoryMediumRoundMat, villagerFurnitureLilyRecordPlayer, villagerFurnitureMonstera, villagerFurniturePetBed, villagerFurniturePetFoodBowl, villagerFurnitureRoseBed, villagerFurnitureTerrarium, villagerFurnitureLCDTV50In}}
	agnesGames         = villagersGames{[]VillagerGame{}} // TBD
	agnesGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	agnesInterest      = villagersInterest{villagerInterestPlay}
	agnesName          = villagersName{villagerNameAgnes}
	agnesNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	agnesMusic         = villagersMusic{}
	agnesPersonality   = villagersPersonality{villagerPersonalityBigSister}
	agnesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	agnesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	agnesWallpaper     = villagersWallpaper{villagerWallpaperGrayMoldedPanelWall}
)
