package atsumori

import "time"

// Bonbon is an Animal Crossing villager.
var Bonbon = villager{
	bonbonAstrology,
	bonbonBirthDay,
	bonbonBirthMonth,
	bonbonBubbleColor,
	bonbonCategory,
	bonbonClothes,
	bonbonClothesColors,
	bonbonFlooring,
	bonbonFurniture,
	bonbonGames,
	bonbonGender,
	bonbonInterest,
	bonbonName,
	bonbonNameColor,
	bonbonMusic,
	bonbonPersonality,
	bonbonSpecies,
	bonbonStyle,
	bonbonWallpaper}

var (
	bonbonAstrology     = villagersAstrology{villagerAstrologyPisces}
	bonbonBirthDay      = villagersBirthDay{3}
	bonbonBirthMonth    = villagersBirthMonth{time.March} // 3
	bonbonBubbleColor   = villagersBubbleColor{villagerBubbleColor7C6559}
	bonbonCategory      = villagersCategory{villagerCategoryA}
	bonbonClothes       = villagersClothes{} // 8360
	bonbonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorYellow}}
	bonbonFlooring      = villagersFlooring{villagerFlooringWhitePaintFlooring}
	bonbonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureCuteWardrobe, villagerFurnitureCuteChair, villagerFurnitureCuteVanity, villagerFurnitureCuteFloorLamp, villagerFurnitureCuteTeaTable, villagerFurnitureWhiteHeartRug, villagerFurnitureCuteMusicPlayer, villagerFurnitureCuteSofa, villagerFurnitureCuteBed, villagerFurnitureCuteWallMountedClock}}
	bonbonGames         = villagersGames{[]VillagerGame{}} // TBD
	bonbonGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	bonbonInterest      = villagersInterest{villagerInterestPlay}
	bonbonName          = villagersName{villagerNameBonbon}
	bonbonNameColor     = villagersNameColor{villagerNameColorEAC113}
	bonbonMusic         = villagersMusic{villagerMusicBubblegumKK}
	bonbonPersonality   = villagersPersonality{villagerPersonalityPeppy}
	bonbonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	bonbonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	bonbonWallpaper     = villagersWallpaper{villagerWallpaperCuteYellowWall}
)
