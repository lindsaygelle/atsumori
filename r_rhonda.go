package atsumori

import "time"

// Rhonda is an Animal Crossing villager.
var Rhonda = villager{
	rhondaAstrology,
	rhondaBirthDay,
	rhondaBirthMonth,
	rhondaBubbleColor,
	rhondaCategory,
	rhondaClothes,
	rhondaClothesColors,
	rhondaFlooring,
	rhondaFurniture,
	rhondaGames,
	rhondaGender,
	rhondaInterest,
	rhondaName,
	rhondaNameColor,
	rhondaMusic,
	rhondaPersonality,
	rhondaSpecies,
	rhondaStyle,
	rhondaWallpaper}

var (
	rhondaAstrology     = villagersAstrology{villagerAstrologyAquarius}
	rhondaBirthDay      = villagersBirthDay{24}
	rhondaBirthMonth    = villagersBirthMonth{time.January} // 1
	rhondaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	rhondaCategory      = villagersCategory{villagerCategoryB}
	rhondaClothes       = villagersClothes{} // 3385
	rhondaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlack}}
	rhondaFlooring      = villagersFlooring{villagerFlooringSimplePurpleFlooring}
	rhondaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShellBed, villagerFurnitureCuteVanity, villagerFurnitureCuteChair, villagerFurnitureClawFootTub, villagerFurnitureCuteSofa, villagerFurnitureShellPartition, villagerFurnitureShellSpeaker, villagerFurnitureShellArch, villagerFurnitureBathroomTowelRack, villagerFurnitureShowerSet}}
	rhondaGames         = villagersGames{[]VillagerGame{}} // TBD
	rhondaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	rhondaInterest      = villagersInterest{villagerInterestMusic}
	rhondaName          = villagersName{villagerNameRhonda}
	rhondaNameColor     = villagersNameColor{villagerNameColor848484}
	rhondaMusic         = villagersMusic{villagerMusicKKBallad}
	rhondaPersonality   = villagersPersonality{villagerPersonalityNormal}
	rhondaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRhino}}
	rhondaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	rhondaWallpaper     = villagersWallpaper{villagerWallpaperPurpleDesertTileWall}
)
