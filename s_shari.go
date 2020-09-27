package atsumori

import "time"

// Shari is an Animal Crossing villager.
var Shari = villager{
	shariAstrology,
	shariBirthDay,
	shariBirthMonth,
	shariBubbleColor,
	shariCategory,
	shariClothes,
	shariClothesColors,
	shariFlooring,
	shariFurniture,
	shariGames,
	shariGender,
	shariInterest,
	shariName,
	shariNameColor,
	shariMusic,
	shariPersonality,
	shariSpecies,
	shariStyle,
	shariWallpaper}

var (
	shariAstrology     = villagersAstrology{villagerAstrologyAries}
	shariBirthDay      = villagersBirthDay{10}
	shariBirthMonth    = villagersBirthMonth{time.April} // 4
	shariBubbleColor   = villagersBubbleColor{villagerBubbleColorBFBFBF}
	shariCategory      = villagersCategory{villagerCategoryA}
	shariClothes       = villagersClothes{villagerClothesGardenTank} // 8075
	shariClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorYellow}}
	shariFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	shariFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanLowTable, villagerFurnitureRattanArmchair, villagerFurnitureRattanWasteBin, villagerFurnitureRattanVanity, villagerFurnitureHangingTerrarium, villagerFurnitureRattanEndTable, villagerFurnitureRattanWardrobe, villagerFurnitureRattanTowelBasket, villagerFurnitureLongBathtub, villagerFurnitureRoseBed, villagerFurnitureIronShelf}}
	shariGames         = villagersGames{[]VillagerGame{}} // TBD
	shariGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	shariInterest      = villagersInterest{villagerInterestMusic}
	shariName          = villagersName{villagerNameShari}
	shariNameColor     = villagersNameColor{villagerNameColor5E5E5E}
	shariMusic         = villagersMusic{villagerMusicKKCruisin}
	shariPersonality   = villagersPersonality{villagerPersonalityBigSister}
	shariSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMonkey}}
	shariStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleActive}}
	shariWallpaper     = villagersWallpaper{villagerWallpaperPurpleDesertTileWall}
)
