package atsumori

import (
	"time"
)

// Astrid is an Animal Crossing villager.
var Astrid = villager{
	astridAstrology,
	astridBirthDay,
	astridBirthMonth,
	astridBubbleColor,
	astridCategory,
	astridClothes,
	astridClothesColors,
	astridFlooring,
	astridFurniture,
	astridGames,
	astridGender,
	astridInterest,
	astridName,
	astridNameColor,
	astridMusic,
	astridPersonality,
	astridSpecies,
	astridStyle,
	astridWallpaper}

var (
	astridAstrology     = villagersAstrology{villagerAstrologyVirgo}
	astridBirthDay      = villagersBirthDay{8}
	astridBirthMonth    = villagersBirthMonth{time.September}
	astridBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	astridCategory      = villagersCategory{villagerCategoryB}
	astridClothes       = villagersClothes{} // 4559
	astridClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorColorful}}
	astridFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	astridFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSwingingBench, villagerFurnitureSpringyRideOn, villagerFurnitureTireToy, villagerFurnitureTireToy, villagerFurnitureTireToy, villagerFurnitureCassettePlayer, villagerFurnitureInflatableSofa, villagerFurnitureDrinkingFountain, villagerFurniturePlaygroundGym, villagerFurnitureTricycle}}
	astridGames         = villagersGames{[]VillagerGame{}} // TBD
	astridGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	astridInterest      = villagersInterest{villagerInterestMusic}
	astridName          = villagersName{villagerNameAstrid}
	astridNameColor     = villagersNameColor{villagerNameColor848484}
	astridMusic         = villagersMusic{villagerMusicKKGumbo}
	astridPersonality   = villagersPersonality{villagerPersonalitySnooty}
	astridSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKangaroo}}
	astridStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	astridWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
