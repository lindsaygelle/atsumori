package atsumori

import "time"

// Hamphrey is an Animal Crossing villager.
var Hamphrey = villager{
	hamphreyAstrology,
	hamphreyBirthDay,
	hamphreyBirthMonth,
	hamphreyBubbleColor,
	hamphreyCategory,
	hamphreyClothes,
	hamphreyClothesColors,
	hamphreyFlooring,
	hamphreyFurniture,
	hamphreyGames,
	hamphreyGender,
	hamphreyInterest,
	hamphreyName,
	hamphreyNameColor,
	hamphreyMusic,
	hamphreyPersonality,
	hamphreySpecies,
	hamphreyStyle,
	hamphreyWallpaper}

var (
	hamphreyAstrology     = villagersAstrology{villagerAstrologyPisces}
	hamphreyBirthDay      = villagersBirthDay{25}
	hamphreyBirthMonth    = villagersBirthMonth{time.February} // 2
	hamphreyBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	hamphreyCategory      = villagersCategory{villagerCategoryA}
	hamphreyClothes       = villagersClothes{villagerClothesFuzzyVest} // 8514
	hamphreyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorBeige}}
	hamphreyFlooring      = villagersFlooring{villagerFlooringRushTatami}
	hamphreyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRaccoonFigurine, villagerFurnitureBonsaiShelf, villagerFurnitureRoundSpaceHeater, villagerFurniturePileOfZenCushions, villagerFurnitureFuton, villagerFurnitureTapeDeck, villagerFurnitureWoodenPlankSign, villagerFurnitureKotatsu}}
	hamphreyGames         = villagersGames{[]VillagerGame{}} // TBD
	hamphreyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	hamphreyInterest      = villagersInterest{villagerInterestNature}
	hamphreyName          = villagersName{villagerNameHamphrey}
	hamphreyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	hamphreyMusic         = villagersMusic{villagerMusicComradeKK}
	hamphreyPersonality   = villagersPersonality{villagerPersonalityCranky}
	hamphreySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHamster}}
	hamphreyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	hamphreyWallpaper     = villagersWallpaper{villagerWallpaperScreenWall}
)
