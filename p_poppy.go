package atsumori

import "time"

// Poppy is an Animal Crossing villager.
var Poppy = villager{
	poppyAstrology,
	poppyBirthDay,
	poppyBirthMonth,
	poppyBubbleColor,
	poppyCategory,
	poppyClothes,
	poppyClothesColors,
	poppyFlooring,
	poppyFurniture,
	poppyGames,
	poppyGender,
	poppyInterest,
	poppyName,
	poppyNameColor,
	poppyMusic,
	poppyPersonality,
	poppySpecies,
	poppyStyle,
	poppyWallpaper}

var (
	poppyAstrology     = villagersAstrology{villagerAstrologyLeo}
	poppyBirthDay      = villagersBirthDay{5}
	poppyBirthMonth    = villagersBirthMonth{time.August} // 8
	poppyBubbleColor   = villagersBubbleColor{villagerBubbleColorDB6161}
	poppyCategory      = villagersCategory{villagerCategoryB}
	poppyClothes       = villagersClothes{} // 3168
	poppyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorYellow}}
	poppyFlooring      = villagersFlooring{villagerFlooringColoredLeavesFlooring}
	poppyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureLogGardenLounge, villagerFurnitureTreesBountyArch, villagerFurnitureWoodenBucket, villagerFurnitureLogRoundTable, villagerFurnitureLogBench, villagerFurnitureTinyLibrary, villagerFurnitureTraditionalBalancingToy, villagerFurniturePicnicBasket, villagerFurnitureLogStool, villagerFurniturePortableRecordPlayer, villagerFurnitureLeafCampfire}}
	poppyGames         = villagersGames{[]VillagerGame{}} // TBD
	poppyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	poppyInterest      = villagersInterest{villagerInterestEducation}
	poppyName          = villagersName{villagerNamePoppy}
	poppyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	poppyMusic         = villagersMusic{villagerMusicForestLife}
	poppyPersonality   = villagersPersonality{villagerPersonalityNormal}
	poppySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	poppyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	poppyWallpaper     = villagersWallpaper{villagerWallpaperAutumnWall}
)
