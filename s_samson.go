package atsumori

import "time"

// Samson is an Animal Crossing villager.
var Samson = villager{
	samsonAstrology,
	samsonBirthDay,
	samsonBirthMonth,
	samsonBubbleColor,
	samsonCategory,
	samsonClothes,
	samsonClothesColors,
	samsonFlooring,
	samsonFurniture,
	samsonGames,
	samsonGender,
	samsonInterest,
	samsonName,
	samsonNameColor,
	samsonMusic,
	samsonPersonality,
	samsonSpecies,
	samsonStyle,
	samsonWallpaper}

var (
	samsonAstrology     = villagersAstrology{villagerAstrologyCancer}
	samsonBirthDay      = villagersBirthDay{5}
	samsonBirthMonth    = villagersBirthMonth{time.July} // 7
	samsonBubbleColor   = villagersBubbleColor{villagerBubbleColorACC8CF}
	samsonCategory      = villagersCategory{villagerCategoryB}
	samsonClothes       = villagersClothes{villagerClothesKanjiTee} // 8426
	samsonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorYellow}}
	samsonFlooring      = villagersFlooring{villagerFlooringCoolPaintFlooring}
	samsonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePearBed, villagerFurniturePearWardrobe, villagerFurnitureFoosballTable, villagerFurnitureThrowbackWallClock, villagerFurnitureJuicyAppleTV, villagerFurnitureOrangeEndTable, villagerFurnitureTrainSet, villagerFurnitureThrowbackWrestlingFigure, villagerFurnitureThrowbackContainer, villagerFurnitureProTapeRecorder}}
	samsonGames         = villagersGames{[]VillagerGame{}} // TBD
	samsonGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	samsonInterest      = villagersInterest{villagerInterestFitness}
	samsonName          = villagersName{villagerNameSamson}
	samsonNameColor     = villagersNameColor{villagerNameColor498992}
	samsonMusic         = villagersMusic{villagerMusicRockinKK}
	samsonPersonality   = villagersPersonality{villagerPersonalityJock}
	samsonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	samsonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	samsonWallpaper     = villagersWallpaper{villagerWallpaperBlueSubwayTileWall}
)
