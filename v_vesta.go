package atsumori

import "time"

// Vesta is an Animal Crossing villager.
var Vesta = villager{
	vestaAstrology,
	vestaBirthDay,
	vestaBirthMonth,
	vestaBubbleColor,
	vestaCategory,
	vestaClothes,
	vestaClothesColors,
	vestaFlooring,
	vestaFurniture,
	vestaGames,
	vestaGender,
	vestaInterest,
	vestaName,
	vestaNameColor,
	vestaMusic,
	vestaPersonality,
	vestaSpecies,
	vestaStyle,
	vestaWallpaper}

var (
	vestaAstrology     = villagersAstrology{villagerAstrologyAries}
	vestaBirthDay      = villagersBirthDay{16}
	vestaBirthMonth    = villagersBirthMonth{time.April} // 4
	vestaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	vestaCategory      = villagersCategory{villagerCategoryB}
	vestaClothes       = villagersClothes{villagerClothesHandKnitTank} // 2746
	vestaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorRed}}
	vestaFlooring      = villagersFlooring{villagerFlooringFloralMosaicTileFlooring}
	vestaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureOldFashionedWashtub, villagerFurnitureUtilitySink, villagerFurnitureWallClock, villagerFurnitureDeluxeWasher, villagerFurnitureRattanTowelBasket, villagerFurnitureBathroomTowelRack, villagerFurnitureDinerCounterTable, villagerFurnitureMenuChalkboard, villagerFurnitureAirConditioner, villagerFurnitureRattanEndTable, villagerFurnitureAutomaticWasher, villagerFurniturePortableRecordPlayer, villagerFurnitureIroningBoard, villagerFurniturePantsPress, villagerFurnitureDinerCounterTable, villagerFurnitureSewingProject, villagerFurnitureSewingMachine}}
	vestaGames         = villagersGames{[]VillagerGame{}} // TBD
	vestaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	vestaInterest      = villagersInterest{villagerInterestFashion}
	vestaName          = villagersName{villagerNameVesta}
	vestaNameColor     = villagersNameColor{villagerNameColor848484}
	vestaMusic         = villagersMusic{villagerMusicForestLife}
	vestaPersonality   = villagersPersonality{villagerPersonalityNormal}
	vestaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	vestaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	vestaWallpaper     = villagersWallpaper{villagerWallpaperWhiteSubwayTileWall}
)
