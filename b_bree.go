package atsumori

import "time"

// Bree is an Animal Crossing villager
var Bree = villager{
	breeAstrology,
	breeBirthDay,
	breeBirthMonth,
	breeBubbleColor,
	breeCategory,
	breeClothes,
	breeClothesColors,
	breeFlooring,
	breeFurniture,
	breeGames,
	breeGender,
	breeInterest,
	breeName,
	breeNameColor,
	breeMusic,
	breePersonality,
	breeSpecies,
	breeStyle,
	breeWallpaper}

var (
	breeAstrology     = villagersAstrology{villagerAstrologyCancer}
	breeBirthDay      = villagersBirthDay{7}
	breeBirthMonth    = villagersBirthMonth{time.July} // 7
	breeBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	breeCategory      = villagersCategory{villagerCategoryB}
	breeClothes       = villagersClothes{} // 3696
	breeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorBlue}}
	breeFlooring      = villagersFlooring{villagerFlooringBlueDesertTileFlooring}
	breeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanVanity, villagerFurnitureRattanWardrobe, villagerFurnitureTanklessToilet, villagerFurnitureHangingTerrarium, villagerFurnitureBidet, villagerFurnitureWhirlpoolBath, villagerFurnitureRattanLowTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureDeluxeWasher, villagerFurnitureRattanTowelBasket}}
	breeGames         = villagersGames{[]VillagerGame{}} // TBD
	breeGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	breeInterest      = villagersInterest{villagerInterestFashion}
	breeName          = villagersName{villagerNameBree}
	breeNameColor     = villagersNameColor{villagerNameColor848484}
	breeMusic         = villagersMusic{villagerMusicKKBossa}
	breePersonality   = villagersPersonality{villagerPersonalitySnooty}
	breeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	breeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	breeWallpaper     = villagersWallpaper{villagerWallpaperWhiteBotanicalTileWall}
)
