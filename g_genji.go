package atsumori

import "time"

// Genji is an Animal Crossing villager
var Genji = villager{
	genjiAstrology,
	genjiBirthDay,
	genjiBirthMonth,
	genjiBubbleColor,
	genjiCategory,
	genjiClothes,
	genjiClothesColors,
	genjiFlooring,
	genjiFurniture,
	genjiGames,
	genjiGender,
	genjiInterest,
	genjiName,
	genjiNameColor,
	genjiMusic,
	genjiPersonality,
	genjiSpecies,
	genjiStyle,
	genjiWallpaper}

var (
	genjiAstrology     = villagersAstrology{villagerAstrologyAquarius}
	genjiBirthDay      = villagersBirthDay{21}
	genjiBirthMonth    = villagersBirthMonth{time.January} // 1
	genjiBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	genjiCategory      = villagersCategory{villagerCategoryB}
	genjiClothes       = villagersClothes{villagerClothesMistyTee} // 8201
	genjiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorPurple}}
	genjiFlooring      = villagersFlooring{villagerFlooringMossyGardenFlooring}
	genjiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBonsaiShelf, villagerFurniturePondStone, villagerFurnitureCherryBlossomBranches, villagerFurnitureScreen, villagerFurnitureBambooBench, villagerFurnitureHangingScroll, villagerFurnitureBambooSphere, villagerFurnitureDeerScare, villagerFurnitureTatamiBed, villagerFurniturePaperLantern}}
	genjiGames         = villagersGames{[]VillagerGame{}} // TBD
	genjiGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	genjiInterest      = villagersInterest{villagerInterestFitness}
	genjiName          = villagersName{villagerNameGenji}
	genjiNameColor     = villagersNameColor{villagerNameColor848484}
	genjiMusic         = villagersMusic{villagerMusicKingKK}
	genjiPersonality   = villagersPersonality{villagerPersonalityJock}
	genjiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	genjiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleSimple}}
	genjiWallpaper     = villagersWallpaper{villagerWallpaperGoldScreenWall}
)
