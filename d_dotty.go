package atsumori

import "time"

// Dotty is an Animal Crossing villager
var Dotty = villager{
	dottyAstrology,
	dottyBirthDay,
	dottyBirthMonth,
	dottyBubbleColor,
	dottyCategory,
	dottyClothes,
	dottyClothesColors,
	dottyFlooring,
	dottyFurniture,
	dottyGames,
	dottyGender,
	dottyInterest,
	dottyName,
	dottyNameColor,
	dottyMusic,
	dottyPersonality,
	dottySpecies,
	dottyStyle,
	dottyWallpaper}

var (
	dottyAstrology     = villagersAstrology{villagerAstrologyPisces}
	dottyBirthDay      = villagersBirthDay{14}
	dottyBirthMonth    = villagersBirthMonth{time.March} // 3
	dottyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	dottyCategory      = villagersCategory{villagerCategoryB}
	dottyClothes       = villagersClothes{villagerClothesSleevelessTunic} // 8359
	dottyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorBlack}}
	dottyFlooring      = villagersFlooring{villagerFlooringPineBoardFlooring}
	dottyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGasRange, villagerFurnitureFloralSwag, villagerFurnitureSimpleKettle, villagerFurnitureFloorLamp, villagerFurnitureYucca, villagerFurniturePotRack, villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenLowTable, villagerFurnitureWoodenWardrobe, villagerFurnitureMiniFridge, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureWoodenEndTable, villagerFurnitureIronwoodClock, villagerFurnitureGreenShaggyRug, villagerFurnitureTeaSet, villagerFurniturePopUpToaster, villagerFurniturePortableRecordPlayer}}
	dottyGames         = villagersGames{[]VillagerGame{}} // TBD
	dottyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	dottyInterest      = villagersInterest{villagerInterestFashion}
	dottyName          = villagersName{villagerNameDotty}
	dottyNameColor     = villagersNameColor{villagerNameColor848484}
	dottyMusic         = villagersMusic{villagerMusicDrivin}
	dottyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	dottySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	dottyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	dottyWallpaper     = villagersWallpaper{villagerWallpaperGreenIntricateWall}
)
