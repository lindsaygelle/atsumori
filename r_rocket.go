package atsumori

import "time"

// Rocket is an Animal Crossing villager.
var Rocket = villager{
	rocketAstrology,
	rocketBirthDay,
	rocketBirthMonth,
	rocketBubbleColor,
	rocketCategory,
	rocketClothes,
	rocketClothesColors,
	rocketFlooring,
	rocketFurniture,
	rocketGames,
	rocketGender,
	rocketInterest,
	rocketName,
	rocketNameColor,
	rocketMusic,
	rocketPersonality,
	rocketSpecies,
	rocketStyle,
	rocketWallpaper}

var (
	rocketAstrology     = villagersAstrology{villagerAstrologyAries}
	rocketBirthDay      = villagersBirthDay{14}
	rocketBirthMonth    = villagersBirthMonth{time.April} // 4
	rocketBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	rocketCategory      = villagersCategory{villagerCategoryB}
	rocketClothes       = villagersClothes{villagerClothesNo4Shirt} // 12038
	rocketClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorRed}}
	rocketFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	rocketFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePullUpBarStand, villagerFurniturePunchingBag, villagerFurnitureRedCarpet, villagerFurnitureRedCarpet, villagerFurnitureRoseBed, villagerFurnitureDIYWorkbench, villagerFurnitureExerciseBall, villagerFurnitureWoodenMiniTable, villagerFurnitureKettlebell, villagerFurnitureToolCart, villagerFurniturePortableRecordPlayer, villagerFurnitureExerciseBike, villagerFurnitureRocketsPhoto}}
	rocketGames         = villagersGames{[]VillagerGame{}} // TBD
	rocketGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	rocketInterest      = villagersInterest{villagerInterestFitness}
	rocketName          = villagersName{villagerNameRocket}
	rocketNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	rocketMusic         = villagersMusic{villagerMusicKKAdventure}
	rocketPersonality   = villagersPersonality{villagerPersonalityBigSister}
	rocketSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGorilla}}
	rocketStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	rocketWallpaper     = villagersWallpaper{villagerWallpaperSkyscraperWall}
)
