package atsumori

import "time"

// Chadder is an Animal Crossing villager
var Chadder = villager{
	chadderAstrology,
	chadderBirthDay,
	chadderBirthMonth,
	chadderBubbleColor,
	chadderCategory,
	chadderClothes,
	chadderClothesColors,
	chadderFlooring,
	chadderFurniture,
	chadderGames,
	chadderGender,
	chadderInterest,
	chadderName,
	chadderNameColor,
	chadderMusic,
	chadderPersonality,
	chadderSpecies,
	chadderStyle,
	chadderWallpaper}

var (
	chadderAstrology     = villagersAstrology{villagerAstrologySagittarius}
	chadderBirthDay      = villagersBirthDay{15}
	chadderBirthMonth    = villagersBirthMonth{time.December} // 12
	chadderBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	chadderCategory      = villagersCategory{villagerCategoryA}
	chadderClothes       = villagersClothes{villagerClothesTailcoat} // 4440
	chadderClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	chadderFlooring      = villagersFlooring{villagerFlooringDarkWoodPatternFlooring}
	chadderFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShowerBooth, villagerFurnitureWhirlpoolBath, villagerFurnitureBeachTowel, villagerFurnitureBidet, villagerFurnitureTanklessToilet, villagerFurniturePlainSink, villagerFurniturePhonograph, villagerFurnitureImperialPartition, villagerFurnitureMonstera, villagerFurnitureAirConditioner, villagerFurnitureBathroomTowelRack, villagerFurnitureWallMountedTV50In}}
	chadderGames         = villagersGames{[]VillagerGame{}} // TBD
	chadderGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	chadderInterest      = villagersInterest{villagerInterestFitness}
	chadderName          = villagersName{villagerNameChadder}
	chadderNameColor     = villagersNameColor{villagerNameColor9B553A}
	chadderMusic         = villagersMusic{} // K.K. Soul
	chadderPersonality   = villagersPersonality{villagerPersonalitySmug}
	chadderSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	chadderStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	chadderWallpaper     = villagersWallpaper{villagerWallpaperCityscapeWall}
)
