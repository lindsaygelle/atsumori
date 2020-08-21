package atsumori

import "time"

// Anicotti is an Animal Crossing villager
var Anicotti = villager{
	anicottiAstrology,
	anicottiBirthDay,
	anicottiBirthMonth,
	anicottiBubbleColor,
	anicottiCategory,
	anicottiClothes,
	anicottiClothesColors,
	anicottiFlooring,
	anicottiFurniture,
	anicottiGames,
	anicottiGender,
	anicottiInterest,
	anicottiName,
	anicottiNameColor,
	anicottiMusic,
	anicottiPersonality,
	anicottiSpecies,
	anicottiStyle,
	anicottiWallpaper}

var (
	anicottiAstrology     = villagersAstrology{villagerAstrologyPisces}
	anicottiBirthDay      = villagersBirthDay{24}
	anicottiBirthMonth    = villagersBirthMonth{time.February}
	anicottiBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	anicottiCategory      = villagersCategory{villagerCategoryB}
	anicottiClothes       = villagersClothes{villagerClothesColorfulStripedSweater} // 7641
	anicottiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorPink}}
	anicottiFlooring      = villagersFlooring{villagerFlooringKitschyTile}
	anicottiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureYucca, villagerFurnitureWoodenChest, villagerFurnitureWoodenSimpleBed, villagerFurnitureMiniFridge, villagerFurniturePotRack, villagerFurnitureDishDryingRack, villagerFurnitureGasRange, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenTable, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureLilyRecordPlayer, villagerFurnitureWoodenChair, villagerFurnitureSimpleKettle, villagerFurnitureHumidifier, villagerFurnitureWhiteSimpleMediumMat}}
	anicottiGames         = villagersGames{[]VillagerGame{}} // TBD
	anicottiGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	anicottiInterest      = villagersInterest{villagerInterestFashion}
	anicottiName          = villagersName{villagerNameAnicotti}
	anicottiNameColor     = villagersNameColor{villagerNameColor28665A}
	anicottiMusic         = villagersMusic{villagerMusicNeapolitan}
	anicottiPersonality   = villagersPersonality{villagerPersonalityPeppy}
	anicottiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	anicottiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	anicottiWallpaper     = villagersWallpaper{villagerWallpaperBeadedCurtainWall}
)
