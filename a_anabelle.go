package atsumori

import (
	"time"
)

// Anabelle is an Animal Crossing villager
var Anabelle = villager{
	anabelleAstrology,
	anabelleBirthDay,
	anabelleBirthMonth,
	anabelleBubbleColor,
	anabelleCategory,
	anabelleClothes,
	anabelleClothesColors,
	anabelleFlooring,
	anabelleFurniture,
	anabelleGames,
	anabelleGender,
	anabelleInterest,
	anabelleName,
	anabelleNameColor,
	anabelleMusic,
	anabellePersonality,
	anabelleSpecies,
	anabelleStyle,
	anabelleWallpaper}

var (
	anabelleAstrology     = villagersAstrology{villagerAstrologyAquarius}
	anabelleBirthDay      = villagersBirthDay{16}
	anabelleBirthMonth    = villagersBirthMonth{time.February}
	anabelleBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	anabelleCategory      = villagersCategory{villagerCategoryB}
	anabelleClothes       = villagersClothes{} // 8190
	anabelleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBlue}}
	anabelleFlooring      = villagersFlooring{villagerFlooringYellowFloralFlooring}
	anabelleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenChair, villagerFurnitureDeluxeWasher, villagerFurnitureCuteDIYTable, villagerFurnitureRefrigerator, villagerFurnitureWoodenWasteBin, villagerFurnitureWoodenTable, villagerFurnitureVentilationFan, villagerFurniturePortableRadio, villagerFurnitureWoodenEndTable, villagerFurnitureIroningSet, villagerFurnitureWoodenSimpleBed, villagerFurnitureIronwoodKitchenette, villagerFurnitureStandMixer, villagerFurnitureWallFan}}
	anabelleGames         = villagersGames{[]VillagerGame{}} // TBD
	anabelleGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	anabelleInterest      = villagersInterest{villagerInterestFashion}
	anabelleName          = villagersName{villagerNameAnabelle}
	anabelleNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	anabelleMusic         = villagersMusic{} // Aloha K.K.
	anabellePersonality   = villagersPersonality{villagerPersonalityPeppy}
	anabelleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAnteater}}
	anabelleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	anabelleWallpaper     = villagersWallpaper{villagerWallpaperGreenPaintedWoodWall}
)
