package atsumori

import "time"

// Del is an Animal Crossing villager
var Del = villager{
	delAstrology,
	delBirthDay,
	delBirthMonth,
	delBubbleColor,
	delCategory,
	delClothes,
	delClothesColors,
	delFlooring,
	delFurniture,
	delGames,
	delGender,
	delInterest,
	delName,
	delNameColor,
	delMusic,
	delPersonality,
	delSpecies,
	delStyle,
	delWallpaper}

var (
	delAstrology     = villagersAstrology{villagerAstrologyGemini}
	delBirthDay      = villagersBirthDay{27}
	delBirthMonth    = villagersBirthMonth{time.May} // 5
	delBubbleColor   = villagersBubbleColor{villagerBubbleColor6B75CE}
	delCategory      = villagersCategory{villagerCategoryB}
	delClothes       = villagersClothes{villagerClothesStripedShirt} // 2655
	delClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	delFlooring      = villagersFlooring{villagerFlooringShipDeck}
	delFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronShelf, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureIronWorktable, villagerFurnitureKeyHolder, villagerFurnitureOilBarrel, villagerFurnitureOilBarrel, villagerFurnitureWallMountedToolBoard, villagerFurnitureFoldingChair, villagerFurnitureFoldingChair, villagerFurnitureOutdoorGenerator, villagerFurnitureUtilitySink, villagerFurnitureGears, villagerFurnitureGears, villagerFurnitureIronWorktable, villagerFurnitureToolCart, villagerFurnitureMetalCan, villagerFurnitureBoardGame, villagerFurnitureGears}}
	delGames         = villagersGames{[]VillagerGame{}} // TBD
	delGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	delInterest      = villagersInterest{villagerInterestFitness}
	delName          = villagersName{villagerNameDel}
	delNameColor     = villagersNameColor{villagerNameColor9AE8DF}
	delMusic         = villagersMusic{villagerMusicPondering}
	delPersonality   = villagersPersonality{villagerPersonalityCranky}
	delSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAlligator}}
	delStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	delWallpaper     = villagersWallpaper{villagerWallpaperIndustrialWall}
)
