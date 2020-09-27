package atsumori

import "time"

// Simon is an Animal Crossing villager.
var Simon = villager{
	simonAstrology,
	simonBirthDay,
	simonBirthMonth,
	simonBubbleColor,
	simonCategory,
	simonClothes,
	simonClothesColors,
	simonFlooring,
	simonFurniture,
	simonGames,
	simonGender,
	simonInterest,
	simonName,
	simonNameColor,
	simonMusic,
	simonPersonality,
	simonSpecies,
	simonStyle,
	simonWallpaper}

var (
	simonAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	simonBirthDay      = villagersBirthDay{19}
	simonBirthMonth    = villagersBirthMonth{time.January} // 1
	simonBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	simonCategory      = villagersCategory{villagerCategoryB}
	simonClothes       = villagersClothes{villagerClothesStripedTee} // 7768
	simonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorColorful}}
	simonFlooring      = villagersFlooring{villagerFlooringSaharahsDesert}
	simonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureBonfire, villagerFurnitureCampfireCookware, villagerFurnitureSmoker, villagerFurnitureCassettePlayer, villagerFurnitureTikiTorch, villagerFurnitureHammock}}
	simonGames         = villagersGames{[]VillagerGame{}} // TBD
	simonGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	simonInterest      = villagersInterest{villagerInterestPlay}
	simonName          = villagersName{villagerNameSimon}
	simonNameColor     = villagersNameColor{villagerNameColor8B5F57}
	simonMusic         = villagersMusic{villagerMusicKKSafari}
	simonPersonality   = villagersPersonality{villagerPersonalityLazy}
	simonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMonkey}}
	simonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	simonWallpaper     = villagersWallpaper{villagerWallpaperStarrySkyWall}
)
