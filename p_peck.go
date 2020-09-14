package atsumori

import "time"

// Peck is an Animal Crossing villager.
var Peck = villager{
	peckAstrology,
	peckBirthDay,
	peckBirthMonth,
	peckBubbleColor,
	peckCategory,
	peckClothes,
	peckClothesColors,
	peckFlooring,
	peckFurniture,
	peckGames,
	peckGender,
	peckInterest,
	peckName,
	peckNameColor,
	peckMusic,
	peckPersonality,
	peckSpecies,
	peckStyle,
	peckWallpaper}

var (
	peckAstrology     = villagersAstrology{villagerAstrologyLeo}
	peckBirthDay      = villagersBirthDay{25}
	peckBirthMonth    = villagersBirthMonth{time.July} // 7
	peckBubbleColor   = villagersBubbleColor{villagerBubbleColor4C3317}
	peckCategory      = villagersCategory{villagerCategoryA}
	peckClothes       = villagersClothes{villagerClothesRaglanShirt} // 2674
	peckClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorRed}}
	peckFlooring      = villagersFlooring{villagerFlooringBluePaintFlooring}
	peckFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDIYWorkbench, villagerFurnitureWoodenBlockStereo, villagerFurnitureMonstera, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureToyBox, villagerFurnitureThrowbackDinoScreen, villagerFurnitureThrowbackWallClock, villagerFurnitureWoodenBlockstool, villagerFurnitureThrowbackWrestlingFigure, villagerFurnitureWoodenEndTable, villagerFurnitureDigitalAlarmClock, villagerFurnitureWallMountedTV20In}}
	peckGames         = villagersGames{[]VillagerGame{}} // TBD
	peckGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	peckInterest      = villagersInterest{villagerInterestPlay}
	peckName          = villagersName{villagerNamePeck}
	peckNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	peckMusic         = villagersMusic{villagerMusicMyPlace}
	peckPersonality   = villagersPersonality{villagerPersonalityJock}
	peckSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	peckStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	peckWallpaper     = villagersWallpaper{villagerWallpaperCamoWall}
)
