package atsumori

import "time"

// Jitters is an Animal Crossing villager.
var Jitters = villager{
	jittersAstrology,
	jittersBirthDay,
	jittersBirthMonth,
	jittersBubbleColor,
	jittersCategory,
	jittersClothes,
	jittersClothesColors,
	jittersFlooring,
	jittersFurniture,
	jittersGames,
	jittersGender,
	jittersInterest,
	jittersName,
	jittersNameColor,
	jittersMusic,
	jittersPersonality,
	jittersSpecies,
	jittersStyle,
	jittersWallpaper}

var (
	jittersAstrology     = villagersAstrology{villagerAstrologyAquarius}
	jittersBirthDay      = villagersBirthDay{2}
	jittersBirthMonth    = villagersBirthMonth{time.February} // 2
	jittersBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	jittersCategory      = villagersCategory{villagerCategoryB}
	jittersClothes       = villagersClothes{villagerClothesSoccerUniformTop} // 3249
	jittersClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorOrange}}
	jittersFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	jittersFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBall, villagerFurnitureBall, villagerFurnitureChampionsPennant, villagerFurnitureDirectorsChair, villagerFurnitureSoccerGoal}}
	jittersGames         = villagersGames{[]VillagerGame{}} // TBD
	jittersGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	jittersInterest      = villagersInterest{villagerInterestFitness}
	jittersName          = villagersName{villagerNameJitters}
	jittersNameColor     = villagersNameColor{villagerNameColor28665A}
	jittersMusic         = villagersMusic{villagerMusicKKAdventure}
	jittersPersonality   = villagersPersonality{villagerPersonalityJock}
	jittersSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	jittersStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleActive}}
	jittersWallpaper     = villagersWallpaper{villagerWallpaperStadiumWall}
)
