package atsumori

import "time"

// Gloria is an Animal Crossing villager
var Gloria = villager{
	gloriaAstrology,
	gloriaBirthDay,
	gloriaBirthMonth,
	gloriaBubbleColor,
	gloriaCategory,
	gloriaClothes,
	gloriaClothesColors,
	gloriaFlooring,
	gloriaFurniture,
	gloriaGames,
	gloriaGender,
	gloriaInterest,
	gloriaName,
	gloriaNameColor,
	gloriaMusic,
	gloriaPersonality,
	gloriaSpecies,
	gloriaStyle,
	gloriaWallpaper}

var (
	gloriaAstrology     = villagersAstrology{villagerAstrologyLeo}
	gloriaBirthDay      = villagersBirthDay{12}
	gloriaBirthMonth    = villagersBirthMonth{time.August} // 8
	gloriaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	gloriaCategory      = villagersCategory{villagerCategoryB}
	gloriaClothes       = villagersClothes{} // 8896
	gloriaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	gloriaFlooring      = villagersFlooring{villagerFlooringSidewalkFlooring}
	gloriaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGardenBench, villagerFurnitureGardenBench, villagerFurnitureFountain, villagerFurnitureStreetlamp, villagerFurnitureStreetlamp, villagerFurniturePopcornMachine, villagerFurnitureManholeCover, villagerFurnitureStreetOrgan}}
	gloriaGames         = villagersGames{[]VillagerGame{}} // TBD
	gloriaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	gloriaInterest      = villagersInterest{villagerInterestFashion}
	gloriaName          = villagersName{villagerNameGloria}
	gloriaNameColor     = villagersNameColor{villagerNameColor848484}
	gloriaMusic         = villagersMusic{villagerMusicAnimalCity}
	gloriaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	gloriaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	gloriaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	gloriaWallpaper     = villagersWallpaper{villagerWallpaperCafeCurtainWall}
)
