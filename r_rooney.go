package atsumori

import "time"

// Rooney is an Animal Crossing villager.
var Rooney = villager{
	rooneyAstrology,
	rooneyBirthDay,
	rooneyBirthMonth,
	rooneyBubbleColor,
	rooneyCategory,
	rooneyClothes,
	rooneyClothesColors,
	rooneyFlooring,
	rooneyFurniture,
	rooneyGames,
	rooneyGender,
	rooneyInterest,
	rooneyName,
	rooneyNameColor,
	rooneyMusic,
	rooneyPersonality,
	rooneySpecies,
	rooneyStyle,
	rooneyWallpaper}

var (
	rooneyAstrology     = villagersAstrology{villagerAstrologySagittarius}
	rooneyBirthDay      = villagersBirthDay{1}
	rooneyBirthMonth    = villagersBirthMonth{time.December} // 12
	rooneyBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	rooneyCategory      = villagersCategory{villagerCategoryA}
	rooneyClothes       = villagersClothes{villagerClothesSleevelessParka} // 7994
	rooneyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorBlack}}
	rooneyFlooring      = villagersFlooring{villagerFlooringSteelFlooring}
	rooneyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurnitureTreadmill, villagerFurnitureNeutralCorner, villagerFurniturePunchingBag, villagerFurnitureSpeedBag, villagerFurnitureIronWorktable, villagerFurnitureHandyWaterCooler, villagerFurnitureWallFan, villagerFurnitureIronWorktable, villagerFurnitureJudgesBell, villagerFurnitureCuteMusicPlayer, villagerFurnitureProteinShakerBottle}}
	rooneyGames         = villagersGames{[]VillagerGame{}} // TBD
	rooneyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rooneyInterest      = villagersInterest{villagerInterestFitness}
	rooneyName          = villagersName{villagerNameRooney}
	rooneyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	rooneyMusic         = villagersMusic{villagerMusicKKAdventure}
	rooneyPersonality   = villagersPersonality{villagerPersonalityCranky}
	rooneySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKangaroo}}
	rooneyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleActive}}
	rooneyWallpaper     = villagersWallpaper{villagerWallpaperGrayShantyWall}
)
