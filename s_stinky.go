package atsumori

import "time"

// Stinky is an Animal Crossing villager.
var Stinky = villager{
	stinkyAstrology,
	stinkyBirthDay,
	stinkyBirthMonth,
	stinkyBubbleColor,
	stinkyCategory,
	stinkyClothes,
	stinkyClothesColors,
	stinkyFlooring,
	stinkyFurniture,
	stinkyGames,
	stinkyGender,
	stinkyInterest,
	stinkyName,
	stinkyNameColor,
	stinkyMusic,
	stinkyPersonality,
	stinkySpecies,
	stinkyStyle,
	stinkyWallpaper}

var (
	stinkyAstrology     = villagersAstrology{villagerAstrologyLeo}
	stinkyBirthDay      = villagersBirthDay{17}
	stinkyBirthMonth    = villagersBirthMonth{time.August} // 8
	stinkyBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	stinkyCategory      = villagersCategory{villagerCategoryB}
	stinkyClothes       = villagersClothes{villagerClothesAthleticJacket} // 3244
	stinkyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlue}}
	stinkyFlooring      = villagersFlooring{villagerFlooringBoxingRingMat}
	stinkyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurnitureCassettePlayer, villagerFurniturePunchingBag, villagerFurnitureBarbell, villagerFurnitureCardboardBox, villagerFurnitureProteinShakerBottle, villagerFurnitureBlueCorner, villagerFurnitureIronWorktable, villagerFurnitureThrowbackWrestlingFigure, villagerFurnitureJudgesBell}}
	stinkyGames         = villagersGames{[]VillagerGame{}} // TBD
	stinkyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	stinkyInterest      = villagersInterest{villagerInterestFitness}
	stinkyName          = villagersName{villagerNameStinky}
	stinkyNameColor     = villagersNameColor{villagerNameColor8B5F57}
	stinkyMusic         = villagersMusic{villagerMusicKKFlamenco}
	stinkyPersonality   = villagersPersonality{villagerPersonalityJock}
	stinkySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	stinkyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	stinkyWallpaper     = villagersWallpaper{villagerWallpaperRingsideSeating}
)
