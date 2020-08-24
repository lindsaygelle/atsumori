package atsumori

import "time"

// Frita is an Animal Crossing villager
var Frita = villager{
	fritaAstrology,
	fritaBirthDay,
	fritaBirthMonth,
	fritaBubbleColor,
	fritaCategory,
	fritaClothes,
	fritaClothesColors,
	fritaFlooring,
	fritaFurniture,
	fritaGames,
	fritaGender,
	fritaInterest,
	fritaName,
	fritaNameColor,
	fritaMusic,
	fritaPersonality,
	fritaSpecies,
	fritaStyle,
	fritaWallpaper}

var (
	fritaAstrology     = villagersAstrology{villagerAstrologyCancer}
	fritaBirthDay      = villagersBirthDay{16}
	fritaBirthMonth    = villagersBirthMonth{time.July} // 7
	fritaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	fritaCategory      = villagersCategory{villagerCategoryA}
	fritaClothes       = villagersClothes{} // 5287
	fritaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorRed}}
	fritaFlooring      = villagersFlooring{villagerFlooringYellowFloralFlooring}
	fritaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJukebox, villagerFurnitureDrinkMachine, villagerFurnitureSnackMachine, villagerFurnitureCandyMachine, villagerFurnitureMenuChalkboard, villagerFurnitureDinerNeonSign, villagerFurnitureDinerDiningTable, villagerFurnitureCoffeeCup, villagerFurnitureDinerCounterTable, villagerFurnitureEspressoMaker, villagerFurnitureSoftServeLamp}}
	fritaGames         = villagersGames{[]VillagerGame{}} // TBD
	fritaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	fritaInterest      = villagersInterest{villagerInterestMusic}
	fritaName          = villagersName{villagerNameFrita}
	fritaNameColor     = villagersNameColor{villagerNameColor874C25}
	fritaMusic         = villagersMusic{villagerMusicKKRockabilly}
	fritaPersonality   = villagersPersonality{villagerPersonalityBigSister}
	fritaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	fritaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCute}}
	fritaWallpaper     = villagersWallpaper{villagerWallpaperPinkDinerWall}
)
