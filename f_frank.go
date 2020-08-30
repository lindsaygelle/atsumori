package atsumori

import "time"

// Frank is an Animal Crossing villager.
var Frank = villager{
	frankAstrology,
	frankBirthDay,
	frankBirthMonth,
	frankBubbleColor,
	frankCategory,
	frankClothes,
	frankClothesColors,
	frankFlooring,
	frankFurniture,
	frankGames,
	frankGender,
	frankInterest,
	frankName,
	frankNameColor,
	frankMusic,
	frankPersonality,
	frankSpecies,
	frankStyle,
	frankWallpaper}

var (
	frankAstrology     = villagersAstrology{villagerAstrologyLeo}
	frankBirthDay      = villagersBirthDay{30}
	frankBirthMonth    = villagersBirthMonth{time.July} // 7
	frankBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	frankCategory      = villagersCategory{villagerCategoryB}
	frankClothes       = villagersClothes{villagerClothesLetterJacket} // 2401
	frankClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorYellow}}
	frankFlooring      = villagersFlooring{villagerFlooringBoxingRingMat}
	frankFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShowerBooth, villagerFurnitureWeightBench, villagerFurnitureNeutralCorner, villagerFurniturePunchingBag, villagerFurnitureUprightLocker, villagerFurnitureIronWorktable, villagerFurnitureWaterCooler, villagerFurnitureTapeDeck, villagerFurnitureProteinShakerBottle}}
	frankGames         = villagersGames{[]VillagerGame{}} // TBD
	frankGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	frankInterest      = villagersInterest{villagerInterestEducation}
	frankName          = villagersName{villagerNameFrank}
	frankNameColor     = villagersNameColor{villagerNameColor28665A}
	frankMusic         = villagersMusic{villagerMusicKKCasbah}
	frankPersonality   = villagersPersonality{villagerPersonalityCranky}
	frankSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	frankStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCool}}
	frankWallpaper     = villagersWallpaper{villagerWallpaperGrayShantyWall}
)
