package atsumori

import "time"

// Lily is an Animal Crossing villager.
var Lily = villager{
	lilyAstrology,
	lilyBirthDay,
	lilyBirthMonth,
	lilyBubbleColor,
	lilyCategory,
	lilyClothes,
	lilyClothesColors,
	lilyFlooring,
	lilyFurniture,
	lilyGames,
	lilyGender,
	lilyInterest,
	lilyName,
	lilyNameColor,
	lilyMusic,
	lilyPersonality,
	lilySpecies,
	lilyStyle,
	lilyWallpaper}

var (
	lilyAstrology     = villagersAstrology{villagerAstrologyAquarius}
	lilyBirthDay      = villagersBirthDay{4}
	lilyBirthMonth    = villagersBirthMonth{time.February} // 2
	lilyBubbleColor   = villagersBubbleColor{villagerBubbleColor00D1BD}
	lilyCategory      = villagersCategory{villagerCategoryB}
	lilyClothes       = villagersClothes{} // 9576
	lilyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorYellow}}
	lilyFlooring      = villagersFlooring{villagerFlooringBlueDotFlooring}
	lilyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGasRange, villagerFurnitureRattanBed, villagerFurnitureClawFootTub, villagerFurnitureShowerSet, villagerFurnitureStandardUmbrellaStand, villagerFurniturePlainSink, villagerFurnitureWoodenBlockstool, villagerFurnitureRattanTowelBasket, villagerFurnitureRattanEndTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureRattanLowTable, villagerFurnitureHyacinthLamp, villagerFurnitureMiniFridge, villagerFurnitureDishDryingRack, villagerFurnitureBathroomTowelRack, villagerFurnitureSimplePanel, villagerFurnitureSimplePanel}}
	lilyGames         = villagersGames{[]VillagerGame{}} // TBD
	lilyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	lilyInterest      = villagersInterest{villagerInterestEducation}
	lilyName          = villagersName{villagerNameLily}
	lilyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	lilyMusic         = villagersMusic{villagerMusicFarewell}
	lilyPersonality   = villagersPersonality{villagerPersonalityNormal}
	lilySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	lilyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	lilyWallpaper     = villagersWallpaper{villagerWallpaperMistyGardenWall}
)
