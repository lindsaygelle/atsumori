package atsumori

import "time"

// Lobo is an Animal Crossing villager.
var Lobo = villager{
	loboAstrology,
	loboBirthDay,
	loboBirthMonth,
	loboBubbleColor,
	loboCategory,
	loboClothes,
	loboClothesColors,
	loboFlooring,
	loboFurniture,
	loboGames,
	loboGender,
	loboInterest,
	loboName,
	loboNameColor,
	loboMusic,
	loboPersonality,
	loboSpecies,
	loboStyle,
	loboWallpaper}

var (
	loboAstrology     = villagersAstrology{villagerAstrologyScorpio}
	loboBirthDay      = villagersBirthDay{5}
	loboBirthMonth    = villagersBirthMonth{time.November} // 11
	loboBubbleColor   = villagersBubbleColor{villagerBubbleColor7352E8}
	loboCategory      = villagersCategory{villagerCategoryB}
	loboClothes       = villagersClothes{villagerClothesBomberStyleJacket} // 4401
	loboClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorBeige}}
	loboFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	loboFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHighEndStereo, villagerFurnitureLogWallMountedClock, villagerFurnitureWallMountedToolBoard, villagerFurnitureLogBed, villagerFurnitureLogExtraLongSofa, villagerFurnitureCoolerBox, villagerFurnitureLogDecorativeShelves, villagerFurnitureLogRoundTable, villagerFurnitureMacrameTapestry, villagerFurnitureChessboard, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureKeyHolder, villagerFurnitureBroomAndDustpan}}
	loboGames         = villagersGames{[]VillagerGame{}} // TBD
	loboGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	loboInterest      = villagersInterest{villagerInterestEducation}
	loboName          = villagersName{villagerNameLobo}
	loboNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	loboMusic         = villagersMusic{villagerMusicKKRock}
	loboPersonality   = villagersPersonality{villagerPersonalityCranky}
	loboSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	loboStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	loboWallpaper     = villagersWallpaper{villagerWallpaperRusticStoneWall}
)
