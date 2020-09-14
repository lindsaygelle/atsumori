package atsumori

import "time"

// Portia is an Animal Crossing villager.
var Portia = villager{
	portiaAstrology,
	portiaBirthDay,
	portiaBirthMonth,
	portiaBubbleColor,
	portiaCategory,
	portiaClothes,
	portiaClothesColors,
	portiaFlooring,
	portiaFurniture,
	portiaGames,
	portiaGender,
	portiaInterest,
	portiaName,
	portiaNameColor,
	portiaMusic,
	portiaPersonality,
	portiaSpecies,
	portiaStyle,
	portiaWallpaper}

var (
	portiaAstrology     = villagersAstrology{villagerAstrologyScorpio}
	portiaBirthDay      = villagersBirthDay{25}
	portiaBirthMonth    = villagersBirthMonth{time.October} // 10
	portiaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	portiaCategory      = villagersCategory{villagerCategoryB}
	portiaClothes       = villagersClothes{} // 7902
	portiaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlack}}
	portiaFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	portiaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDeluxeWasher, villagerFurnitureRoseBed, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureAntiqueVanity, villagerFurnitureFloralSwag, villagerFurnitureBathroomTowelRack, villagerFurnitureIronwoodCupboard, villagerFurnitureStandMixer, villagerFurnitureAnalogKitchenScale, villagerFurnitureKitchenIsland, villagerFurnitureMixer, villagerFurniturePotRack, villagerFurnitureIronWorktable, villagerFurnitureEspressoMaker, villagerFurniturePortableRecordPlayer}}
	portiaGames         = villagersGames{[]VillagerGame{}} // TBD
	portiaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	portiaInterest      = villagersInterest{villagerInterestFashion}
	portiaName          = villagersName{villagerNamePortia}
	portiaNameColor     = villagersNameColor{villagerNameColor848484}
	portiaMusic         = villagersMusic{villagerMusicKKDisco}
	portiaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	portiaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	portiaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	portiaWallpaper     = villagersWallpaper{villagerWallpaperBlueDesertTileWall}
)
