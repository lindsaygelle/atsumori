package atsumori

import "time"

// Peggy is an Animal Crossing villager.
var Peggy = villager{
	peggyAstrology,
	peggyBirthDay,
	peggyBirthMonth,
	peggyBubbleColor,
	peggyCategory,
	peggyClothes,
	peggyClothesColors,
	peggyFlooring,
	peggyFurniture,
	peggyGames,
	peggyGender,
	peggyInterest,
	peggyName,
	peggyNameColor,
	peggyMusic,
	peggyPersonality,
	peggySpecies,
	peggyStyle,
	peggyWallpaper}

var (
	peggyAstrology     = villagersAstrology{villagerAstrologyGemini}
	peggyBirthDay      = villagersBirthDay{23}
	peggyBirthMonth    = villagersBirthMonth{time.May} // 5
	peggyBubbleColor   = villagersBubbleColor{villagerBubbleColorE4B887}
	peggyCategory      = villagersCategory{villagerCategoryB}
	peggyClothes       = villagersClothes{villagerClothesAranKnitSweater} // 3633
	peggyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorRed}}
	peggyFlooring      = villagersFlooring{villagerFlooringBirchFlooring}
	peggyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteSofa, villagerFurnitureFanPalm, villagerFurnitureShowerSet, villagerFurnitureBathroomTowelRack, villagerFurnitureJuicyAppleTV, villagerFurnitureCuteChair, villagerFurnitureIronShelf, villagerFurnitureDeluxeWasher, villagerFurnitureFloralSwag, villagerFurnitureRattanTowelBasket, villagerFurnitureClawFootTub, villagerFurnitureIroningBoard, villagerFurnitureLoftBedWithDesk, villagerFurnitureHomeworkSet}}
	peggyGames         = villagersGames{[]VillagerGame{}} // TBD
	peggyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	peggyInterest      = villagersInterest{villagerInterestFashion}
	peggyName          = villagersName{villagerNamePeggy}
	peggyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	peggyMusic         = villagersMusic{villagerMusicKKBossa}
	peggyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	peggySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	peggyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleActive}}
	peggyWallpaper     = villagersWallpaper{villagerWallpaperRedIntricateWall}
)
