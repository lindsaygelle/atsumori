package atsumori

import "time"

// Midge is an Animal Crossing villager.
var Midge = villager{
	midgeAstrology,
	midgeBirthDay,
	midgeBirthMonth,
	midgeBubbleColor,
	midgeCategory,
	midgeClothes,
	midgeClothesColors,
	midgeFlooring,
	midgeFurniture,
	midgeGames,
	midgeGender,
	midgeInterest,
	midgeName,
	midgeNameColor,
	midgeMusic,
	midgePersonality,
	midgeSpecies,
	midgeStyle,
	midgeWallpaper}

var (
	midgeAstrology     = villagersAstrology{villagerAstrologyPisces}
	midgeBirthDay      = villagersBirthDay{12}
	midgeBirthMonth    = villagersBirthMonth{time.March} // 3
	midgeBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	midgeCategory      = villagersCategory{villagerCategoryB}
	midgeClothes       = villagersClothes{villagerClothesSilkShirt} // 9128
	midgeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorPink}}
	midgeFlooring      = villagersFlooring{villagerFlooringCuteWhiteTileFlooring}
	midgeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureWoodenLowTable, villagerFurnitureCuteChair, villagerFurnitureClawFootTub, villagerFurnitureCuteWallMountedClock, villagerFurnitureCuteBed, villagerFurnitureCuteWardrobe, villagerFurnitureCuteVanity, villagerFurnitureCuteTeaTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureCuteSofa}}
	midgeGames         = villagersGames{[]VillagerGame{}} // TBD
	midgeGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	midgeInterest      = villagersInterest{villagerInterestEducation}
	midgeName          = villagersName{villagerNameMidge}
	midgeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	midgeMusic         = villagersMusic{villagerMusicKKStroll}
	midgePersonality   = villagersPersonality{villagerPersonalityNormal}
	midgeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	midgeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	midgeWallpaper     = villagersWallpaper{villagerWallpaperYellowQuiltWall}
)
