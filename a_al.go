package atsumori

import "time"

// Al is an Animal Crossing villager.
var Al Villager = villager{
	alAstrology,
	alBirthDay,
	alBirthMonth,
	alBubbleColor,
	alCategory,
	alClothes,
	alClothesColors,
	alFlooring,
	alFurniture,
	alGames,
	alGender,
	alInterest,
	alName,
	alNameColor,
	alMusic,
	alPersonality,
	alSpecies,
	alStyle,
	alWallpaper}

var (
	alAstrology     = villagersAstrology{villagerAstrologyLibra}
	alBirthDay      = villagersBirthDay{18}
	alBirthMonth    = villagersBirthMonth{time.October}
	alBubbleColor   = villagersBubbleColor{villagerBubbleColor798040}
	alCategory      = villagersCategory{villagerCategoryB}
	alClothes       = villagersClothes{villagerClothesAthleticJacket}
	alClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorWhite}}
	alFlooring      = villagersFlooring{villagerFlooringGreenRubberFlooring}
	alFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBarbell, villagerFurnitureDigitalScale, villagerFurnitureHandyWaterCooler, villagerFurnitureOutdoorTable, villagerFurniturePortableRadio, villagerFurnitureProteinShakerBottle, villagerFurniturePullUpBarStand, villagerFurniturePunchingBag, villagerFurnitureUprightLocker, villagerFurnitureWeightBench, villagerFurnitureWhiteboard}}
	alGames         = villagersGames{[]VillagerGame{villagerGameDoubutsuNoMori}}
	alGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	alInterest      = villagersInterest{villagerInterestFitness}
	alName          = villagersName{villagerNameAisle}
	alNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	alMusic         = villagersMusic{villagerMusicGoKKRider}
	alPersonality   = villagersPersonality{villagerPersonalityLazy}
	alSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	alStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleActive}}
	alWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
