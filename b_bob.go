package atsumori

import "time"

// Bob is an Animal Crossing villager
var Bob = villager{
	bobAstrology,
	bobBirthDay,
	bobBirthMonth,
	bobBubbleColor,
	bobCategory,
	bobClothes,
	bobClothesColors,
	bobFlooring,
	bobFurniture,
	bobGames,
	bobGender,
	bobInterest,
	bobName,
	bobNameColor,
	bobMusic,
	bobPersonality,
	bobSpecies,
	bobStyle,
	bobWallpaper}

var (
	bobAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	bobBirthDay      = villagersBirthDay{1}
	bobBirthMonth    = villagersBirthMonth{time.January} // 1
	bobBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	bobCategory      = villagersCategory{villagerCategoryB}
	bobClothes       = villagersClothes{} // 8205
	bobClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorRed}}
	bobFlooring      = villagersFlooring{villagerFlooringColorfulPuzzleFlooring}
	bobFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBlockStereo, villagerFurnitureWoodenBlockBed, villagerFurnitureWoodenBlockChair, villagerFurnitureWoodenBlockTable, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureWoodenBlockBench, villagerFurnitureWoodenBlockChest, villagerFurnitureWoodenBlockToy, villagerFurnitureToyBox, villagerFurnitureWoodenBlockWallClock}}
	bobGames         = villagersGames{[]VillagerGame{}} // TBD
	bobGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	bobInterest      = villagersInterest{villagerInterestPlay}
	bobName          = villagersName{villagerNameBob}
	bobNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bobMusic         = villagersMusic{villagerMusicNeapolitan}
	bobPersonality   = villagersPersonality{villagerPersonalityLazy}
	bobSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	bobStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	bobWallpaper     = villagersWallpaper{villagerWallpaperBluePlayroomWall}
)
