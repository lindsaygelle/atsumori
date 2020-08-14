package atsumori

import "time"

// Biff is an Animal Crossing villager
var Biff = villager{
	biffAstrology,
	biffBirthDay,
	biffBirthMonth,
	biffBubbleColor,
	biffCategory,
	biffClothes,
	biffClothesColors,
	biffFlooring,
	biffFurniture,
	biffGames,
	biffGender,
	biffInterest,
	biffName,
	biffNameColor,
	biffMusic,
	biffPersonality,
	biffSpecies,
	biffStyle,
	biffWallpaper}

var (
	biffAstrology     = villagersAstrology{villagerAstrologyAries}
	biffBirthDay      = villagersBirthDay{29}
	biffBirthMonth    = villagersBirthMonth{time.March} // 3
	biffBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	biffCategory      = villagersCategory{villagerCategoryB}
	biffClothes       = villagersClothes{villagerClothesGoldPrintTee} // 4327
	biffClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorBlue}}
	biffFlooring      = villagersFlooring{villagerFlooringPaintballFlooring}
	biffFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBasketballHoop, villagerFurnitureBall, villagerFurnitureUtilityPole, villagerFurnitureElectricKickScooter, villagerFurniturePaintingSet, villagerFurniturePhoneBox, villagerFurnitureMountainBike, villagerFurnitureTapeDeck}}
	biffGames         = villagersGames{[]VillagerGame{}} // TBD
	biffGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	biffInterest      = villagersInterest{villagerInterestFitness}
	biffName          = villagersName{villagerNameBiff}
	biffNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	biffMusic         = villagersMusic{} // The K. Funk
	biffPersonality   = villagersPersonality{villagerPersonalityJock}
	biffSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHippo}}
	biffStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleActive}}
	biffWallpaper     = villagersWallpaper{villagerWallpaperStreetArtWall}
)
