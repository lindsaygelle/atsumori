package atsumori

import "time"

// Cally is an Animal Crossing villager.
var Cally = villager{
	callyAstrology,
	callyBirthDay,
	callyBirthMonth,
	callyBubbleColor,
	callyCategory,
	callyClothes,
	callyClothesColors,
	callyFlooring,
	callyFurniture,
	callyGames,
	callyGender,
	callyInterest,
	callyName,
	callyNameColor,
	callyMusic,
	callyPersonality,
	callySpecies,
	callyStyle,
	callyWallpaper}

var (
	callyAstrology     = villagersAstrology{villagerAstrologyVirgo}
	callyBirthDay      = villagersBirthDay{4}
	callyBirthMonth    = villagersBirthMonth{time.September} // 9
	callyBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	callyCategory      = villagersCategory{villagerCategoryB}
	callyClothes       = villagersClothes{} // 3137
	callyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorGreen}}
	callyFlooring      = villagersFlooring{villagerFlooringDaisyMeadow}
	callyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLilyRecordPlayer, villagerFurnitureGardenLantern, villagerFurnitureDrinkingFountain, villagerFurnitureSwingingBench, villagerFurnitureHammock, villagerFurnitureTreesBountyBigTree, villagerFurnitureGardenGnome, villagerFurnitureBirdhouse, villagerFurnitureNaturalGardenTable, villagerFurniturePicnicBasket}}
	callyGames         = villagersGames{[]VillagerGame{}} // TBD
	callyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	callyInterest      = villagersInterest{villagerInterestNature}
	callyName          = villagersName{villagerNameCally}
	callyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	callyMusic         = villagersMusic{villagerMusicKKStroll}
	callyPersonality   = villagersPersonality{villagerPersonalityNormal}
	callySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	callyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	callyWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
