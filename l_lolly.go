package atsumori

import "time"

// Lolly is an Animal Crossing villager.
var Lolly = villager{
	lollyAstrology,
	lollyBirthDay,
	lollyBirthMonth,
	lollyBubbleColor,
	lollyCategory,
	lollyClothes,
	lollyClothesColors,
	lollyFlooring,
	lollyFurniture,
	lollyGames,
	lollyGender,
	lollyInterest,
	lollyName,
	lollyNameColor,
	lollyMusic,
	lollyPersonality,
	lollySpecies,
	lollyStyle,
	lollyWallpaper}

var (
	lollyAstrology     = villagersAstrology{villagerAstrologyAries}
	lollyBirthDay      = villagersBirthDay{27}
	lollyBirthMonth    = villagersBirthMonth{time.March} // 3
	lollyBubbleColor   = villagersBubbleColor{villagerBubbleColorBFF2FF}
	lollyCategory      = villagersCategory{villagerCategoryB}
	lollyClothes       = villagersClothes{villagerClothesSnowySweater} // 3631
	lollyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorPink}}
	lollyFlooring      = villagersFlooring{villagerFlooringGreenVinylFlooring}
	lollyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureWoodenBlockBed, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureWoodenStool, villagerFurnitureWoodenBlockChest, villagerFurniturePortableRecordPlayer, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureKitchenIsland, villagerFurnitureMiniDIYWorkbench, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureWoodenBlockWallClock, villagerFurniturePottedIvy, villagerFurniturePottedIvy}}
	lollyGames         = villagersGames{[]VillagerGame{}} // TBD
	lollyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	lollyInterest      = villagersInterest{villagerInterestMusic}
	lollyName          = villagersName{villagerNameLolly}
	lollyNameColor     = villagersNameColor{villagerNameColor85A2AF}
	lollyMusic         = villagersMusic{villagerMusicForestLife}
	lollyPersonality   = villagersPersonality{villagerPersonalityNormal}
	lollySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	lollyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	lollyWallpaper     = villagersWallpaper{villagerWallpaperModernWoodWall}
)
