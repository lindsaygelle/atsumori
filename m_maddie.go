package atsumori

import "time"

// Maddie is an Animal Crossing villager.
var Maddie = villager{
	maddieAstrology,
	maddieBirthDay,
	maddieBirthMonth,
	maddieBubbleColor,
	maddieCategory,
	maddieClothes,
	maddieClothesColors,
	maddieFlooring,
	maddieFurniture,
	maddieGames,
	maddieGender,
	maddieInterest,
	maddieName,
	maddieNameColor,
	maddieMusic,
	maddiePersonality,
	maddieSpecies,
	maddieStyle,
	maddieWallpaper}

var (
	maddieAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	maddieBirthDay      = villagersBirthDay{11}
	maddieBirthMonth    = villagersBirthMonth{time.January} // 1
	maddieBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	maddieCategory      = villagersCategory{villagerCategoryA}
	maddieClothes       = villagersClothes{} // 8378
	maddieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorPink}}
	maddieFlooring      = villagersFlooring{villagerFlooringWildflowerMeadow}
	maddieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteMusicPlayer, villagerFurnitureOutdoorPicnicSet, villagerFurnitureBoardGame, villagerFurnitureFruitBasket, villagerFurnitureMountainBike, villagerFurnitureInflatableSofa, villagerFurnitureCoolerBox, villagerFurniturePicnicBasket, villagerFurnitureLogBench, villagerFurnitureSoupKettle, villagerFurnitureHandyWaterCooler, villagerFurniturePeachCheckedRug}}
	maddieGames         = villagersGames{[]VillagerGame{}} // TBD
	maddieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	maddieInterest      = villagersInterest{villagerInterestPlay}
	maddieName          = villagersName{villagerNameMaddie}
	maddieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	maddieMusic         = villagersMusic{villagerMusicKKStroll}
	maddiePersonality   = villagersPersonality{villagerPersonalityPeppy}
	maddieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	maddieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	maddieWallpaper     = villagersWallpaper{villagerWallpaperMeadowVista}
)
