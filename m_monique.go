package atsumori

import "time"

// Monique is an Animal Crossing villager.
var Monique = villager{
	moniqueAstrology,
	moniqueBirthDay,
	moniqueBirthMonth,
	moniqueBubbleColor,
	moniqueCategory,
	moniqueClothes,
	moniqueClothesColors,
	moniqueFlooring,
	moniqueFurniture,
	moniqueGames,
	moniqueGender,
	moniqueInterest,
	moniqueName,
	moniqueNameColor,
	moniqueMusic,
	moniquePersonality,
	moniqueSpecies,
	moniqueStyle,
	moniqueWallpaper}

var (
	moniqueAstrology     = villagersAstrology{villagerAstrologyLibra}
	moniqueBirthDay      = villagersBirthDay{30}
	moniqueBirthMonth    = villagersBirthMonth{time.September} // 9
	moniqueBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	moniqueCategory      = villagersCategory{villagerCategoryB}
	moniqueClothes       = villagersClothes{} // 7887
	moniqueClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorPink}}
	moniqueFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	moniqueFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanLowTable, villagerFurnitureRattanVanity, villagerFurnitureShowerBooth, villagerFurnitureLilyRecordPlayer, villagerFurnitureRattanWasteBin, villagerFurnitureRattanWardrobe, villagerFurnitureRattanBed, villagerFurnitureRattanStool, villagerFurnitureRattanArmchair, villagerFurnitureDeluxeWasher, villagerFurnitureBathroomTowelRack, villagerFurnitureRattanEndTable, villagerFurnitureRattanTableLamp, villagerFurnitureAirConditioner, villagerFurnitureRattanTowelBasket}}
	moniqueGames         = villagersGames{[]VillagerGame{}} // TBD
	moniqueGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	moniqueInterest      = villagersInterest{villagerInterestFashion}
	moniqueName          = villagersName{villagerNameMonique}
	moniqueNameColor     = villagersNameColor{villagerNameColor879B96}
	moniqueMusic         = villagersMusic{villagerMusicKKCruisin}
	moniquePersonality   = villagersPersonality{villagerPersonalitySnooty}
	moniqueSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	moniqueStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	moniqueWallpaper     = villagersWallpaper{villagerWallpaperBeigeDesertTileWall}
)
