package atsumori

import "time"

// Rory is an Animal Crossing villager.
var Rory = villager{
	roryAstrology,
	roryBirthDay,
	roryBirthMonth,
	roryBubbleColor,
	roryCategory,
	roryClothes,
	roryClothesColors,
	roryFlooring,
	roryFurniture,
	roryGames,
	roryGender,
	roryInterest,
	roryName,
	roryNameColor,
	roryMusic,
	roryPersonality,
	rorySpecies,
	roryStyle,
	roryWallpaper}

var (
	roryAstrology     = villagersAstrology{villagerAstrologyLeo}
	roryBirthDay      = villagersBirthDay{7}
	roryBirthMonth    = villagersBirthMonth{time.August} // 8
	roryBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	roryCategory      = villagersCategory{villagerCategoryA}
	roryClothes       = villagersClothes{villagerClothesSeaHantenShirt} // 9218
	roryClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorRed}}
	roryFlooring      = villagersFlooring{villagerFlooringStarrySandsFlooring}
	roryFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureStoneLionDog, villagerFurnitureStoneLionDog, villagerFurniturePlasticPool, villagerFurnitureBeachBall, villagerFurniturePoolsideBed, villagerFurnitureBeachTowel, villagerFurnitureShellTable, villagerFurnitureShellSpeaker, villagerFurniturePortableToilet}}
	roryGames         = villagersGames{[]VillagerGame{}} // TBD
	roryGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	roryInterest      = villagersInterest{villagerInterestMusic}
	roryName          = villagersName{villagerNameRory}
	roryNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	roryMusic         = villagersMusic{villagerMusicKKFaire}
	roryPersonality   = villagersPersonality{villagerPersonalityJock}
	rorySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesLion}}
	roryStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	roryWallpaper     = villagersWallpaper{villagerWallpaperMangroveWall}
)
