package atsumori

import "time"

// Naomi is an Animal Crossing villager.
var Naomi = villager{
	naomiAstrology,
	naomiBirthDay,
	naomiBirthMonth,
	naomiBubbleColor,
	naomiCategory,
	naomiClothes,
	naomiClothesColors,
	naomiFlooring,
	naomiFurniture,
	naomiGames,
	naomiGender,
	naomiInterest,
	naomiName,
	naomiNameColor,
	naomiMusic,
	naomiPersonality,
	naomiSpecies,
	naomiStyle,
	naomiWallpaper}

var (
	naomiAstrology     = villagersAstrology{villagerAstrologyPisces}
	naomiBirthDay      = villagersBirthDay{28}
	naomiBirthMonth    = villagersBirthMonth{time.February} // 2
	naomiBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	naomiCategory      = villagersCategory{villagerCategoryA}
	naomiClothes       = villagersClothes{} // 7926
	naomiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorPurple}}
	naomiFlooring      = villagersFlooring{villagerFlooringLobbyFlooring}
	naomiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePhonograph, villagerFurnitureAntiqueBed, villagerFurnitureAntiqueConsoleTable, villagerFurnitureAntiqueBureau, villagerFurnitureAntiqueWardrobe, villagerFurnitureAntiqueChair, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueClock, villagerFurnitureAntiqueVanity, villagerFurnitureAntiquePhone}}
	naomiGames         = villagersGames{[]VillagerGame{}} // TBD
	naomiGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	naomiInterest      = villagersInterest{villagerInterestFashion}
	naomiName          = villagersName{villagerNameNaomi}
	naomiNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	naomiMusic         = villagersMusic{villagerMusicKKMilonga}
	naomiPersonality   = villagersPersonality{villagerPersonalitySnooty}
	naomiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCow}}
	naomiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	naomiWallpaper     = villagersWallpaper{villagerWallpaperRedArtDecoWall}
)
