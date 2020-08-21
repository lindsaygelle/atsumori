package atsumori

import "time"

// Buzz is an Animal Crossing villager
var Buzz = villager{
	buzzAstrology,
	buzzBirthDay,
	buzzBirthMonth,
	buzzBubbleColor,
	buzzCategory,
	buzzClothes,
	buzzClothesColors,
	buzzFlooring,
	buzzFurniture,
	buzzGames,
	buzzGender,
	buzzInterest,
	buzzName,
	buzzNameColor,
	buzzMusic,
	buzzPersonality,
	buzzSpecies,
	buzzStyle,
	buzzWallpaper}

var (
	buzzAstrology     = villagersAstrology{villagerAstrologySagittarius}
	buzzBirthDay      = villagersBirthDay{7}
	buzzBirthMonth    = villagersBirthMonth{time.December} // 12
	buzzBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	buzzCategory      = villagersCategory{villagerCategoryA}
	buzzClothes       = villagersClothes{villagerClothesNineBallTee} // 3326
	buzzClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorRed}}
	buzzFlooring      = villagersFlooring{villagerFlooringParkingFlooring}
	buzzFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureOutdoorTable, villagerFurnitureGarbagePail, villagerFurnitureDrinkMachine, villagerFurnitureSnackMachine, villagerFurnitureDrinkMachine, villagerFurnitureGarbageBin, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureOutdoorBench, villagerFurnitureCassettePlayer}}
	buzzGames         = villagersGames{[]VillagerGame{}} // TBD
	buzzGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	buzzInterest      = villagersInterest{villagerInterestNature}
	buzzName          = villagersName{villagerNameBuzz}
	buzzNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	buzzMusic         = villagersMusic{villagerMusicDrivin}
	buzzPersonality   = villagersPersonality{villagerPersonalityCranky}
	buzzSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	buzzStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	buzzWallpaper     = villagersWallpaper{villagerWallpaperSummitWall}
)
