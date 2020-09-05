package atsumori

import "time"

// Monty is an Animal Crossing villager.
var Monty = villager{
	montyAstrology,
	montyBirthDay,
	montyBirthMonth,
	montyBubbleColor,
	montyCategory,
	montyClothes,
	montyClothesColors,
	montyFlooring,
	montyFurniture,
	montyGames,
	montyGender,
	montyInterest,
	montyName,
	montyNameColor,
	montyMusic,
	montyPersonality,
	montySpecies,
	montyStyle,
	montyWallpaper}

var (
	montyAstrology     = villagersAstrology{villagerAstrologySagittarius}
	montyBirthDay      = villagersBirthDay{7}
	montyBirthMonth    = villagersBirthMonth{time.December} // 12
	montyBubbleColor   = villagersBubbleColor{villagerBubbleColorACC8CF}
	montyCategory      = villagersCategory{villagerCategoryB}
	montyClothes       = villagersClothes{villagerClothesGuayaberaShirt} // 3687
	montyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorGray}}
	montyFlooring      = villagersFlooring{villagerFlooringForestFlooring}
	montyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBambooSpeaker, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureMushLog, villagerFurnitureMushParasol, villagerFurnitureMushTable, villagerFurnitureMushLowStool, villagerFurnitureMushLamp, villagerFurnitureMushPartition}}
	montyGames         = villagersGames{[]VillagerGame{}} // TBD
	montyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	montyInterest      = villagersInterest{villagerInterestEducation}
	montyName          = villagersName{villagerNameMonty}
	montyNameColor     = villagersNameColor{villagerNameColor498992}
	montyMusic         = villagersMusic{villagerMusicToTheEdge}
	montyPersonality   = villagersPersonality{villagerPersonalityCranky}
	montySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMonkey}}
	montyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	montyWallpaper     = villagersWallpaper{villagerWallpaperAutumnWall}
)
