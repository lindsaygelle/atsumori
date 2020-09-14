package atsumori

import "time"

// Pippy is an Animal Crossing villager.
var Pippy = villager{
	pippyAstrology,
	pippyBirthDay,
	pippyBirthMonth,
	pippyBubbleColor,
	pippyCategory,
	pippyClothes,
	pippyClothesColors,
	pippyFlooring,
	pippyFurniture,
	pippyGames,
	pippyGender,
	pippyInterest,
	pippyName,
	pippyNameColor,
	pippyMusic,
	pippyPersonality,
	pippySpecies,
	pippyStyle,
	pippyWallpaper}

var (
	pippyAstrology     = villagersAstrology{villagerAstrologyGemini}
	pippyBirthDay      = villagersBirthDay{14}
	pippyBirthMonth    = villagersBirthMonth{time.June} // 6
	pippyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	pippyCategory      = villagersCategory{villagerCategoryB}
	pippyClothes       = villagersClothes{} // 9012
	pippyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBrown}}
	pippyFlooring      = villagersFlooring{villagerFlooringBluePaintFlooring}
	pippyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureShowerBooth, villagerFurniturePlainSink, villagerFurnitureWoodenChair, villagerFurnitureToolCart, villagerFurnitureCuteMusicPlayer, villagerFurnitureMiniDIYWorkbench, villagerFurnitureCushion, villagerFurnitureLoftBedWithDesk, villagerFurnitureBlueMediumRoundMat, villagerFurnitureSewingProject, villagerFurnitureLCDTV50In, villagerFurnitureCorkboard, villagerFurnitureWhiteMessageMat}}
	pippyGames         = villagersGames{[]VillagerGame{}} // TBD
	pippyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pippyInterest      = villagersInterest{villagerInterestFashion}
	pippyName          = villagersName{villagerNamePippy}
	pippyNameColor     = villagersNameColor{villagerNameColor879B96}
	pippyMusic         = villagersMusic{villagerMusicDrivin}
	pippyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	pippySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	pippyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	pippyWallpaper     = villagersWallpaper{villagerWallpaperModWall}
)
