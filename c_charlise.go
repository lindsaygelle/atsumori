package atsumori

import "time"

// Charlise is an Animal Crossing villager
var Charlise = villager{
	charliseAstrology,
	charliseBirthDay,
	charliseBirthMonth,
	charliseBubbleColor,
	charliseCategory,
	charliseClothes,
	charliseClothesColors,
	charliseFlooring,
	charliseFurniture,
	charliseGames,
	charliseGender,
	charliseInterest,
	charliseName,
	charliseNameColor,
	charliseMusic,
	charlisePersonality,
	charliseSpecies,
	charliseStyle,
	charliseWallpaper}

var (
	charliseAstrology     = villagersAstrology{villagerAstrologyAries}
	charliseBirthDay      = villagersBirthDay{17}
	charliseBirthMonth    = villagersBirthMonth{time.April} // 4
	charliseBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	charliseCategory      = villagersCategory{villagerCategoryB}
	charliseClothes       = villagersClothes{villagerClothesAthleticJacket} // 3244
	charliseClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorOrange}}
	charliseFlooring      = villagersFlooring{villagerFlooringRacetrackFlooring}
	charliseFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBall, villagerFurnitureExerciseBall, villagerFurnitureCampingCot, villagerFurnitureBall, villagerFurnitureOutdoorBench, villagerFurnitureTapeDeck, villagerFurnitureBasketballHoop, villagerFurnitureOutdoorTable, villagerFurnitureHandyWaterCooler, villagerFurnitureMug}}
	charliseGames         = villagersGames{[]VillagerGame{}} // TBD
	charliseGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	charliseInterest      = villagersInterest{villagerInterestFitness}
	charliseName          = villagersName{villagerNameCharlise}
	charliseNameColor     = villagersNameColor{villagerNameColor28665A}
	charliseMusic         = villagersMusic{} // Mr. K.K.
	charlisePersonality   = villagersPersonality{villagerPersonalityBigSister}
	charliseSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	charliseStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCute}}
	charliseWallpaper     = villagersWallpaper{villagerWallpaperTreeLinedWall}
)
