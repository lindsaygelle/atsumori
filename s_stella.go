package atsumori

import "time"

// Stella is an Animal Crossing villager.
var Stella = villager{
	stellaAstrology,
	stellaBirthDay,
	stellaBirthMonth,
	stellaBubbleColor,
	stellaCategory,
	stellaClothes,
	stellaClothesColors,
	stellaFlooring,
	stellaFurniture,
	stellaGames,
	stellaGender,
	stellaInterest,
	stellaName,
	stellaNameColor,
	stellaMusic,
	stellaPersonality,
	stellaSpecies,
	stellaStyle,
	stellaWallpaper}

var (
	stellaAstrology     = villagersAstrology{villagerAstrologyAries}
	stellaBirthDay      = villagersBirthDay{9}
	stellaBirthMonth    = villagersBirthMonth{time.April} // 4
	stellaBubbleColor   = villagersBubbleColor{villagerBubbleColor951D43}
	stellaCategory      = villagersCategory{villagerCategoryA}
	stellaClothes       = villagersClothes{} // 3597
	stellaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorBeige}}
	stellaFlooring      = villagersFlooring{villagerFlooringDarkHerringboneFlooring}
	stellaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePlainSink, villagerFurnitureWoodenSimpleBed, villagerFurnitureCuteDIYTable, villagerFurnitureWoodenChair, villagerFurnitureYucca, villagerFurnitureBathroomTowelRack, villagerFurnitureWoodenChest, villagerFurnitureSturdySewingBox, villagerFurnitureWoodenTableMirror, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenTable, villagerFurnitureWoodenWasteBin, villagerFurnitureCorkboard, villagerFurnitureMomsTissueBox, villagerFurnitureSewingProject, villagerFurnitureCuteMusicPlayer}}
	stellaGames         = villagersGames{[]VillagerGame{}} // TBD
	stellaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	stellaInterest      = villagersInterest{villagerInterestNature}
	stellaName          = villagersName{villagerNameStella}
	stellaNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	stellaMusic         = villagersMusic{villagerMusicOnlyMe}
	stellaPersonality   = villagersPersonality{villagerPersonalityNormal}
	stellaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	stellaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	stellaWallpaper     = villagersWallpaper{villagerWallpaperRoseWall}
)
