package atsumori

import "time"

// Bertha is an Animal Crossing villager
var Bertha = villager{
	berthaAstrology,
	berthaBirthDay,
	berthaBirthMonth,
	berthaBubbleColor,
	berthaCategory,
	berthaClothes,
	berthaClothesColors,
	berthaFlooring,
	berthaFurniture,
	berthaGames,
	berthaGender,
	berthaInterest,
	berthaName,
	berthaNameColor,
	berthaMusic,
	berthaPersonality,
	berthaSpecies,
	berthaStyle,
	berthaWallpaper}

var (
	berthaAstrology     = villagersAstrology{villagerAstrologyTaurus}
	berthaBirthDay      = villagersBirthDay{25}
	berthaBirthMonth    = villagersBirthMonth{time.April} // 4
	berthaBubbleColor   = villagersBubbleColor{villagerBubbleColor00D1BD}
	berthaCategory      = villagersCategory{villagerCategoryB}
	berthaClothes       = villagersClothes{} // 4506
	berthaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorWhite}}
	berthaFlooring      = villagersFlooring{villagerFlooringRoseFlooring}
	berthaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureWoodenSimpleBed, villagerFurnitureRefrigerator, villagerFurnitureIronwoodKitchenette, villagerFurnitureIronwoodClock, villagerFurnitureIronwoodChair, villagerFurnitureHangingTerrarium, villagerFurnitureAirConditioner, villagerFurnitureCorkboard, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureDeluxeWasher, villagerFurnitureRattanTowelBasket, villagerFurnitureIronwoodLowTable, villagerFurnitureUnfinishedPuzzle}}
	berthaGames         = villagersGames{[]VillagerGame{}} // TBD
	berthaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	berthaInterest      = villagersInterest{villagerInterestEducation}
	berthaName          = villagersName{villagerNameBertha}
	berthaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	berthaMusic         = villagersMusic{villagerMusicTwoDaysAgo}
	berthaPersonality   = villagersPersonality{villagerPersonalityNormal}
	berthaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHippo}}
	berthaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	berthaWallpaper     = villagersWallpaper{villagerWallpaperWhiteRoseWall}
)
