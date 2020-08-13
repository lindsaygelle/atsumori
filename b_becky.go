package atsumori

import "time"

// Becky is an Animal Crossing villager
var Becky = villager{
	beckyAstrology,
	beckyBirthDay,
	beckyBirthMonth,
	beckyBubbleColor,
	beckyCategory,
	beckyClothes,
	beckyClothesColors,
	beckyFlooring,
	beckyFurniture,
	beckyGames,
	beckyGender,
	beckyInterest,
	beckyName,
	beckyNameColor,
	beckyMusic,
	beckyPersonality,
	beckySpecies,
	beckyStyle,
	beckyWallpaper}

var (
	beckyAstrology     = villagersAstrology{villagerAstrologySagittarius}
	beckyBirthDay      = villagersBirthDay{9}
	beckyBirthMonth    = villagersBirthMonth{time.December} // 12
	beckyBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	beckyCategory      = villagersCategory{villagerCategoryB}
	beckyClothes       = villagersClothes{} // 8173
	beckyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorPink}}
	beckyFlooring      = villagersFlooring{villagerFlooringPalaceTile}
	beckyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueClock, villagerFurnitureCello, villagerFurnitureAntiqueMiniTable, villagerFurniturePhonograph, villagerFurnitureFancyViolin, villagerFurnitureMusicStand, villagerFurnitureGrandPiano, villagerFurniturePianoBench, villagerFurnitureHarp}}
	beckyGames         = villagersGames{[]VillagerGame{}} // TBD
	beckyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	beckyInterest      = villagersInterest{villagerInterestMusic}
	beckyName          = villagersName{villagerNameBecky}
	beckyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	beckyMusic         = villagersMusic{} // K.K. Chorale
	beckyPersonality   = villagersPersonality{villagerPersonalitySnooty}
	beckySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	beckyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	beckyWallpaper     = villagersWallpaper{villagerWallpaperPalaceWall}
)
