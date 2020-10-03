package atsumori

import "time"

// Velma is an Animal Crossing villager.
var Velma = villager{
	velmaAstrology,
	velmaBirthDay,
	velmaBirthMonth,
	velmaBubbleColor,
	velmaCategory,
	velmaClothes,
	velmaClothesColors,
	velmaFlooring,
	velmaFurniture,
	velmaGames,
	velmaGender,
	velmaInterest,
	velmaName,
	velmaNameColor,
	velmaMusic,
	velmaPersonality,
	velmaSpecies,
	velmaStyle,
	velmaWallpaper}

var (
	velmaAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	velmaBirthDay      = villagersBirthDay{14}
	velmaBirthMonth    = villagersBirthMonth{time.January} // 1
	velmaBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	velmaCategory      = villagersCategory{villagerCategoryB}
	velmaClothes       = villagersClothes{} // 2784
	velmaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorPurple}}
	velmaFlooring      = villagersFlooring{villagerFlooringLightParquetFlooring}
	velmaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGrandPiano, villagerFurniturePianoBench, villagerFurnitureAntiqueVanity, villagerFurnitureMusicStand, villagerFurnitureFancyViolin, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueMiniTable, villagerFurnitureAntiqueConsoleTable, villagerFurniturePhonograph, villagerFurnitureMetronome}}
	velmaGames         = villagersGames{[]VillagerGame{}} // TBD
	velmaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	velmaInterest      = villagersInterest{villagerInterestEducation}
	velmaName          = villagersName{villagerNameVelma}
	velmaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	velmaMusic         = villagersMusic{villagerMusicMrKK}
	velmaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	velmaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGoat}}
	velmaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	velmaWallpaper     = villagersWallpaper{villagerWallpaperClassicLibraryWall}
)
