package atsumori

import "time"

// Octavian is an Animal Crossing villager.
var Octavian = villager{
	octavianAstrology,
	octavianBirthDay,
	octavianBirthMonth,
	octavianBubbleColor,
	octavianCategory,
	octavianClothes,
	octavianClothesColors,
	octavianFlooring,
	octavianFurniture,
	octavianGames,
	octavianGender,
	octavianInterest,
	octavianName,
	octavianNameColor,
	octavianMusic,
	octavianPersonality,
	octavianSpecies,
	octavianStyle,
	octavianWallpaper}

var (
	octavianAstrology     = villagersAstrology{villagerAstrologyVirgo}
	octavianBirthDay      = villagersBirthDay{20}
	octavianBirthMonth    = villagersBirthMonth{time.September} // 9
	octavianBubbleColor   = villagersBubbleColor{villagerBubbleColorFF4040}
	octavianCategory      = villagersCategory{villagerCategoryB}
	octavianClothes       = villagersClothes{villagerClothesGoldPrintTee} // 4327
	octavianClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorWhite}}
	octavianFlooring      = villagersFlooring{villagerFlooringLunarSurface}
	octavianFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAsteroid, villagerFurnitureCrewedSpaceship, villagerFurnitureRocket, villagerFurnitureFlyingSaucer, villagerFurnitureAstronautSuit, villagerFurnitureSatellite}}
	octavianGames         = villagersGames{[]VillagerGame{}} // TBD
	octavianGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	octavianInterest      = villagersInterest{villagerInterestPlay}
	octavianName          = villagersName{villagerNameOctavian}
	octavianNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	octavianMusic         = villagersMusic{villagerMusicKKDB}
	octavianPersonality   = villagersPersonality{villagerPersonalityCranky}
	octavianSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOctopus}}
	octavianStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	octavianWallpaper     = villagersWallpaper{villagerWallpaperStarrySkyWall}
)
