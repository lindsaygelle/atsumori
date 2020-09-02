package atsumori

import "time"

// Melba is an Animal Crossing villager.
var Melba = villager{
	melbaAstrology,
	melbaBirthDay,
	melbaBirthMonth,
	melbaBubbleColor,
	melbaCategory,
	melbaClothes,
	melbaClothesColors,
	melbaFlooring,
	melbaFurniture,
	melbaGames,
	melbaGender,
	melbaInterest,
	melbaName,
	melbaNameColor,
	melbaMusic,
	melbaPersonality,
	melbaSpecies,
	melbaStyle,
	melbaWallpaper}

var (
	melbaAstrology     = villagersAstrology{villagerAstrologyAries}
	melbaBirthDay      = villagersBirthDay{12}
	melbaBirthMonth    = villagersBirthMonth{time.April} // 4
	melbaBubbleColor   = villagersBubbleColor{villagerBubbleColorE4B887}
	melbaCategory      = villagersCategory{villagerCategoryB}
	melbaClothes       = villagersClothes{} // 4399
	melbaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorGreen}}
	melbaFlooring      = villagersFlooring{villagerFlooringGreenPaintFlooring}
	melbaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePlainSink, villagerFurnitureCuteMusicPlayer, villagerFurnitureYucca, villagerFurnitureStandardUmbrellaStand, villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureWoodenWardrobe, villagerFurnitureWallMountedTV50In, villagerFurnitureWoodenEndTable, villagerFurnitureAutomaticWasher, villagerFurnitureWoodenLowTable, villagerFurnitureCushion, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureBotanicalRug, villagerFurniturePottedIvy, villagerFurnitureUnfinishedPuzzle}}
	melbaGames         = villagersGames{[]VillagerGame{}} // TBD
	melbaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	melbaInterest      = villagersInterest{villagerInterestEducation}
	melbaName          = villagersName{villagerNameMelba}
	melbaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	melbaMusic         = villagersMusic{villagerMusicKKStroll}
	melbaPersonality   = villagersPersonality{villagerPersonalityNormal}
	melbaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	melbaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	melbaWallpaper     = villagersWallpaper{villagerWallpaperBlueBlossomingWall}
)
