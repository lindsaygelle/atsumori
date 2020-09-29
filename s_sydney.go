package atsumori

import "time"

// Sydney is an Animal Crossing villager.
var Sydney = villager{
	sydneyAstrology,
	sydneyBirthDay,
	sydneyBirthMonth,
	sydneyBubbleColor,
	sydneyCategory,
	sydneyClothes,
	sydneyClothesColors,
	sydneyFlooring,
	sydneyFurniture,
	sydneyGames,
	sydneyGender,
	sydneyInterest,
	sydneyName,
	sydneyNameColor,
	sydneyMusic,
	sydneyPersonality,
	sydneySpecies,
	sydneyStyle,
	sydneyWallpaper}

var (
	sydneyAstrology     = villagersAstrology{villagerAstrologyGemini}
	sydneyBirthDay      = villagersBirthDay{21}
	sydneyBirthMonth    = villagersBirthMonth{time.June} // 6
	sydneyBubbleColor   = villagersBubbleColor{villagerBubbleColor7352E8}
	sydneyCategory      = villagersCategory{villagerCategoryB}
	sydneyClothes       = villagersClothes{} // 4165
	sydneyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorYellow}}
	sydneyFlooring      = villagersFlooring{villagerFlooringBluePaintFlooring}
	sydneyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBlockBed, villagerFurnitureWoodenBlockstool, villagerFurnitureAnthuriumPlant, villagerFurniturePotRack, villagerFurnitureGasRange, villagerFurnitureSimpleKettle, villagerFurnitureAirConditioner, villagerFurnitureWoodenBlockStereo, villagerFurnitureWoodenBlockChest, villagerFurnitureWoodenBlockTable, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureHumidifier, villagerFurnitureMiniFridge, villagerFurnitureWoodenBlockChair, villagerFurnitureMicrowave}}
	sydneyGames         = villagersGames{[]VillagerGame{}} // TBD
	sydneyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	sydneyInterest      = villagersInterest{villagerInterestMusic}
	sydneyName          = villagersName{villagerNameSydney}
	sydneyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	sydneyMusic         = villagersMusic{villagerMusicDrivin}
	sydneyPersonality   = villagersPersonality{villagerPersonalityNormal}
	sydneySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	sydneyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	sydneyWallpaper     = villagersWallpaper{villagerWallpaperPurplePuzzleWall}
)
