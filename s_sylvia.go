package atsumori

import "time"

// Sylvia is an Animal Crossing villager.
var Sylvia = villager{
	sylviaAstrology,
	sylviaBirthDay,
	sylviaBirthMonth,
	sylviaBubbleColor,
	sylviaCategory,
	sylviaClothes,
	sylviaClothesColors,
	sylviaFlooring,
	sylviaFurniture,
	sylviaGames,
	sylviaGender,
	sylviaInterest,
	sylviaName,
	sylviaNameColor,
	sylviaMusic,
	sylviaPersonality,
	sylviaSpecies,
	sylviaStyle,
	sylviaWallpaper}

var (
	sylviaAstrology     = villagersAstrology{villagerAstrologyTaurus}
	sylviaBirthDay      = villagersBirthDay{3}
	sylviaBirthMonth    = villagersBirthMonth{time.May} // 5
	sylviaBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	sylviaCategory      = villagersCategory{villagerCategoryB}
	sylviaClothes       = villagersClothes{} // 9578
	sylviaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorGreen}}
	sylviaFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	sylviaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackRaceCarBed, villagerFurnitureDIYWorkbench, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureDinerSofa, villagerFurnitureKitchenIsland, villagerFurniturePortableRadio, villagerFurnitureBlackWoodenDeckRug, villagerFurnitureBlackWoodenDeckRug, villagerFurnitureToolCart, villagerFurnitureIronwoodCupboard, villagerFurnitureKeyHolder, villagerFurnitureToolbox, villagerFurnitureMicrowave, villagerFurnitureMixer, villagerFurnitureWallMountedToolBoard}}
	sylviaGames         = villagersGames{[]VillagerGame{}} // TBD
	sylviaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	sylviaInterest      = villagersInterest{villagerInterestMusic}
	sylviaName          = villagersName{villagerNameSylvia}
	sylviaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	sylviaMusic         = villagersMusic{villagerMusicDJKK}
	sylviaPersonality   = villagersPersonality{villagerPersonalityBigSister}
	sylviaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKangaroo}}
	sylviaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleGorgeous}}
	sylviaWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
