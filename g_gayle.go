package atsumori

import "time"

// Gayle is an Animal Crossing villager
var Gayle = villager{
	gayleAstrology,
	gayleBirthDay,
	gayleBirthMonth,
	gayleBubbleColor,
	gayleCategory,
	gayleClothes,
	gayleClothesColors,
	gayleFlooring,
	gayleFurniture,
	gayleGames,
	gayleGender,
	gayleInterest,
	gayleName,
	gayleNameColor,
	gayleMusic,
	gaylePersonality,
	gayleSpecies,
	gayleStyle,
	gayleWallpaper}

var (
	gayleAstrology     = villagersAstrology{villagerAstrologyTaurus}
	gayleBirthDay      = villagersBirthDay{17}
	gayleBirthMonth    = villagersBirthMonth{time.May} // 5
	gayleBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	gayleCategory      = villagersCategory{villagerCategoryA}
	gayleClothes       = villagersClothes{} // 3070
	gayleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorWhite}}
	gayleFlooring      = villagersFlooring{villagerFlooringPineBoardFlooring}
	gayleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureCuteBed, villagerFurnitureCuteChair, villagerFurnitureCuteVanity, villagerFurnitureCuteSofa, villagerFurnitureFloorLamp, villagerFurnitureWoodenLowTable, villagerFurnitureCuteWallMountedClock, villagerFurnitureMug, villagerFurnitureCuteTeaTable, villagerFurnitureCuteMusicPlayer}}
	gayleGames         = villagersGames{[]VillagerGame{}} // TBD
	gayleGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	gayleInterest      = villagersInterest{villagerInterestNature}
	gayleName          = villagersName{villagerNameGayle}
	gayleNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	gayleMusic         = villagersMusic{villagerClothesBubblegumKK}
	gaylePersonality   = villagersPersonality{villagerPersonalityNormal}
	gayleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAlligator}}
	gayleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	gayleWallpaper     = villagersWallpaper{villagerWallpaperPinkQuiltWall}
)
