package atsumori

import "time"

// Paula is an Animal Crossing villager.
var Paula = villager{
	paulaAstrology,
	paulaBirthDay,
	paulaBirthMonth,
	paulaBubbleColor,
	paulaCategory,
	paulaClothes,
	paulaClothesColors,
	paulaFlooring,
	paulaFurniture,
	paulaGames,
	paulaGender,
	paulaInterest,
	paulaName,
	paulaNameColor,
	paulaMusic,
	paulaPersonality,
	paulaSpecies,
	paulaStyle,
	paulaWallpaper}

var (
	paulaAstrology     = villagersAstrology{villagerAstrologyAries}
	paulaBirthDay      = villagersBirthDay{22}
	paulaBirthMonth    = villagersBirthMonth{time.March} // 3
	paulaBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	paulaCategory      = villagersCategory{villagerCategoryB}
	paulaClothes       = villagersClothes{} // 4616
	paulaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorGreen}}
	paulaFlooring      = villagersFlooring{villagerFlooringBerryChocolatesFlooring}
	paulaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureDoubleSofa, villagerFurnitureWoodenWasteBin, villagerFurnitureWoodenEndTable, villagerFurniturePortableRecordPlayer, villagerFurnitureWoodenWardrobe, villagerFurnitureWoodenChest, villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenLowTable, villagerFurnitureCorkboard}}
	paulaGames         = villagersGames{[]VillagerGame{}} // TBD
	paulaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	paulaInterest      = villagersInterest{villagerInterestFitness}
	paulaName          = villagersName{villagerNamePaula}
	paulaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	paulaMusic         = villagersMusic{villagerMusicKKDisco}
	paulaPersonality   = villagersPersonality{villagerPersonalityBigSister}
	paulaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	paulaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleActive}}
	paulaWallpaper     = villagersWallpaper{villagerWallpaperStarryWall}
)
