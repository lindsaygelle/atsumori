package atsumori

import "time"

// Twiggy is an Animal Crossing villager.
var Twiggy = villager{
	twiggyAstrology,
	twiggyBirthDay,
	twiggyBirthMonth,
	twiggyBubbleColor,
	twiggyCategory,
	twiggyClothes,
	twiggyClothesColors,
	twiggyFlooring,
	twiggyFurniture,
	twiggyGames,
	twiggyGender,
	twiggyInterest,
	twiggyName,
	twiggyNameColor,
	twiggyMusic,
	twiggyPersonality,
	twiggySpecies,
	twiggyStyle,
	twiggyWallpaper}

var (
	twiggyAstrology     = villagersAstrology{villagerAstrologyCancer}
	twiggyBirthDay      = villagersBirthDay{13}
	twiggyBirthMonth    = villagersBirthMonth{time.July} // 7
	twiggyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF80D}
	twiggyCategory      = villagersCategory{villagerCategoryB}
	twiggyClothes       = villagersClothes{villagerClothesStripedTee} // 7768
	twiggyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorBlue}}
	twiggyFlooring      = villagersFlooring{villagerFlooringGrayArgyleTileFlooring}
	twiggyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLogGardenLounge, villagerFurnitureLongBathtub, villagerFurnitureRattanVanity, villagerFurnitureRattanEndTable, villagerFurnitureRattanLowTable, villagerFurnitureImperialPartition, villagerFurnitureRattanTowelBasket, villagerFurniturePortableRecordPlayer, villagerFurnitureAccessoriesStand, villagerFurnitureBirdbath, villagerFurnitureCypressPlant, villagerFurnitureCypressPlant}}
	twiggyGames         = villagersGames{[]VillagerGame{}} // TBD
	twiggyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	twiggyInterest      = villagersInterest{villagerInterestFashion}
	twiggyName          = villagersName{villagerNameTwiggy}
	twiggyNameColor     = villagersNameColor{villagerNameColor9B8986}
	twiggyMusic         = villagersMusic{villagerMusicBubblegumKK}
	twiggyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	twiggySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	twiggyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	twiggyWallpaper     = villagersWallpaper{villagerWallpaperWhiteBotanicalTileWall}
)
