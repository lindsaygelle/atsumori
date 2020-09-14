package atsumori

import "time"

// Peanut is an Animal Crossing villager.
var Peanut = villager{
	peanutAstrology,
	peanutBirthDay,
	peanutBirthMonth,
	peanutBubbleColor,
	peanutCategory,
	peanutClothes,
	peanutClothesColors,
	peanutFlooring,
	peanutFurniture,
	peanutGames,
	peanutGender,
	peanutInterest,
	peanutName,
	peanutNameColor,
	peanutMusic,
	peanutPersonality,
	peanutSpecies,
	peanutStyle,
	peanutWallpaper}

var (
	peanutAstrology     = villagersAstrology{villagerAstrologyGemini}
	peanutBirthDay      = villagersBirthDay{8}
	peanutBirthMonth    = villagersBirthMonth{time.June} // 6
	peanutBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	peanutCategory      = villagersCategory{villagerCategoryB}
	peanutClothes       = villagersClothes{} // 7784
	peanutClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorColorful}}
	peanutFlooring      = villagersFlooring{villagerFlooringCuteWhiteTileFlooring}
	peanutFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteBed, villagerFurnitureCuteVanity, villagerFurnitureCuteWardrobe, villagerFurnitureLCDTV50In, villagerFurnitureCuteFloorLamp, villagerFurnitureCushion, villagerFurnitureMug, villagerFurnitureCuteTeaTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureCuteDIYTable, villagerFurnitureCorkboard, villagerFurniturePinkHeartRug, villagerFurnitureCuteWallMountedClock}}
	peanutGames         = villagersGames{[]VillagerGame{}} // TBD
	peanutGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	peanutInterest      = villagersInterest{villagerInterestFashion}
	peanutName          = villagersName{villagerNamePeanut}
	peanutNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	peanutMusic         = villagersMusic{villagerMusicForestLife}
	peanutPersonality   = villagersPersonality{villagerPersonalityPeppy}
	peanutSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	peanutStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	peanutWallpaper     = villagersWallpaper{villagerWallpaperCuteRedWall}
)
