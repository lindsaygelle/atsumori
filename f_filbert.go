package atsumori

import "time"

// Filbert is an Animal Crossing villager
var Filbert = villager{
	filbertAstrology,
	filbertBirthDay,
	filbertBirthMonth,
	filbertBubbleColor,
	filbertCategory,
	filbertClothes,
	filbertClothesColors,
	filbertFlooring,
	filbertFurniture,
	filbertGames,
	filbertGender,
	filbertInterest,
	filbertName,
	filbertNameColor,
	filbertMusic,
	filbertPersonality,
	filbertSpecies,
	filbertStyle,
	filbertWallpaper}

var (
	filbertAstrology     = villagersAstrology{villagerAstrologyGemini}
	filbertBirthDay      = villagersBirthDay{3}
	filbertBirthMonth    = villagersBirthMonth{time.June} // 6
	filbertBubbleColor   = villagersBubbleColor{villagerBubbleColor55CFFF}
	filbertCategory      = villagersCategory{villagerCategoryA}
	filbertClothes       = villagersClothes{villagerClothesTreeSweater} // 8194
	filbertClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	filbertFlooring      = villagersFlooring{villagerFlooringBlueDotFlooring}
	filbertFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackRocket, villagerFurnitureMiniDIYWorkbench, villagerFurnitureCrescentMoonChair, villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenMiniTable, villagerFurnitureStarryGarland, villagerFurnitureWoodenMiniTable, villagerFurnitureWritingPoster, villagerFurnitureStarClock, villagerFurnitureNovaLight, villagerFurnitureWoodenBlockstool, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenLowTable, villagerFurnitureToyBox, villagerFurnitureStarryGarland, villagerFurnitureModelingClay, villagerFurnitureCuteMusicPlayer, villagerFurnitureStarryGarland}}
	filbertGames         = villagersGames{[]VillagerGame{}} // TBD
	filbertGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	filbertInterest      = villagersInterest{villagerInterestNature}
	filbertName          = villagersName{villagerNameFilbert}
	filbertNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	filbertMusic         = villagersMusic{villagerMusicStaleCupcakes}
	filbertPersonality   = villagersPersonality{villagerPersonalityLazy}
	filbertSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	filbertStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	filbertWallpaper     = villagersWallpaper{villagerWallpaperStarrySkyWall}
)
