package atsumori

import "time"

// Spike is an Animal Crossing villager.
var Spike = villager{
	spikeAstrology,
	spikeBirthDay,
	spikeBirthMonth,
	spikeBubbleColor,
	spikeCategory,
	spikeClothes,
	spikeClothesColors,
	spikeFlooring,
	spikeFurniture,
	spikeGames,
	spikeGender,
	spikeInterest,
	spikeName,
	spikeNameColor,
	spikeMusic,
	spikePersonality,
	spikeSpecies,
	spikeStyle,
	spikeWallpaper}

var (
	spikeAstrology     = villagersAstrology{villagerAstrologyGemini}
	spikeBirthDay      = villagersBirthDay{17}
	spikeBirthMonth    = villagersBirthMonth{time.June} // 6
	spikeBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	spikeCategory      = villagersCategory{villagerCategoryA}
	spikeClothes       = villagersClothes{villagerClothesGoldPrintTee} // 4327
	spikeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	spikeFlooring      = villagersFlooring{villagerFlooringLeopardPrintFlooring}
	spikeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShowerBooth, villagerFurnitureRattanBed, villagerFurnitureDinerSofa, villagerFurniturePullUpBarStand, villagerFurnitureJukebox, villagerFurnitureMountainBike, villagerFurnitureRattanEndTable, villagerFurnitureKettlebell, villagerFurnitureWallMountedTV20In, villagerFurnitureIronWallRack, villagerFurnitureOpenFrameKitchen, villagerFurnitureProteinShakerBottle}}
	spikeGames         = villagersGames{[]VillagerGame{}} // TBD
	spikeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	spikeInterest      = villagersInterest{villagerInterestNature}
	spikeName          = villagersName{villagerNameSpike}
	spikeNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	spikeMusic         = villagersMusic{villagerMusicKKHouse}
	spikePersonality   = villagersPersonality{villagerPersonalityCranky}
	spikeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRhino}}
	spikeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	spikeWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
