package atsumori

import "time"

// Scoot is an Animal Crossing villager.
var Scoot = villager{
	scootAstrology,
	scootBirthDay,
	scootBirthMonth,
	scootBubbleColor,
	scootCategory,
	scootClothes,
	scootClothesColors,
	scootFlooring,
	scootFurniture,
	scootGames,
	scootGender,
	scootInterest,
	scootName,
	scootNameColor,
	scootMusic,
	scootPersonality,
	scootSpecies,
	scootStyle,
	scootWallpaper}

var (
	scootAstrology     = villagersAstrology{villagerAstrologyGemini}
	scootBirthDay      = villagersBirthDay{13}
	scootBirthMonth    = villagersBirthMonth{time.June} // 6
	scootBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	scootCategory      = villagersCategory{villagerCategoryB}
	scootClothes       = villagersClothes{villagerClothesFrogTee} // 8195
	scootClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorBlue}}
	scootFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	scootFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHoseReel, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureBeachTowel, villagerFurnitureGardenFaucet, villagerFurniturePlasticPool, villagerFurnitureLogBench, villagerFurniturePortableToilet, villagerFurniturePoolsideBed, villagerFurnitureHandyWaterCooler, villagerFurnitureCuteMusicPlayer}}
	scootGames         = villagersGames{[]VillagerGame{}} // TBD
	scootGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	scootInterest      = villagersInterest{villagerInterestFitness}
	scootName          = villagersName{villagerNameScoot}
	scootNameColor     = villagersNameColor{villagerNameColor28665A}
	scootMusic         = villagersMusic{villagerMusicMyPlace}
	scootPersonality   = villagersPersonality{villagerPersonalityJock}
	scootSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	scootStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	scootWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
