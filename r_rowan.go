package atsumori

import "time"

// Rowan is an Animal Crossing villager.
var Rowan = villager{
	rowanAstrology,
	rowanBirthDay,
	rowanBirthMonth,
	rowanBubbleColor,
	rowanCategory,
	rowanClothes,
	rowanClothesColors,
	rowanFlooring,
	rowanFurniture,
	rowanGames,
	rowanGender,
	rowanInterest,
	rowanName,
	rowanNameColor,
	rowanMusic,
	rowanPersonality,
	rowanSpecies,
	rowanStyle,
	rowanWallpaper}

var (
	rowanAstrology     = villagersAstrology{villagerAstrologyVirgo}
	rowanBirthDay      = villagersBirthDay{26}
	rowanBirthMonth    = villagersBirthMonth{time.August} // 8
	rowanBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	rowanCategory      = villagersCategory{villagerCategoryB}
	rowanClothes       = villagersClothes{villagerClothesSimpleDotsTee} // 9498
	rowanClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorGray}}
	rowanFlooring      = villagersFlooring{villagerFlooringTigerPrintFlooring}
	rowanFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBambooBench, villagerFurnitureImperialLowTable, villagerFurnitureScreen, villagerFurnitureImperialBed, villagerFurnitureBambooWallDecoration, villagerFurnitureImperialChest, villagerFurniturePortableRecordPlayer, villagerFurniturePaperTiger, villagerFurnitureImperialDecorativeShelves}}
	rowanGames         = villagersGames{[]VillagerGame{}} // TBD
	rowanGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rowanInterest      = villagersInterest{villagerInterestFitness}
	rowanName          = villagersName{villagerNameRowan}
	rowanNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	rowanMusic         = villagersMusic{villagerMusicSurfinKK}
	rowanPersonality   = villagersPersonality{villagerPersonalityJock}
	rowanSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesTiger}}
	rowanStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	rowanWallpaper     = villagersWallpaper{villagerWallpaperBambooWall}
)
