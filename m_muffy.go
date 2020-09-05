package atsumori

import "time"

// Muffy is an Animal Crossing villager.
var Muffy = villager{
	muffyAstrology,
	muffyBirthDay,
	muffyBirthMonth,
	muffyBubbleColor,
	muffyCategory,
	muffyClothes,
	muffyClothesColors,
	muffyFlooring,
	muffyFurniture,
	muffyGames,
	muffyGender,
	muffyInterest,
	muffyName,
	muffyNameColor,
	muffyMusic,
	muffyPersonality,
	muffySpecies,
	muffyStyle,
	muffyWallpaper}

var (
	muffyAstrology     = villagersAstrology{villagerAstrologyAquarius}
	muffyBirthDay      = villagersBirthDay{14}
	muffyBirthMonth    = villagersBirthMonth{time.February} // 2
	muffyBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	muffyCategory      = villagersCategory{villagerCategoryA}
	muffyClothes       = villagersClothes{} // 3385
	muffyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorPurple}}
	muffyFlooring      = villagersFlooring{villagerFlooringPurpleDesertTileFlooring}
	muffyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSpinningWheel, villagerFurnitureWoodenStool, villagerFurnitureLilyRecordPlayer, villagerFurnitureIronCloset, villagerFurnitureOldSewingMachine, villagerFurnitureFloralSwag, villagerFurnitureDeluxeWasher, villagerFurnitureRattanTowelBasket, villagerFurnitureIroningBoard, villagerFurnitureAntiqueBed, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureIronWorktable, villagerFurnitureIronEntranceMat, villagerFurnitureSturdySewingBox, villagerFurnitureSewingMachine}}
	muffyGames         = villagersGames{[]VillagerGame{}} // TBD
	muffyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	muffyInterest      = villagersInterest{villagerInterestMusic}
	muffyName          = villagersName{villagerNameMuffy}
	muffyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	muffyMusic         = villagersMusic{villagerMusicKKSonata}
	muffyPersonality   = villagersPersonality{villagerPersonalityBigSister}
	muffySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	muffyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	muffyWallpaper     = villagersWallpaper{villagerWallpaperBlackBotanicalTileWall}
)
