package atsumori

import "time"

// Curlos is an Animal Crossing villager.
var Curlos = villager{
	curlosAstrology,
	curlosBirthDay,
	curlosBirthMonth,
	curlosBubbleColor,
	curlosCategory,
	curlosClothes,
	curlosClothesColors,
	curlosFlooring,
	curlosFurniture,
	curlosGames,
	curlosGender,
	curlosInterest,
	curlosName,
	curlosNameColor,
	curlosMusic,
	curlosPersonality,
	curlosSpecies,
	curlosStyle,
	curlosWallpaper}

var (
	curlosAstrology     = villagersAstrology{villagerAstrologyTaurus}
	curlosBirthDay      = villagersBirthDay{8}
	curlosBirthMonth    = villagersBirthMonth{time.May} // 5
	curlosBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	curlosCategory      = villagersCategory{villagerCategoryB}
	curlosClothes       = villagersClothes{villagerClothesZigzagShirt} // 8204
	curlosClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorGreen}}
	curlosFlooring      = villagersFlooring{villagerFlooringDarkHerringboneFlooring}
	curlosFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJukebox, villagerFurnitureDoubleSofa, villagerFurnitureRedCarpet, villagerFurnitureDrumSet, villagerFurnitureCacaoTree, villagerFurnitureFoosballTable, villagerFurnitureDinerCounterTable, villagerFurnitureEspressoMaker, villagerFurnitureMiniCactusSet, villagerFurnitureDinerCounterChair, villagerFurnitureAcousticGuitar, villagerFurnitureDinerNeonClock, villagerFurnitureTapestry}}
	curlosGames         = villagersGames{[]VillagerGame{}} // TBD
	curlosGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	curlosInterest      = villagersInterest{villagerInterestNature}
	curlosName          = villagersName{villagerNameCurlos}
	curlosNameColor     = villagersNameColor{villagerNameColor9B553A}
	curlosMusic         = villagersMusic{villagerMusicKKSalsa}
	curlosPersonality   = villagersPersonality{villagerPersonalitySmug}
	curlosSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	curlosStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleGorgeous}}
	curlosWallpaper     = villagersWallpaper{villagerWallpaperWildWoodWall}
)
