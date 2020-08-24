package atsumori

import "time"

// Flip is an Animal Crossing villager
var Flip = villager{
	flipAstrology,
	flipBirthDay,
	flipBirthMonth,
	flipBubbleColor,
	flipCategory,
	flipClothes,
	flipClothesColors,
	flipFlooring,
	flipFurniture,
	flipGames,
	flipGender,
	flipInterest,
	flipName,
	flipNameColor,
	flipMusic,
	flipPersonality,
	flipSpecies,
	flipStyle,
	flipWallpaper}

var (
	flipAstrology     = villagersAstrology{villagerAstrologyScorpio}
	flipBirthDay      = villagersBirthDay{21}
	flipBirthMonth    = villagersBirthMonth{time.November} // 11
	flipBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	flipCategory      = villagersCategory{villagerCategoryA}
	flipClothes       = villagersClothes{villagerClothesMuscleTank} // 8058
	flipClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorYellow}}
	flipFlooring      = villagersFlooring{villagerFlooringGravelFlooring}
	flipFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBambooSpeaker, villagerFurniturePhoneBox, villagerFurnitureDrinkMachine, villagerFurniturePortableToilet, villagerFurnitureBlossomViewingLantern, villagerFurnitureBlossomViewingLantern, villagerFurnitureBambooBench, villagerFurnitureLectureHallDesk, villagerFurnitureEspressoMaker, villagerFurnitureDishDryingRack, villagerFurnitureUtilitySink, villagerFurnitureBambooStool, villagerFurnitureStall, villagerFurnitureShavedIceMaker}}
	flipGames         = villagersGames{[]VillagerGame{}} // TBD
	flipGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	flipInterest      = villagersInterest{villagerInterestMusic}
	flipName          = villagersName{villagerNameFlip}
	flipNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	flipMusic         = villagersMusic{villagerMusicKKRally}
	flipPersonality   = villagersPersonality{villagerPersonalityJock}
	flipSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMonkey}}
	flipStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	flipWallpaper     = villagersWallpaper{villagerWallpaperCherryBlossomTreesWall}
)
