package atsumori

import "time"

// Keaton is an Animal Crossing villager.
var Keaton = villager{
	keatonAstrology,
	keatonBirthDay,
	keatonBirthMonth,
	keatonBubbleColor,
	keatonCategory,
	keatonClothes,
	keatonClothesColors,
	keatonFlooring,
	keatonFurniture,
	keatonGames,
	keatonGender,
	keatonInterest,
	keatonName,
	keatonNameColor,
	keatonMusic,
	keatonPersonality,
	keatonSpecies,
	keatonStyle,
	keatonWallpaper}

var (
	keatonAstrology     = villagersAstrology{villagerAstrologyGemini}
	keatonBirthDay      = villagersBirthDay{1}
	keatonBirthMonth    = villagersBirthMonth{time.June} // 6
	keatonBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	keatonCategory      = villagersCategory{villagerCategoryA}
	keatonClothes       = villagersClothes{villagerClothesFuzzyVest} // 4212
	keatonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	keatonFlooring      = villagersFlooring{villagerFlooringParkingFlooring}
	keatonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackRaceCarBed, villagerFurnitureThrowbackRaceCarBed, villagerFurniturePlasticCanister, villagerFurnitureCone, villagerFurnitureCone, villagerFurnitureMetalCan, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureRetroGasPump, villagerFurnitureManholeCover, villagerFurnitureBambooStopblock, villagerFurnitureCardboardBox, villagerFurnitureCassettePlayer}}
	keatonGames         = villagersGames{[]VillagerGame{}} // TBD
	keatonGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	keatonInterest      = villagersInterest{villagerInterestMusic}
	keatonName          = villagersName{villagerNameKeaton}
	keatonNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	keatonMusic         = villagersMusic{villagerMusicDrivin}
	keatonPersonality   = villagersPersonality{villagerPersonalitySmug}
	keatonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	keatonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCool}}
	keatonWallpaper     = villagersWallpaper{villagerWallpaperTreeLinedWall}
)
