package atsumori

import "time"

// Freya is an Animal Crossing villager
var Freya = villager{
	freyaAstrology,
	freyaBirthDay,
	freyaBirthMonth,
	freyaBubbleColor,
	freyaCategory,
	freyaClothes,
	freyaClothesColors,
	freyaFlooring,
	freyaFurniture,
	freyaGames,
	freyaGender,
	freyaInterest,
	freyaName,
	freyaNameColor,
	freyaMusic,
	freyaPersonality,
	freyaSpecies,
	freyaStyle,
	freyaWallpaper}

var (
	freyaAstrology     = villagersAstrology{villagerAstrologySagittarius}
	freyaBirthDay      = villagersBirthDay{14}
	freyaBirthMonth    = villagersBirthMonth{time.December} // 12
	freyaBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	freyaCategory      = villagersCategory{villagerCategoryB}
	freyaClothes       = villagersClothes{villagerClothesReindeerSweater} // 4566
	freyaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBlue}}
	freyaFlooring      = villagersFlooring{villagerFlooringRoseFlooring}
	freyaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBed, villagerFurnitureAntiqueVanity, villagerFurnitureAntiqueChair, villagerFurnitureAntiqueChair, villagerFurnitureMusicStand, villagerFurnitureRetroRadiator, villagerFurnitureAntiqueMiniTable, villagerFurnitureFancyViolin, villagerFurnitureAntiqueTable, villagerFurnitureFireplace, villagerFurniturePhonograph, villagerFurnitureGlassHolderWithCandle, villagerFurnitureMatryoshka}}
	freyaGames         = villagersGames{[]VillagerGame{}} // TBD
	freyaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	freyaInterest      = villagersInterest{villagerInterestFashion}
	freyaName          = villagersName{villagerNameFreya}
	freyaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	freyaMusic         = villagersMusic{villagerMusicLuckyKK}
	freyaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	freyaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	freyaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	freyaWallpaper     = villagersWallpaper{villagerWallpaperHeavyCurtainWall}
)
