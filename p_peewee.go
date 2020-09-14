package atsumori

import "time"

// Peewee is an Animal Crossing villager.
var Peewee = villager{
	peeweeAstrology,
	peeweeBirthDay,
	peeweeBirthMonth,
	peeweeBubbleColor,
	peeweeCategory,
	peeweeClothes,
	peeweeClothesColors,
	peeweeFlooring,
	peeweeFurniture,
	peeweeGames,
	peeweeGender,
	peeweeInterest,
	peeweeName,
	peeweeNameColor,
	peeweeMusic,
	peeweePersonality,
	peeweeSpecies,
	peeweeStyle,
	peeweeWallpaper}

var (
	peeweeAstrology     = villagersAstrology{villagerAstrologyVirgo}
	peeweeBirthDay      = villagersBirthDay{11}
	peeweeBirthMonth    = villagersBirthMonth{time.September} // 9
	peeweeBubbleColor   = villagersBubbleColor{villagerBubbleColor7FA9FF}
	peeweeCategory      = villagersCategory{villagerCategoryB}
	peeweeClothes       = villagersClothes{villagerClothesFiveBallTee} // 3322
	peeweeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorBlue}}
	peeweeFlooring      = villagersFlooring{villagerFlooringConcreteFlooring}
	peeweeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurnitureDrumSet, villagerFurnitureDinerSofa, villagerFurnitureDinerCounterChair, villagerFurnitureGasRange, villagerFurnitureWallMountedToolBoard, villagerFurnitureBarbell, villagerFurnitureDIYWorkbench, villagerFurnitureMiniFridge, villagerFurnitureDinerMiniTable, villagerFurnitureToolCart, villagerFurniturePotRack, villagerFurnitureRoughRug, villagerFurnitureProteinShakerBottle, villagerFurnitureTapeDeck}}
	peeweeGames         = villagersGames{[]VillagerGame{}} // TBD
	peeweeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	peeweeInterest      = villagersInterest{villagerInterestFitness}
	peeweeName          = villagersName{villagerNamePeewee}
	peeweeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	peeweeMusic         = villagersMusic{villagerMusicKKRock}
	peeweePersonality   = villagersPersonality{villagerPersonalityCranky}
	peeweeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGorilla}}
	peeweeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	peeweeWallpaper     = villagersWallpaper{villagerWallpaperSteelFrameWall}
)
