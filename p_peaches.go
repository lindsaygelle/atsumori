package atsumori

import "time"

// Peaches is an Animal Crossing villager.
var Peaches = villager{
	peachesAstrology,
	peachesBirthDay,
	peachesBirthMonth,
	peachesBubbleColor,
	peachesCategory,
	peachesClothes,
	peachesClothesColors,
	peachesFlooring,
	peachesFurniture,
	peachesGames,
	peachesGender,
	peachesInterest,
	peachesName,
	peachesNameColor,
	peachesMusic,
	peachesPersonality,
	peachesSpecies,
	peachesStyle,
	peachesWallpaper}

var (
	peachesAstrology     = villagersAstrology{villagerAstrologySagittarius}
	peachesBirthDay      = villagersBirthDay{28}
	peachesBirthMonth    = villagersBirthMonth{time.November} // 11
	peachesBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	peachesCategory      = villagersCategory{villagerCategoryB}
	peachesClothes       = villagersClothes{} // 4555
	peachesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorLightBlue}}
	peachesFlooring      = villagersFlooring{villagerFlooringDirtFlooring}
	peachesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHayBed, villagerFurnitureFirewood, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureWoodBurningStove, villagerFurnitureGasRange, villagerFurnitureLogDecorativeShelves, villagerFurnitureLogStakes, villagerFurnitureLogStool, villagerFurniturePotRack, villagerFurnitureCuteMusicPlayer, villagerFurnitureLogDiningTable, villagerFurniturePennant}}
	peachesGames         = villagersGames{[]VillagerGame{}} // TBD
	peachesGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	peachesInterest      = villagersInterest{villagerInterestEducation}
	peachesName          = villagersName{villagerNamePeaches}
	peachesNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	peachesMusic         = villagersMusic{villagerMusicForestLife}
	peachesPersonality   = villagersPersonality{villagerPersonalityNormal}
	peachesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	peachesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	peachesWallpaper     = villagersWallpaper{villagerWallpaperCabinWall}
)
