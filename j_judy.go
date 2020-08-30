package atsumori

import "time"

// Judy is an Animal Crossing villager.
var Judy = villager{
	judyAstrology,
	judyBirthDay,
	judyBirthMonth,
	judyBubbleColor,
	judyCategory,
	judyClothes,
	judyClothesColors,
	judyFlooring,
	judyFurniture,
	judyGames,
	judyGender,
	judyInterest,
	judyName,
	judyNameColor,
	judyMusic,
	judyPersonality,
	judySpecies,
	judyStyle,
	judyWallpaper}

var (
	judyAstrology     = villagersAstrology{villagerAstrologyPisces}
	judyBirthDay      = villagersBirthDay{10}
	judyBirthMonth    = villagersBirthMonth{time.March} // 3
	judyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFEBFF}
	judyCategory      = villagersCategory{villagerCategoryA}
	judyClothes       = villagersClothes{} // 8869
	judyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorWhite}}
	judyFlooring      = villagersFlooring{villagerFlooringBlueDotFlooring}
	judyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureWoodenBlockBed, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureWoodenBlockTable, villagerFurnitureStarryGarland, villagerFurnitureWoodenBlockWallClock, villagerFurnitureWoodenBlockBench, villagerFurnitureWoodenBlockstool, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureHumidifier, villagerFurnitureWoodenBlockChair, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureWoodenBlockChest, villagerFurnitureWoodenBlockStereo}}
	judyGames         = villagersGames{[]VillagerGame{}} // TBD
	judyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	judyInterest      = villagersInterest{villagerInterestMusic}
	judyName          = villagersName{villagerNameJudy}
	judyNameColor     = villagersNameColor{villagerNameColor725661}
	judyMusic         = villagersMusic{villagerMusicKKLullaby}
	judyPersonality   = villagersPersonality{villagerPersonalitySnooty}
	judySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	judyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	judyWallpaper     = villagersWallpaper{villagerWallpaperStarryWall}
)
