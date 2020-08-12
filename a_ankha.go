package atsumori

import "time"

// Ankha is an Animal Crossing villager
var Ankha = villager{
	ankhaAstrology,
	ankhaBirthDay,
	ankhaBirthMonth,
	ankhaBubbleColor,
	ankhaCategory,
	ankhaClothes,
	ankhaClothesColors,
	ankhaFlooring,
	ankhaFurniture,
	ankhaGames,
	ankhaGender,
	ankhaInterest,
	ankhaName,
	ankhaNameColor,
	ankhaMusic,
	ankhaPersonality,
	ankhaSpecies,
	ankhaStyle,
	ankhaWallpaper}

var (
	ankhaAstrology     = villagersAstrology{villagerAstrologyVirgo}
	ankhaBirthDay      = villagersBirthDay{22}
	ankhaBirthMonth    = villagersBirthMonth{time.September}
	ankhaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF80D}
	ankhaCategory      = villagersCategory{villagerCategoryB}
	ankhaClothes       = villagersClothes{} // 3441
	ankhaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorBrown}}
	ankhaFlooring      = villagersFlooring{villagerFlooringPyramidTile}
	ankhaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGoldenToilet, villagerFurniturePyramid, villagerFurnitureWallMountedCandle, villagerFurnitureGoldenCasket, villagerFurnitureLibraScale, villagerFurnitureGoldenCasket, villagerFurnitureGoldenCasket, villagerFurnitureGoldBars, villagerFurnitureGoldenDishes}}
	ankhaGames         = villagersGames{[]VillagerGame{}} // TBD
	ankhaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	ankhaInterest      = villagersInterest{villagerInterestNature}
	ankhaName          = villagersName{villagerNameAnkha}
	ankhaNameColor     = villagersNameColor{villagerNameColor9B8986}
	ankhaMusic         = villagersMusic{} // K.K. Bazaar
	ankhaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	ankhaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	ankhaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleSimple}}
	ankhaWallpaper     = villagersWallpaper{villagerWallpaperAncientWall}
)
