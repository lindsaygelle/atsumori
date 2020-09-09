package atsumori

import "time"

// Pango is an Animal Crossing villager.
var Pango = villager{
	pangoAstrology,
	pangoBirthDay,
	pangoBirthMonth,
	pangoBubbleColor,
	pangoCategory,
	pangoClothes,
	pangoClothesColors,
	pangoFlooring,
	pangoFurniture,
	pangoGames,
	pangoGender,
	pangoInterest,
	pangoName,
	pangoNameColor,
	pangoMusic,
	pangoPersonality,
	pangoSpecies,
	pangoStyle,
	pangoWallpaper}

var (
	pangoAstrology     = villagersAstrology{villagerAstrologyScorpio}
	pangoBirthDay      = villagersBirthDay{9}
	pangoBirthMonth    = villagersBirthMonth{time.November} // 11
	pangoBubbleColor   = villagersBubbleColor{villagerBubbleColor00D1BD}
	pangoCategory      = villagersCategory{villagerCategoryB}
	pangoClothes       = villagersClothes{} // 7920
	pangoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorPurple}}
	pangoFlooring      = villagersFlooring{villagerFlooringGoldIronParquetFlooring}
	pangoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureImperialBed, villagerFurnitureImperialPartition, villagerFurnitureImperialLowTable, villagerFurnitureGoldenToilet, villagerFurnitureRattanLowTable, villagerFurnitureRattanArmchair, villagerFurnitureImperialChest, villagerFurniturePhonograph, villagerFurnitureClawFootTub}}
	pangoGames         = villagersGames{[]VillagerGame{}} // TBD
	pangoGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pangoInterest      = villagersInterest{villagerInterestFashion}
	pangoName          = villagersName{villagerNamePango}
	pangoNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	pangoMusic         = villagersMusic{villagerMusicKKOasis}
	pangoPersonality   = villagersPersonality{villagerPersonalityPeppy}
	pangoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAnteater}}
	pangoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	pangoWallpaper     = villagersWallpaper{villagerWallpaperImperialWall}
)
