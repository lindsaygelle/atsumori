package atsumori

import "time"

// Rolf is an Animal Crossing villager.
var Rolf = villager{
	rolfAstrology,
	rolfBirthDay,
	rolfBirthMonth,
	rolfBubbleColor,
	rolfCategory,
	rolfClothes,
	rolfClothesColors,
	rolfFlooring,
	rolfFurniture,
	rolfGames,
	rolfGender,
	rolfInterest,
	rolfName,
	rolfNameColor,
	rolfMusic,
	rolfPersonality,
	rolfSpecies,
	rolfStyle,
	rolfWallpaper}

var (
	rolfAstrology     = villagersAstrology{villagerAstrologyLeo}
	rolfBirthDay      = villagersBirthDay{22}
	rolfBirthMonth    = villagersBirthMonth{time.August} // 8
	rolfBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	rolfCategory      = villagersCategory{villagerCategoryB}
	rolfClothes       = villagersClothes{villagerClothesDownJacket} // 2778
	rolfClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorBlack}}
	rolfFlooring      = villagersFlooring{villagerFlooringFlowingRiverFlooring}
	rolfFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureOldFashionedWashtub, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureFirewood, villagerFurnitureCampfireCookware, villagerFurnitureSleepingBag, villagerFurnitureStoneStool, villagerFurnitureCampStove, villagerFurnitureLantern, villagerFurnitureWildLogBench, villagerFurniturePortableRadio}}
	rolfGames         = villagersGames{[]VillagerGame{}} // TBD
	rolfGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rolfInterest      = villagersInterest{villagerInterestFitness}
	rolfName          = villagersName{villagerNameRolf}
	rolfNameColor     = villagersNameColor{villagerNameColor848484}
	rolfMusic         = villagersMusic{villagerMusicKKCondor}
	rolfPersonality   = villagersPersonality{villagerPersonalityCranky}
	rolfSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesTiger}}
	rolfStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	rolfWallpaper     = villagersWallpaper{villagerWallpaperSummitWall}
)
