package atsumori

import "time"

// Victoria is an Animal Crossing villager.
var Victoria = villager{
	victoriaAstrology,
	victoriaBirthDay,
	victoriaBirthMonth,
	victoriaBubbleColor,
	victoriaCategory,
	victoriaClothes,
	victoriaClothesColors,
	victoriaFlooring,
	victoriaFurniture,
	victoriaGames,
	victoriaGender,
	victoriaInterest,
	victoriaName,
	victoriaNameColor,
	victoriaMusic,
	victoriaPersonality,
	victoriaSpecies,
	victoriaStyle,
	victoriaWallpaper}

var (
	victoriaAstrology     = villagersAstrology{villagerAstrologyCancer}
	victoriaBirthDay      = villagersBirthDay{11}
	victoriaBirthMonth    = villagersBirthMonth{time.July} // 7
	victoriaBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	victoriaCategory      = villagersCategory{villagerCategoryB}
	victoriaClothes       = villagersClothes{villagerClothesOneBallTee} // 3320
	victoriaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorOrange}}
	victoriaFlooring      = villagersFlooring{villagerFlooringRacetrackFlooring}
	victoriaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHayBed, villagerFurnitureChampionsPennant, villagerFurnitureSpringyRideOn, villagerFurnitureSpringyRideOn, villagerFurnitureSpringyRideOn, villagerFurnitureSpringyRideOn, villagerFurnitureSpringyRideOn, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurnitureSpringyRideOn, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurnitureSpringyRideOn, villagerFurniturePartyGarland, villagerFurniturePartyGarland}}
	victoriaGames         = villagersGames{[]VillagerGame{}} // TBD
	victoriaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	victoriaInterest      = villagersInterest{villagerInterestFitness}
	victoriaName          = villagersName{villagerNameVictoria}
	victoriaNameColor     = villagersNameColor{villagerNameColor8B5F57}
	victoriaMusic         = villagersMusic{villagerMusicKKAdventure}
	victoriaPersonality   = villagersPersonality{villagerPersonalityPeppy}
	victoriaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	victoriaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	victoriaWallpaper     = villagersWallpaper{villagerWallpaperStadiumWall}
)
