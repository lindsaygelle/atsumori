package atsumori

import "time"

// Broccolo is an Animal Crossing villager.
var Broccolo = villager{
	broccoloAstrology,
	broccoloBirthDay,
	broccoloBirthMonth,
	broccoloBubbleColor,
	broccoloCategory,
	broccoloClothes,
	broccoloClothesColors,
	broccoloFlooring,
	broccoloFurniture,
	broccoloGames,
	broccoloGender,
	broccoloInterest,
	broccoloName,
	broccoloNameColor,
	broccoloMusic,
	broccoloPersonality,
	broccoloSpecies,
	broccoloStyle,
	broccoloWallpaper}

var (
	broccoloAstrology     = villagersAstrology{villagerAstrologyCancer}
	broccoloBirthDay      = villagersBirthDay{30}
	broccoloBirthMonth    = villagersBirthMonth{time.June} // 6
	broccoloBubbleColor   = villagersBubbleColor{villagerBubbleColor55CFFF}
	broccoloCategory      = villagersCategory{villagerCategoryB}
	broccoloClothes       = villagersClothes{villagerClothesRaglanShirt} // 2674
	broccoloClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorYellow}}
	broccoloFlooring      = villagersFlooring{villagerFlooringPastelPuzzleFlooring}
	broccoloFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenLowTable, villagerFurnitureBlueDottedRug, villagerFurniturePaintingSet, villagerFurnitureWoodenBlockBench, villagerFurnitureToyBox, villagerFurnitureWoodenBlockstool, villagerFurnitureWoodenBlockWallClock, villagerFurnitureMomsPlushie, villagerFurnitureWoodenBlockChest, villagerFurnitureElephantSlide, villagerFurnitureMomsCushion, villagerFurnitureCuteMusicPlayer}}
	broccoloGames         = villagersGames{[]VillagerGame{}} // TBD
	broccoloGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	broccoloInterest      = villagersInterest{villagerInterestPlay}
	broccoloName          = villagersName{villagerNameBroccolo}
	broccoloNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	broccoloMusic         = villagersMusic{villagerMusicILoveYou}
	broccoloPersonality   = villagersPersonality{villagerPersonalityLazy}
	broccoloSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	broccoloStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	broccoloWallpaper     = villagersWallpaper{villagerWallpaperBluePlayroomWall}
)
