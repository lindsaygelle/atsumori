package atsumori

import "time"

// Boyd is an Animal Crossing villager
var Boyd = villager{
	boydAstrology,
	boydBirthDay,
	boydBirthMonth,
	boydBubbleColor,
	boydCategory,
	boydClothes,
	boydClothesColors,
	boydFlooring,
	boydFurniture,
	boydGames,
	boydGender,
	boydInterest,
	boydName,
	boydNameColor,
	boydMusic,
	boydPersonality,
	boydSpecies,
	boydStyle,
	boydWallpaper}

var (
	boydAstrology     = villagersAstrology{villagerAstrologyLibra}
	boydBirthDay      = villagersBirthDay{1}
	boydBirthMonth    = villagersBirthMonth{time.October} // 10
	boydBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	boydCategory      = villagersCategory{villagerCategoryA}
	boydClothes       = villagersClothes{villagerClothesThreeBallTee} // 2498
	boydClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlack}}
	boydFlooring      = villagersFlooring{villagerFlooringSidewalkFlooring}
	boydFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCone, villagerFurnitureCone, villagerFurnitureOilBarrel, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureOilBarrel, villagerFurnitureConstructionSign, villagerFurnitureOutdoorsyShovel, villagerFurnitureHandcart, villagerFurnitureUtilityPole, villagerFurnitureUtilityPole, villagerFurnitureIronWorktable, villagerFurniturePortableRadio, villagerFurnitureKettle, villagerFurnitureIronFrame, villagerFurnitureIronFrame}}
	boydGames         = villagersGames{[]VillagerGame{}} // TBD
	boydGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	boydInterest      = villagersInterest{villagerInterestFitness}
	boydName          = villagersName{villagerNameBoyd}
	boydNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	boydMusic         = villagersMusic{} // K.K. Groove
	boydPersonality   = villagersPersonality{villagerPersonalityCranky}
	boydSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGorilla}}
	boydStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	boydWallpaper     = villagersWallpaper{villagerWallpaperTreeLinedWall}
)
