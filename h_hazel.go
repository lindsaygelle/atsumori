package atsumori

import "time"

// Hazel is an Animal Crossing villager.
var Hazel = villager{
	hazelAstrology,
	hazelBirthDay,
	hazelBirthMonth,
	hazelBubbleColor,
	hazelCategory,
	hazelClothes,
	hazelClothesColors,
	hazelFlooring,
	hazelFurniture,
	hazelGames,
	hazelGender,
	hazelInterest,
	hazelName,
	hazelNameColor,
	hazelMusic,
	hazelPersonality,
	hazelSpecies,
	hazelStyle,
	hazelWallpaper}

var (
	hazelAstrology     = villagersAstrology{villagerAstrologyVirgo}
	hazelBirthDay      = villagersBirthDay{30}
	hazelBirthMonth    = villagersBirthMonth{time.August} // 8
	hazelBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	hazelCategory      = villagersCategory{villagerCategoryA}
	hazelClothes       = villagersClothes{villagerClothesAthleticJacket} // 3244
	hazelClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorYellow}}
	hazelFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	hazelFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFoldingChair, villagerFurnitureTennisTable, villagerFurnitureTennisTable, villagerFurnitureBasketballHoop, villagerFurnitureBall, villagerFurnitureTapeDeck, villagerFurnitureWallClock, villagerFurnitureChampionsPennant}}
	hazelGames         = villagersGames{[]VillagerGame{}} // TBD
	hazelGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	hazelInterest      = villagersInterest{villagerInterestPlay}
	hazelName          = villagersName{villagerNameHazel}
	hazelNameColor     = villagersNameColor{villagerNameColor874C25}
	hazelMusic         = villagersMusic{villagerMusicDJKK}
	hazelPersonality   = villagersPersonality{villagerPersonalityBigSister}
	hazelSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	hazelStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCute}}
	hazelWallpaper     = villagersWallpaper{villagerWallpaperDojoWall}
)
