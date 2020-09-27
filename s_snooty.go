package atsumori

import "time"

// Snooty is an Animal Crossing villager.
var Snooty = villager{
	snootyAstrology,
	snootyBirthDay,
	snootyBirthMonth,
	snootyBubbleColor,
	snootyCategory,
	snootyClothes,
	snootyClothesColors,
	snootyFlooring,
	snootyFurniture,
	snootyGames,
	snootyGender,
	snootyInterest,
	snootyName,
	snootyNameColor,
	snootyMusic,
	snootyPersonality,
	snootySpecies,
	snootyStyle,
	snootyWallpaper}

var (
	snootyAstrology     = villagersAstrology{villagerAstrologyScorpio}
	snootyBirthDay      = villagersBirthDay{24}
	snootyBirthMonth    = villagersBirthMonth{time.October} // 10
	snootyBubbleColor   = villagersBubbleColor{villagerBubbleColorF2BDC7}
	snootyCategory      = villagersCategory{villagerCategoryA}
	snootyClothes       = villagersClothes{villagerClothesMistyTee} // 8201
	snootyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorYellow}}
	snootyFlooring      = villagersFlooring{villagerFlooringTatamiFlooring}
	snootyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureKimonoStand, villagerFurnitureZenCushion, villagerFurnitureTatamiBed, villagerFurnitureBambooPartition, villagerFurniturePaperLantern, villagerFurniturePileOfZenCushions, villagerFurnitureClayFurnace, villagerFurnitureBambooWallDecoration, villagerFurnitureTeaTable, villagerFurnitureBambooLunchBox}}
	snootyGames         = villagersGames{[]VillagerGame{}} // TBD
	snootyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	snootyInterest      = villagersInterest{villagerInterestEducation}
	snootyName          = villagersName{villagerNameSnooty}
	snootyNameColor     = villagersNameColor{villagerNameColor634B4B}
	snootyMusic         = villagersMusic{villagerMusicKKFolk}
	snootyPersonality   = villagersPersonality{villagerPersonalitySnooty}
	snootySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAnteater}}
	snootyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	snootyWallpaper     = villagersWallpaper{villagerWallpaperShojiScreen}
)
