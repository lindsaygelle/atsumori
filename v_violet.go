package atsumori

import "time"

// Violet is an Animal Crossing villager.
var Violet = villager{
	violetAstrology,
	violetBirthDay,
	violetBirthMonth,
	violetBubbleColor,
	violetCategory,
	violetClothes,
	violetClothesColors,
	violetFlooring,
	violetFurniture,
	violetGames,
	violetGender,
	violetInterest,
	violetName,
	violetNameColor,
	violetMusic,
	violetPersonality,
	violetSpecies,
	violetStyle,
	violetWallpaper}

var (
	violetAstrology     = villagersAstrology{villagerAstrologyVirgo}
	violetBirthDay      = villagersBirthDay{1}
	violetBirthMonth    = villagersBirthMonth{time.September} // 9
	violetBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	violetCategory      = villagersCategory{villagerCategoryB}
	violetClothes       = villagersClothes{villagerClothesPonchoStyleSweater} // 4661
	violetClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorPink}}
	violetFlooring      = villagersFlooring{villagerFlooringLobbyFlooring}
	violetFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRoseBed, villagerFurnitureWhirlpoolBath, villagerFurnitureCherryBlossomBranches, villagerFurnitureDoubleSofa, villagerFurnitureRattanLowTable, villagerFurnitureFragranceSticks, villagerFurnitureShowerSet, villagerFurnitureBathroomTowelRack, villagerFurnitureAntiqueConsoleTable, villagerFurniturePortableRecordPlayer, villagerFurnitureFloralSwag}}
	violetGames         = villagersGames{[]VillagerGame{}} // TBD
	violetGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	violetInterest      = villagersInterest{villagerInterestFitness}
	violetName          = villagersName{villagerNameViolet}
	violetNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	violetMusic         = villagersMusic{villagerMusicKKBazaar}
	violetPersonality   = villagersPersonality{villagerPersonalitySnooty}
	violetSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGorilla}}
	violetStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCool}}
	violetWallpaper     = villagersWallpaper{villagerWallpaperBlackBotanicalTileWall}
)
