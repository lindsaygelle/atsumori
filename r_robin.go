package atsumori

import "time"

// Robin is an Animal Crossing villager.
var Robin = villager{
	robinAstrology,
	robinBirthDay,
	robinBirthMonth,
	robinBubbleColor,
	robinCategory,
	robinClothes,
	robinClothesColors,
	robinFlooring,
	robinFurniture,
	robinGames,
	robinGender,
	robinInterest,
	robinName,
	robinNameColor,
	robinMusic,
	robinPersonality,
	robinSpecies,
	robinStyle,
	robinWallpaper}

var (
	robinAstrology     = villagersAstrology{villagerAstrologySagittarius}
	robinBirthDay      = villagersBirthDay{4}
	robinBirthMonth    = villagersBirthMonth{time.December} // 12
	robinBubbleColor   = villagersBubbleColor{villagerBubbleColor6B75CE}
	robinCategory      = villagersCategory{villagerCategoryB}
	robinClothes       = villagersClothes{villagerClothesFrontTieTee} // 4320
	robinClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorPurple}}
	robinFlooring      = villagersFlooring{villagerFlooringBrownArgyleTileFlooring}
	robinFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLongBathtub, villagerFurnitureRattanVanity, villagerFurnitureRattanEndTable, villagerFurnitureRattanLowTable, villagerFurnitureImperialPartition, villagerFurnitureLogGardenLounge, villagerFurnitureRattanTowelBasket, villagerFurniturePortableRecordPlayer, villagerFurnitureAccessoriesStand, villagerFurnitureCypressPlant, villagerFurnitureCypressPlant, villagerFurnitureBirdbath}}
	robinGames         = villagersGames{[]VillagerGame{}} // TBD
	robinGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	robinInterest      = villagersInterest{villagerInterestFashion}
	robinName          = villagersName{villagerNameRobin}
	robinNameColor     = villagersNameColor{villagerNameColor9AE8DF}
	robinMusic         = villagersMusic{villagerMusicKKCruisin}
	robinPersonality   = villagersPersonality{villagerPersonalitySnooty}
	robinSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	robinStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	robinWallpaper     = villagersWallpaper{villagerWallpaperBrownBotanicalTileWall}
)
