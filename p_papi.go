package atsumori

import "time"

// Papi is an Animal Crossing villager.
var Papi = villager{
	papiAstrology,
	papiBirthDay,
	papiBirthMonth,
	papiBubbleColor,
	papiCategory,
	papiClothes,
	papiClothesColors,
	papiFlooring,
	papiFurniture,
	papiGames,
	papiGender,
	papiInterest,
	papiName,
	papiNameColor,
	papiMusic,
	papiPersonality,
	papiSpecies,
	papiStyle,
	papiWallpaper}

var (
	papiAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	papiBirthDay      = villagersBirthDay{10}
	papiBirthMonth    = villagersBirthMonth{time.January} // 1
	papiBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	papiCategory      = villagersCategory{villagerCategoryA}
	papiClothes       = villagersClothes{villagerClothesColorfulStripedSweater} // 4613
	papiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorLightBlue}}
	papiFlooring      = villagersFlooring{villagerFlooringJointedMatFlooring}
	papiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureYucca, villagerFurnitureWoodenBlockBed, villagerFurnitureWoodenBlockstool, villagerFurnitureDalaHorse, villagerFurnitureWoodenBlockWallClock, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureWoodenBlockStereo, villagerFurnitureWoodenBlockTable, villagerFurnitureWoodenBlockToy, villagerFurnitureRockingHorse, villagerFurnitureWoodenBlockChest}}
	papiGames         = villagersGames{[]VillagerGame{}} // TBD
	papiGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	papiInterest      = villagersInterest{villagerInterestNature}
	papiName          = villagersName{villagerNamePapi}
	papiNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	papiMusic         = villagersMusic{villagerMusicHypnoKK}
	papiPersonality   = villagersPersonality{villagerPersonalityLazy}
	papiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	papiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	papiWallpaper     = villagersWallpaper{villagerWallpaperColorfulPuzzleWall}
)
