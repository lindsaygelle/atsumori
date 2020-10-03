package atsumori

import "time"

// Tucker is an Animal Crossing villager.
var Tucker = villager{
	tuckerAstrology,
	tuckerBirthDay,
	tuckerBirthMonth,
	tuckerBubbleColor,
	tuckerCategory,
	tuckerClothes,
	tuckerClothesColors,
	tuckerFlooring,
	tuckerFurniture,
	tuckerGames,
	tuckerGender,
	tuckerInterest,
	tuckerName,
	tuckerNameColor,
	tuckerMusic,
	tuckerPersonality,
	tuckerSpecies,
	tuckerStyle,
	tuckerWallpaper}

var (
	tuckerAstrology     = villagersAstrology{villagerAstrologyVirgo}
	tuckerBirthDay      = villagersBirthDay{7}
	tuckerBirthMonth    = villagersBirthMonth{time.September} // 9
	tuckerBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	tuckerCategory      = villagersCategory{villagerCategoryA}
	tuckerClothes       = villagersClothes{} // 9626
	tuckerClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorOrange}}
	tuckerFlooring      = villagersFlooring{villagerFlooringSwampFlooring}
	tuckerFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureTikiTorch, villagerFurnitureTikiTorch, villagerFurnitureBambooSpeaker, villagerFurnitureStonehenge, villagerFurnitureStoneStool, villagerFurnitureStoneStool, villagerFurnitureShantyMat, villagerFurnitureBambooDrum, villagerFurnitureStoneStool, villagerFurnitureClassicPitcher}}
	tuckerGames         = villagersGames{[]VillagerGame{}} // TBD
	tuckerGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	tuckerInterest      = villagersInterest{villagerInterestNature}
	tuckerName          = villagersName{villagerNameTucker}
	tuckerNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	tuckerMusic         = villagersMusic{villagerMusicKKSafari}
	tuckerPersonality   = villagersPersonality{villagerPersonalityLazy}
	tuckerSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	tuckerStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	tuckerWallpaper     = villagersWallpaper{villagerWallpaperJungleWall}
)
