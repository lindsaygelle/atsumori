package atsumori

import "time"

// Joey is an Animal Crossing villager
var Joey = villager{
	joeyAstrology,
	joeyBirthDay,
	joeyBirthMonth,
	joeyBubbleColor,
	joeyCategory,
	joeyClothes,
	joeyClothesColors,
	joeyFlooring,
	joeyFurniture,
	joeyGames,
	joeyGender,
	joeyInterest,
	joeyName,
	joeyNameColor,
	joeyMusic,
	joeyPersonality,
	joeySpecies,
	joeyStyle,
	joeyWallpaper}

var (
	joeyAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	joeyBirthDay      = villagersBirthDay{3}
	joeyBirthMonth    = villagersBirthMonth{time.January} // 1
	joeyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	joeyCategory      = villagersCategory{villagerCategoryB}
	joeyClothes       = villagersClothes{villagerClothesBearTee} // 8196
	joeyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBlue}}
	joeyFlooring      = villagersFlooring{villagerFlooringSandyBeachFlooring}
	joeyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLifeRing, villagerFurnitureRattanTowelBasket, villagerFurniturePlasticPool, villagerFurnitureBeachBall, villagerFurnitureBeachChair, villagerFurnitureSandCastle, villagerFurnitureCassettePlayer, villagerFurnitureCoconutJuice, villagerFurnitureTropicalRug, villagerFurnitureLifeguardChair}}
	joeyGames         = villagersGames{[]VillagerGame{}} // TBD
	joeyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	joeyInterest      = villagersInterest{villagerInterestPlay}
	joeyName          = villagersName{villagerNameJoey}
	joeyNameColor     = villagersNameColor{villagerNameColor9B553A}
	joeyMusic         = villagersMusic{villagerMusicMarineSong2001}
	joeyPersonality   = villagersPersonality{villagerPersonalityLazy}
	joeySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	joeyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	joeyWallpaper     = villagersWallpaper{villagerWallpaperTropicalVista}
)
