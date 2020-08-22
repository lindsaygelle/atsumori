package atsumori

import "time"

// Drift is an Animal Crossing villager
var Drift = villager{
	driftAstrology,
	driftBirthDay,
	driftBirthMonth,
	driftBubbleColor,
	driftCategory,
	driftClothes,
	driftClothesColors,
	driftFlooring,
	driftFurniture,
	driftGames,
	driftGender,
	driftInterest,
	driftName,
	driftNameColor,
	driftMusic,
	driftPersonality,
	driftSpecies,
	driftStyle,
	driftWallpaper}

var (
	driftAstrology     = villagersAstrology{villagerAstrologyLibra}
	driftBirthDay      = villagersBirthDay{9}
	driftBirthMonth    = villagersBirthMonth{time.October} // 10
	driftBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	driftCategory      = villagersCategory{villagerCategoryB}
	driftClothes       = villagersClothes{villagerClothesSixBallTee} // 3323
	driftClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorRed}}
	driftFlooring      = villagersFlooring{villagerFlooringSkullPrintFlooring}
	driftFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDJsTurntable, villagerFurnitureIronShelf, villagerFurnitureWallMountedToolBoard, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureIronwoodBed, villagerFurnitureIronwoodLowTable, villagerFurnitureIronwoodDresser, villagerFurnitureSucculentPlant, villagerFurnitureThrowbackSkullRadio, villagerFurnitureFanPalm}}
	driftGames         = villagersGames{[]VillagerGame{}} // TBD
	driftGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	driftInterest      = villagersInterest{villagerInterestFitness}
	driftName          = villagersName{villagerNameDrift}
	driftNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	driftMusic         = villagersMusic{villagerMusicKKMetal}
	driftPersonality   = villagersPersonality{villagerPersonalityJock}
	driftSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	driftStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	driftWallpaper     = villagersWallpaper{villagerWallpaperGrayShantyWall}
)
