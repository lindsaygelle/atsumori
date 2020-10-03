package atsumori

import "time"

// Ursala is an Animal Crossing villager.
var Ursala = villager{
	ursalaAstrology,
	ursalaBirthDay,
	ursalaBirthMonth,
	ursalaBubbleColor,
	ursalaCategory,
	ursalaClothes,
	ursalaClothesColors,
	ursalaFlooring,
	ursalaFurniture,
	ursalaGames,
	ursalaGender,
	ursalaInterest,
	ursalaName,
	ursalaNameColor,
	ursalaMusic,
	ursalaPersonality,
	ursalaSpecies,
	ursalaStyle,
	ursalaWallpaper}

var (
	ursalaAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	ursalaBirthDay      = villagersBirthDay{16}
	ursalaBirthMonth    = villagersBirthMonth{time.January} // 1
	ursalaBubbleColor   = villagersBubbleColor{villagerBubbleColorF2BDC7}
	ursalaCategory      = villagersCategory{villagerCategoryA}
	ursalaClothes       = villagersClothes{} // 4729
	ursalaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorOrange}}
	ursalaFlooring      = villagersFlooring{villagerFlooringGreenHoneycombTile}
	ursalaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDeluxeWasher, villagerFurnitureYucca, villagerFurnitureIronwoodKitchenette, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenDoubleBed, villagerFurnitureNaturalGardenTable, villagerFurnitureNaturalGardenChair, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureRefrigerator, villagerFurnitureCuteMusicPlayer, villagerFurnitureHangingTerrarium}}
	ursalaGames         = villagersGames{[]VillagerGame{}} // TBD
	ursalaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	ursalaInterest      = villagersInterest{villagerInterestMusic}
	ursalaName          = villagersName{villagerNameUrsala}
	ursalaNameColor     = villagersNameColor{villagerNameColor634B4B}
	ursalaMusic         = villagersMusic{villagerMusicCafeKK}
	ursalaPersonality   = villagersPersonality{villagerPersonalityBigSister}
	ursalaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	ursalaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	ursalaWallpaper     = villagersWallpaper{villagerWallpaperCafeCurtainWall}
)
