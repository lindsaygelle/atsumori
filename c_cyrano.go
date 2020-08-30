package atsumori

import "time"

// Cyrano is an Animal Crossing villager.
var Cyrano = villager{
	cyranoAstrology,
	cyranoBirthDay,
	cyranoBirthMonth,
	cyranoBubbleColor,
	cyranoCategory,
	cyranoClothes,
	cyranoClothesColors,
	cyranoFlooring,
	cyranoFurniture,
	cyranoGames,
	cyranoGender,
	cyranoInterest,
	cyranoName,
	cyranoNameColor,
	cyranoMusic,
	cyranoPersonality,
	cyranoSpecies,
	cyranoStyle,
	cyranoWallpaper}

var (
	cyranoAstrology     = villagersAstrology{villagerAstrologyPisces}
	cyranoBirthDay      = villagersBirthDay{9}
	cyranoBirthMonth    = villagersBirthMonth{time.March} // 3
	cyranoBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	cyranoCategory      = villagersCategory{villagerCategoryB}
	cyranoClothes       = villagersClothes{} // 4249
	cyranoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorBeige}}
	cyranoFlooring      = villagersFlooring{villagerFlooringTatami}
	cyranoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHearth, villagerFurnitureFuton, villagerFurnitureBambooBasket, villagerFurnitureBambooStopblock, villagerFurnitureBambooCandleholder, villagerFurniturePileOfZenCushions, villagerFurnitureZenCushion, villagerFurnitureScreen}}
	cyranoGames         = villagersGames{[]VillagerGame{}} // TBD
	cyranoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	cyranoInterest      = villagersInterest{villagerInterestEducation}
	cyranoName          = villagersName{villagerNameCyrano}
	cyranoNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	cyranoMusic         = villagersMusic{villagerMusicKKLament}
	cyranoPersonality   = villagersPersonality{villagerPersonalityCranky}
	cyranoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAnteater}}
	cyranoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	cyranoWallpaper     = villagersWallpaper{villagerWallpaperShojiScreen}
)
