package atsumori

import "time"

// Merry is an Animal Crossing villager.
var Merry = villager{
	merryAstrology,
	merryBirthDay,
	merryBirthMonth,
	merryBubbleColor,
	merryCategory,
	merryClothes,
	merryClothesColors,
	merryFlooring,
	merryFurniture,
	merryGames,
	merryGender,
	merryInterest,
	merryName,
	merryNameColor,
	merryMusic,
	merryPersonality,
	merrySpecies,
	merryStyle,
	merryWallpaper}

var (
	merryAstrology     = villagersAstrology{villagerAstrologyCancer}
	merryBirthDay      = villagersBirthDay{29}
	merryBirthMonth    = villagersBirthMonth{time.June} // 6
	merryBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	merryCategory      = villagersCategory{villagerCategoryB}
	merryClothes       = villagersClothes{villagerClothesDreamySweater} // 3635
	merryClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorLightBlue}}
	merryFlooring      = villagersFlooring{villagerFlooringOrangeRetroFlooring}
	merryFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureYucca, villagerFurnitureDeluxeWasher, villagerFurnitureDoubleSofa, villagerFurnitureRefrigerator, villagerFurnitureHiFiStereo, villagerFurnitureSystemKitchen, villagerFurnitureWoodenTable, villagerFurnitureWoodenChair, villagerFurnitureCoconutWallPlanter}}
	merryGames         = villagersGames{[]VillagerGame{}} // TBD
	merryGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	merryInterest      = villagersInterest{villagerInterestFashion}
	merryName          = villagersName{villagerNameMerry}
	merryNameColor     = villagersNameColor{villagerNameColor879B96}
	merryMusic         = villagersMusic{villagerMusicNeapolitan}
	merryPersonality   = villagersPersonality{villagerPersonalityPeppy}
	merrySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	merryStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	merryWallpaper     = villagersWallpaper{villagerWallpaperBeadedCurtainWall}
)
