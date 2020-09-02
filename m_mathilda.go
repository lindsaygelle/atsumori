package atsumori

import "time"

// Mathilda is an Animal Crossing villager.
var Mathilda = villager{
	mathildaAstrology,
	mathildaBirthDay,
	mathildaBirthMonth,
	mathildaBubbleColor,
	mathildaCategory,
	mathildaClothes,
	mathildaClothesColors,
	mathildaFlooring,
	mathildaFurniture,
	mathildaGames,
	mathildaGender,
	mathildaInterest,
	mathildaName,
	mathildaNameColor,
	mathildaMusic,
	mathildaPersonality,
	mathildaSpecies,
	mathildaStyle,
	mathildaWallpaper}

var (
	mathildaAstrology     = villagersAstrology{villagerAstrologyScorpio}
	mathildaBirthDay      = villagersBirthDay{12}
	mathildaBirthMonth    = villagersBirthMonth{time.November} // 11
	mathildaBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	mathildaCategory      = villagersCategory{villagerCategoryB}
	mathildaClothes       = villagersClothes{} // 3261
	mathildaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorRed}}
	mathildaFlooring      = villagersFlooring{villagerFlooringArabesqueFlooring}
	mathildaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureOpenFrameKitchen, villagerFurnitureAntiqueBed, villagerFurnitureRockingHorse, villagerFurnitureWoodenBlockstool, villagerFurnitureToyBox, villagerFurnitureAntiqueTable, villagerFurnitureCuckooClock, villagerFurnitureWoodenBlockToy, villagerFurnitureAntiqueMiniTable, villagerFurnitureAntiqueConsoleTable, villagerFurnitureTableLamp, villagerFurniturePhonograph, villagerFurnitureNutcracker}}
	mathildaGames         = villagersGames{[]VillagerGame{}} // TBD
	mathildaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	mathildaInterest      = villagersInterest{villagerInterestFitness}
	mathildaName          = villagersName{villagerNameMathilda}
	mathildaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	mathildaMusic         = villagersMusic{villagerMusicKKChorale}
	mathildaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	mathildaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKangaroo}}
	mathildaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	mathildaWallpaper     = villagersWallpaper{villagerWallpaperBlueDelicateBloomsWall}
)
