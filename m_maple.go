package atsumori

import "time"

// Maple is an Animal Crossing villager.
var Maple = villager{
	mapleAstrology,
	mapleBirthDay,
	mapleBirthMonth,
	mapleBubbleColor,
	mapleCategory,
	mapleClothes,
	mapleClothesColors,
	mapleFlooring,
	mapleFurniture,
	mapleGames,
	mapleGender,
	mapleInterest,
	mapleName,
	mapleNameColor,
	mapleMusic,
	maplePersonality,
	mapleSpecies,
	mapleStyle,
	mapleWallpaper}

var (
	mapleAstrology     = villagersAstrology{villagerAstrologyGemini}
	mapleBirthDay      = villagersBirthDay{15}
	mapleBirthMonth    = villagersBirthMonth{time.June} // 6
	mapleBubbleColor   = villagersBubbleColor{villagerBubbleColorA87850}
	mapleCategory      = villagersCategory{villagerCategoryB}
	mapleClothes       = villagersClothes{villagerClothesTreeSweater} // 8194
	mapleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorGreen}}
	mapleFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	mapleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGasRange, villagerFurnitureWoodenSimpleBed, villagerFurnitureMagazineRack, villagerFurnitureBirdcage, villagerFurnitureWoodBurningStove, villagerFurnitureMomsCushion, villagerFurnitureAnthuriumPlant, villagerFurnitureWoodenEndTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureMiniFridge, villagerFurnitureRevolvingSpiceRack, villagerFurnitureMacrameTapestry, villagerFurnitureRetroDottedRug, villagerFurnitureWoodenLowTable, villagerFurnitureMagazine, villagerFurnitureMug, villagerFurniturePotRack, villagerFurnitureBabyBear}}
	mapleGames         = villagersGames{[]VillagerGame{}} // TBD
	mapleGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	mapleInterest      = villagersInterest{villagerInterestEducation}
	mapleName          = villagersName{villagerNameMaple}
	mapleNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	mapleMusic         = villagersMusic{villagerMusicForestLife}
	maplePersonality   = villagersPersonality{villagerPersonalityNormal}
	mapleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	mapleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	mapleWallpaper     = villagersWallpaper{villagerWallpaperGreenBlossomingWall}
)
