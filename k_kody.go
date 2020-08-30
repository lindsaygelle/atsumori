package atsumori

import "time"

// Kody is an Animal Crossing villager.
var Kody = villager{
	kodyAstrology,
	kodyBirthDay,
	kodyBirthMonth,
	kodyBubbleColor,
	kodyCategory,
	kodyClothes,
	kodyClothesColors,
	kodyFlooring,
	kodyFurniture,
	kodyGames,
	kodyGender,
	kodyInterest,
	kodyName,
	kodyNameColor,
	kodyMusic,
	kodyPersonality,
	kodySpecies,
	kodyStyle,
	kodyWallpaper}

var (
	kodyAstrology     = villagersAstrology{villagerAstrologyLibra}
	kodyBirthDay      = villagersBirthDay{28}
	kodyBirthMonth    = villagersBirthMonth{time.September} // 9
	kodyBubbleColor   = villagersBubbleColor{villagerBubbleColor55CFFF}
	kodyCategory      = villagersCategory{villagerCategoryB}
	kodyClothes       = villagersClothes{villagerClothesWesternShirt} // 8477
	kodyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorWhite}}
	kodyFlooring      = villagersFlooring{villagerFlooringBlueDotFlooring}
	kodyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackRaceCarBed, villagerFurnitureThrowbackRocket, villagerFurnitureWoodenTable, villagerFurnitureWoodenBlockstool, villagerFurnitureMiniDIYWorkbench, villagerFurnitureThrowbackContainer, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenStool, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureWoodenChest, villagerFurnitureThrowbackWallClock, villagerFurnitureThrowbackWrestlingFigure, villagerFurnitureCuteMusicPlayer}}
	kodyGames         = villagersGames{[]VillagerGame{}} // TBD
	kodyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	kodyInterest      = villagersInterest{villagerInterestFitness}
	kodyName          = villagersName{villagerNameKody}
	kodyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	kodyMusic         = villagersMusic{villagerMusicMyPlace}
	kodyPersonality   = villagersPersonality{villagerPersonalityJock}
	kodySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	kodyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	kodyWallpaper     = villagersWallpaper{villagerWallpaperBluePlayroomWall}
)
