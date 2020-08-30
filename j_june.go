package atsumori

import "time"

// June is an Animal Crossing villager.
var June = villager{
	juneAstrology,
	juneBirthDay,
	juneBirthMonth,
	juneBubbleColor,
	juneCategory,
	juneClothes,
	juneClothesColors,
	juneFlooring,
	juneFurniture,
	juneGames,
	juneGender,
	juneInterest,
	juneName,
	juneNameColor,
	juneMusic,
	junePersonality,
	juneSpecies,
	juneStyle,
	juneWallpaper}

var (
	juneAstrology     = villagersAstrology{villagerAstrologyGemini}
	juneBirthDay      = villagersBirthDay{21}
	juneBirthMonth    = villagersBirthMonth{time.May} // 5
	juneBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	juneCategory      = villagersCategory{villagerCategoryA}
	juneClothes       = villagersClothes{} // 8189
	juneClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorRed}}
	juneFlooring      = villagersFlooring{villagerFlooringStarrySandsFlooring}
	juneFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShellArch, villagerFurnitureShellBed, villagerFurnitureShellSpeaker, villagerFurnitureShellLamp, villagerFurnitureShellFountain, villagerFurnitureShellPartition}}
	juneGames         = villagersGames{[]VillagerGame{}} // TBD
	juneGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	juneInterest      = villagersInterest{villagerInterestNature}
	juneName          = villagersName{villagerNameJune}
	juneNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	juneMusic         = villagersMusic{villagerMusicKKIsland}
	junePersonality   = villagersPersonality{villagerPersonalityNormal}
	juneSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	juneStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	juneWallpaper     = villagersWallpaper{villagerWallpaperOceanHorizonWall}
)
