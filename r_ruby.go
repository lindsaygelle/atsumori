package atsumori

import "time"

// Ruby is an Animal Crossing villager.
var Ruby = villager{
	rubyAstrology,
	rubyBirthDay,
	rubyBirthMonth,
	rubyBubbleColor,
	rubyCategory,
	rubyClothes,
	rubyClothesColors,
	rubyFlooring,
	rubyFurniture,
	rubyGames,
	rubyGender,
	rubyInterest,
	rubyName,
	rubyNameColor,
	rubyMusic,
	rubyPersonality,
	rubySpecies,
	rubyStyle,
	rubyWallpaper}

var (
	rubyAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	rubyBirthDay      = villagersBirthDay{25}
	rubyBirthMonth    = villagersBirthMonth{time.December} // 12
	rubyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	rubyCategory      = villagersCategory{villagerCategoryB}
	rubyClothes       = villagersClothes{villagerClothesRabbitTee} // 8197
	rubyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorPink}}
	rubyFlooring      = villagersFlooring{villagerFlooringLunarSurface}
	rubyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAsteroid, villagerFurnitureAsteroid, villagerFurnitureAstronautSuit, villagerFurnitureAstronautSuit, villagerFurnitureCrewedSpaceship, villagerFurnitureMoon, villagerFurnitureLunarRover}}
	rubyGames         = villagersGames{[]VillagerGame{}} // TBD
	rubyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	rubyInterest      = villagersInterest{villagerInterestNature}
	rubyName          = villagersName{villagerNameRuby}
	rubyNameColor     = villagersNameColor{villagerNameColor848484}
	rubyMusic         = villagersMusic{villagerMusicStaleCupcakes}
	rubyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	rubySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	rubyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleActive}}
	rubyWallpaper     = villagersWallpaper{villagerWallpaperStarrySkyWall}
)
