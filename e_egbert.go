package atsumori

import "time"

// Egbert is an Animal Crossing villager
var Egbert = villager{
	egbertAstrology,
	egbertBirthDay,
	egbertBirthMonth,
	egbertBubbleColor,
	egbertCategory,
	egbertClothes,
	egbertClothesColors,
	egbertFlooring,
	egbertFurniture,
	egbertGames,
	egbertGender,
	egbertInterest,
	egbertName,
	egbertNameColor,
	egbertMusic,
	egbertPersonality,
	egbertSpecies,
	egbertStyle,
	egbertWallpaper}

var (
	egbertAstrology     = villagersAstrology{villagerAstrologyLibra}
	egbertBirthDay      = villagersBirthDay{14}
	egbertBirthMonth    = villagersBirthMonth{time.October} // 10
	egbertBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	egbertCategory      = villagersCategory{villagerCategoryB}
	egbertClothes       = villagersClothes{villagerClothesFolkShirt} // 4446
	egbertClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorBrown}}
	egbertFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	egbertFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureBeachTowel, villagerFurnitureWaterCooler, villagerFurnitureShowerBooth, villagerFurnitureSaunaHeater, villagerFurnitureSaunaHeater, villagerFurnitureLogExtraLongSofa, villagerFurnitureLogExtraLongSofa, villagerFurnitureLogWallMountedClock, villagerFurnitureIvorySimpleBathMat, villagerFurnitureIvorySimpleBathMat, villagerFurnitureLogBench, villagerFurnitureIvorySimpleBathMat, villagerFurnitureHourglass, villagerFurnitureRattanTowelBasket}}
	egbertGames         = villagersGames{[]VillagerGame{}} // TBD
	egbertGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	egbertInterest      = villagersInterest{villagerInterestPlay}
	egbertName          = villagersName{villagerNameEgbert}
	egbertNameColor     = villagersNameColor{villagerNameColor9B553A}
	egbertMusic         = villagersMusic{villagerMusicLuckyKK}
	egbertPersonality   = villagersPersonality{villagerPersonalityLazy}
	egbertSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	egbertStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	egbertWallpaper     = villagersWallpaper{villagerWallpaperCabinWall}
)
