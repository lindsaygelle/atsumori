package atsumori

import "time"

// Sterling is an Animal Crossing villager.
var Sterling = villager{
	sterlingAstrology,
	sterlingBirthDay,
	sterlingBirthMonth,
	sterlingBubbleColor,
	sterlingCategory,
	sterlingClothes,
	sterlingClothesColors,
	sterlingFlooring,
	sterlingFurniture,
	sterlingGames,
	sterlingGender,
	sterlingInterest,
	sterlingName,
	sterlingNameColor,
	sterlingMusic,
	sterlingPersonality,
	sterlingSpecies,
	sterlingStyle,
	sterlingWallpaper}

var (
	sterlingAstrology     = villagersAstrology{villagerAstrologySagittarius}
	sterlingBirthDay      = villagersBirthDay{11}
	sterlingBirthMonth    = villagersBirthMonth{time.December} // 12
	sterlingBubbleColor   = villagersBubbleColor{villagerBubbleColorBFBFBF}
	sterlingCategory      = villagersCategory{villagerCategoryB}
	sterlingClothes       = villagersClothes{villagerClothesCavalierShirt} // 3233
	sterlingClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorRed}}
	sterlingFlooring      = villagersFlooring{villagerFlooringBasementFlooring}
	sterlingFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBarrel, villagerFurnitureDenChair, villagerFurnitureUtilitySink, villagerFurnitureBunkBed, villagerFurnitureJailBars, villagerFurnitureWallMountedCandle, villagerFurnitureKeyHolder, villagerFurnitureJailBars, villagerFurnitureDenDesk, villagerFurnitureCassettePlayer, villagerFurnitureDocumentStack, villagerFurnitureToolCart, villagerFurnitureWallMountedCandle, villagerFurnitureDinnerware}}
	sterlingGames         = villagersGames{[]VillagerGame{}} // TBD
	sterlingGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	sterlingInterest      = villagersInterest{villagerInterestFitness}
	sterlingName          = villagersName{villagerNameSterling}
	sterlingNameColor     = villagersNameColor{villagerNameColor5E5E5E}
	sterlingMusic         = villagersMusic{villagerMusicKKDirge}
	sterlingPersonality   = villagersPersonality{villagerPersonalityJock}
	sterlingSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	sterlingStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	sterlingWallpaper     = villagersWallpaper{villagerWallpaperDungeonWall}
)
