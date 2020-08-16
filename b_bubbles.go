package atsumori

import "time"

// Bubbles is an Animal Crossing villager
var Bubbles = villager{
	bubblesAstrology,
	bubblesBirthDay,
	bubblesBirthMonth,
	bubblesBubbleColor,
	bubblesCategory,
	bubblesClothes,
	bubblesClothesColors,
	bubblesFlooring,
	bubblesFurniture,
	bubblesGames,
	bubblesGender,
	bubblesInterest,
	bubblesName,
	bubblesNameColor,
	bubblesMusic,
	bubblesPersonality,
	bubblesSpecies,
	bubblesStyle,
	bubblesWallpaper}

var (
	bubblesAstrology     = villagersAstrology{villagerAstrologyVirgo}
	bubblesBirthDay      = villagersBirthDay{18}
	bubblesBirthMonth    = villagersBirthMonth{time.September} // 9
	bubblesBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	bubblesCategory      = villagersCategory{villagerCategoryB}
	bubblesClothes       = villagersClothes{} // 3290
	bubblesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorPink}}
	bubblesFlooring      = villagersFlooring{villagerFlooringPineBoardFlooring}
	bubblesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureKettlebell, villagerFurnitureBasketballHoop, villagerFurnitureOutdoorTable, villagerFurnitureHandyWaterCooler, villagerFurnitureCuteMusicPlayer, villagerFurnitureFoldingChair, villagerFurnitureKettlebell, villagerFurnitureChalkboard, villagerFurnitureBall, villagerFurnitureWallClock, villagerFurnitureKeyHolder, villagerFurnitureChampionsPennant, villagerFurnitureExitSign}}
	bubblesGames         = villagersGames{[]VillagerGame{}} // TBD
	bubblesGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	bubblesInterest      = villagersInterest{villagerInterestFashion}
	bubblesName          = villagersName{villagerNameBubbles}
	bubblesNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bubblesMusic         = villagersMusic{} // K.K. Stroll
	bubblesPersonality   = villagersPersonality{villagerPersonalityPeppy}
	bubblesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHippo}}
	bubblesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCute}}
	bubblesWallpaper     = villagersWallpaper{villagerWallpaperPerforatedBoardWall}
)
