package atsumori

import "time"

// Pekoe is an Animal Crossing villager.
var Pekoe = villager{
	pekoeAstrology,
	pekoeBirthDay,
	pekoeBirthMonth,
	pekoeBubbleColor,
	pekoeCategory,
	pekoeClothes,
	pekoeClothesColors,
	pekoeFlooring,
	pekoeFurniture,
	pekoeGames,
	pekoeGender,
	pekoeInterest,
	pekoeName,
	pekoeNameColor,
	pekoeMusic,
	pekoePersonality,
	pekoeSpecies,
	pekoeStyle,
	pekoeWallpaper}

var (
	pekoeAstrology     = villagersAstrology{villagerAstrologyTaurus}
	pekoeBirthDay      = villagersBirthDay{18}
	pekoeBirthMonth    = villagersBirthMonth{time.May} // 5
	pekoeBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	pekoeCategory      = villagersCategory{villagerCategoryB}
	pekoeClothes       = villagersClothes{} // 5131
	pekoeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBeige}}
	pekoeFlooring      = villagersFlooring{villagerFlooringBambooFlooring}
	pekoeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureImperialPartition, villagerFurnitureImperialPartition, villagerFurnitureImperialBed, villagerFurnitureImperialChest, villagerFurnitureCherryBlossomBranches, villagerFurniturePhonograph, villagerFurnitureImperialDiningLantern, villagerFurnitureImperialLowTable, villagerFurnitureTraditionalTeaSet, villagerFurnitureImperialDecorativeShelves}}
	pekoeGames         = villagersGames{[]VillagerGame{}} // TBD
	pekoeGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pekoeInterest      = villagersInterest{villagerInterestNature}
	pekoeName          = villagersName{villagerNamePekoe}
	pekoeNameColor     = villagersNameColor{villagerNameColor848484}
	pekoeMusic         = villagersMusic{villagerMusicImperialKK}
	pekoePersonality   = villagersPersonality{villagerPersonalityNormal}
	pekoeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	pekoeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCute}}
	pekoeWallpaper     = villagersWallpaper{villagerWallpaperExquisiteWall}
)
