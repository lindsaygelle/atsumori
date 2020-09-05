package atsumori

import "time"

// Mitzi is an Animal Crossing villager.
var Mitzi = villager{
	mitziAstrology,
	mitziBirthDay,
	mitziBirthMonth,
	mitziBubbleColor,
	mitziCategory,
	mitziClothes,
	mitziClothesColors,
	mitziFlooring,
	mitziFurniture,
	mitziGames,
	mitziGender,
	mitziInterest,
	mitziName,
	mitziNameColor,
	mitziMusic,
	mitziPersonality,
	mitziSpecies,
	mitziStyle,
	mitziWallpaper}

var (
	mitziAstrology     = villagersAstrology{villagerAstrologyLibra}
	mitziBirthDay      = villagersBirthDay{25}
	mitziBirthMonth    = villagersBirthMonth{time.September} // 9
	mitziBubbleColor   = villagersBubbleColor{villagerBubbleColor7FA9FF}
	mitziCategory      = villagersCategory{villagerCategoryB}
	mitziClothes       = villagersClothes{} // 2686
	mitziClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorBeige}}
	mitziFlooring      = villagersFlooring{villagerFlooringLightHerringboneFlooring}
	mitziFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenChest, villagerFurnitureGasRange, villagerFurnitureWoodenSimpleBed, villagerFurnitureRefrigerator, villagerFurnitureWoodenLowTable, villagerFurnitureIronWallRack, villagerFurnitureWoodenMiniTable, villagerFurnitureCatGrass, villagerFurnitureFanPalm}}
	mitziGames         = villagersGames{[]VillagerGame{}} // TBD
	mitziGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	mitziInterest      = villagersInterest{villagerInterestEducation}
	mitziName          = villagersName{villagerNameMitzi}
	mitziNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	mitziMusic         = villagersMusic{villagerMusicKKLoveSong}
	mitziPersonality   = villagersPersonality{villagerPersonalityNormal}
	mitziSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	mitziStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	mitziWallpaper     = villagersWallpaper{villagerWallpaperBluePaintedWoodWall}
)
