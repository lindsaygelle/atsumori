package atsumori

import "time"

// Marcel is an Animal Crossing villager.
var Marcel = villager{
	marcelAstrology,
	marcelBirthDay,
	marcelBirthMonth,
	marcelBubbleColor,
	marcelCategory,
	marcelClothes,
	marcelClothesColors,
	marcelFlooring,
	marcelFurniture,
	marcelGames,
	marcelGender,
	marcelInterest,
	marcelName,
	marcelNameColor,
	marcelMusic,
	marcelPersonality,
	marcelSpecies,
	marcelStyle,
	marcelWallpaper}

var (
	marcelAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	marcelBirthDay      = villagersBirthDay{31}
	marcelBirthMonth    = villagersBirthMonth{time.December} // 12
	marcelBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	marcelCategory      = villagersCategory{villagerCategoryB}
	marcelClothes       = villagersClothes{villagerClothesSeaHantenShirt} // 4407
	marcelClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBlue}}
	marcelFlooring      = villagersFlooring{villagerFlooringRushTatami}
	marcelFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureOldFashionedWashtub, villagerFurnitureFloorSeat, villagerFurnitureKotatsu, villagerFurnitureFuton, villagerFurnitureZenCushion, villagerFurnitureBambooSpeaker, villagerFurniturePaperLantern, villagerFurnitureBrownWoodenDeckRug, villagerFurniturePileOfZenCushions, villagerFurnitureBrownWoodenDeckRug, villagerFurnitureClayFurnace, villagerFurnitureBambooStool, villagerFurnitureGoldfish}}
	marcelGames         = villagersGames{[]VillagerGame{}} // TBD
	marcelGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	marcelInterest      = villagersInterest{villagerInterestPlay}
	marcelName          = villagersName{villagerNameMarcel}
	marcelNameColor     = villagersNameColor{villagerNameColor848484}
	marcelMusic         = villagersMusic{villagerMusicComradeKK}
	marcelPersonality   = villagersPersonality{villagerPersonalityLazy}
	marcelSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	marcelStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	marcelWallpaper     = villagersWallpaper{villagerWallpaperModernShojiScreenWall}
)
