package atsumori

import "time"

// BigTop is an Animal Crossing villager.
var BigTop = villager{
	bigTopAstrology,
	bigTopBirthDay,
	bigTopBirthMonth,
	bigTopBubbleColor,
	bigTopCategory,
	bigTopClothes,
	bigTopClothesColors,
	bigTopFlooring,
	bigTopFurniture,
	bigTopGames,
	bigTopGender,
	bigTopInterest,
	bigTopName,
	bigTopNameColor,
	bigTopMusic,
	bigTopPersonality,
	bigTopSpecies,
	bigTopStyle,
	bigTopWallpaper}

var (
	bigTopAstrology     = villagersAstrology{villagerAstrologyLibra}
	bigTopBirthDay      = villagersBirthDay{3}
	bigTopBirthMonth    = villagersBirthMonth{time.October} // 10
	bigTopBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	bigTopCategory      = villagersCategory{villagerCategoryB}
	bigTopClothes       = villagersClothes{villagerClothesNo3Shirt} // 12037
	bigTopClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorGreen}}
	bigTopFlooring      = villagersFlooring{villagerFlooringSteelFlooring}
	bigTopFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurnitureGarbagePail, villagerFurnitureNeutralCorner, villagerFurniturePunchingBag, villagerFurnitureWhiteboard, villagerFurnitureTreadmill, villagerFurnitureCardboardBox, villagerFurnitureTapeDeck, villagerFurnitureBigTopsPhoto}}
	bigTopGames         = villagersGames{[]VillagerGame{}} // TBD
	bigTopGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	bigTopInterest      = villagersInterest{villagerInterestPlay}
	bigTopName          = villagersName{villagerNameBigTop}
	bigTopNameColor     = villagersNameColor{villagerNameColor28665A}
	bigTopMusic         = villagersMusic{villagerMusicGoKKRider}
	bigTopPersonality   = villagersPersonality{villagerPersonalityLazy}
	bigTopSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	bigTopStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	bigTopWallpaper     = villagersWallpaper{villagerWallpaperConstructionSiteWall}
)
