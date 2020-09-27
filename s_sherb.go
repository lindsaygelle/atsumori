package atsumori

import "time"

// Sherb is an Animal Crossing villager.
var Sherb = villager{
	sherbAstrology,
	sherbBirthDay,
	sherbBirthMonth,
	sherbBubbleColor,
	sherbCategory,
	sherbClothes,
	sherbClothesColors,
	sherbFlooring,
	sherbFurniture,
	sherbGames,
	sherbGender,
	sherbInterest,
	sherbName,
	sherbNameColor,
	sherbMusic,
	sherbPersonality,
	sherbSpecies,
	sherbStyle,
	sherbWallpaper}

var (
	sherbAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	sherbBirthDay      = villagersBirthDay{18}
	sherbBirthMonth    = villagersBirthMonth{time.January} // 1
	sherbBubbleColor   = villagersBubbleColor{villagerBubbleColorBFF2FF}
	sherbCategory      = villagersCategory{villagerCategoryA}
	sherbClothes       = villagersClothes{villagerClothesSnowySweater} // 3631
	sherbClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorBlue}}
	sherbFlooring      = villagersFlooring{villagerFlooringSimplePurpleFlooring}
	sherbFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureCuteFloorLamp, villagerFurnitureWoodenMiniTable, villagerFurnitureCuteSofa, villagerFurnitureCuteBed, villagerFurnitureCuteWardrobe, villagerFurnitureCuteTeaTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureOldFashionedAlarmClock}}
	sherbGames         = villagersGames{[]VillagerGame{}} // TBD
	sherbGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	sherbInterest      = villagersInterest{villagerInterestNature}
	sherbName          = villagersName{villagerNameSherb}
	sherbNameColor     = villagersNameColor{villagerNameColor85A2AF}
	sherbMusic         = villagersMusic{villagerMusicHypnoKK}
	sherbPersonality   = villagersPersonality{villagerPersonalityLazy}
	sherbSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGoat}}
	sherbStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	sherbWallpaper     = villagersWallpaper{villagerWallpaperPurpleDesertTileWall}
)
