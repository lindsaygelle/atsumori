package atsumori

import "time"

// Eunice is an Animal Crossing villager.
var Eunice = villager{
	euniceAstrology,
	euniceBirthDay,
	euniceBirthMonth,
	euniceBubbleColor,
	euniceCategory,
	euniceClothes,
	euniceClothesColors,
	euniceFlooring,
	euniceFurniture,
	euniceGames,
	euniceGender,
	euniceInterest,
	euniceName,
	euniceNameColor,
	euniceMusic,
	eunicePersonality,
	euniceSpecies,
	euniceStyle,
	euniceWallpaper}

var (
	euniceAstrology     = villagersAstrology{villagerAstrologyAries}
	euniceBirthDay      = villagersBirthDay{3}
	euniceBirthMonth    = villagersBirthMonth{time.April} // 4
	euniceBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	euniceCategory      = villagersCategory{villagerCategoryB}
	euniceClothes       = villagersClothes{villagerClothesAranKnitCardigan} // 3643
	euniceClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorOrange}}
	euniceFlooring      = villagersFlooring{villagerFlooringBlueRubberFlooring}
	euniceFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDeluxeWasher, villagerFurnitureDeluxeWasher, villagerFurnitureDeluxeWasher, villagerFurnitureDeluxeWasher, villagerFurnitureDeluxeWasher, villagerFurnitureDeluxeWasher, villagerFurnitureDeluxeWasher, villagerFurnitureDeluxeWasher, villagerFurnitureFoldingChair, villagerFurnitureDrinkMachine, villagerFurnitureSurveillanceCamera, villagerFurnitureGarbagePail, villagerFurnitureFoldingChair, villagerFurnitureDeluxeWasher, villagerFurnitureDeluxeWasher, villagerFurnitureFoldingChair, villagerFurnitureDeluxeWasher, villagerFurnitureFoldingChair, villagerFurnitureDeluxeWasher, villagerFurnitureIronWorktable, villagerFurnitureRattanTowelBasket, villagerFurnitureWallFan, villagerFurniturePortableRadio, villagerFurnitureMagazine, villagerFurnitureWallClock}}
	euniceGames         = villagersGames{[]VillagerGame{}} // TBD
	euniceGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	euniceInterest      = villagersInterest{villagerInterestFashion}
	euniceName          = villagersName{villagerNameEunice}
	euniceNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	euniceMusic         = villagersMusic{villagerMusicKKDisco}
	eunicePersonality   = villagersPersonality{villagerPersonalityNormal}
	euniceSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	euniceStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	euniceWallpaper     = villagersWallpaper{villagerWallpaperBlueTileWall}
)
