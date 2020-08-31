package atsumori

import "time"

// Maelle is an Animal Crossing villager.
var Maelle = villager{
	maelleAstrology,
	maelleBirthDay,
	maelleBirthMonth,
	maelleBubbleColor,
	maelleCategory,
	maelleClothes,
	maelleClothesColors,
	maelleFlooring,
	maelleFurniture,
	maelleGames,
	maelleGender,
	maelleInterest,
	maelleName,
	maelleNameColor,
	maelleMusic,
	maellePersonality,
	maelleSpecies,
	maelleStyle,
	maelleWallpaper}

var (
	maelleAstrology     = villagersAstrology{villagerAstrologyAries}
	maelleBirthDay      = villagersBirthDay{8}
	maelleBirthMonth    = villagersBirthMonth{time.April} // 4
	maelleBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	maelleCategory      = villagersCategory{villagerCategoryB}
	maelleClothes       = villagersClothes{villagerClothesFlowerSweater} // 3597
	maelleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorWhite}}
	maelleFlooring      = villagersFlooring{villagerFlooringSidewalkFlooring}
	maelleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronGardenChair, villagerFurnitureIronGardenChair, villagerFurnitureLilyRecordPlayer, villagerFurniturePhoneBox, villagerFurnitureCottonCandyStall, villagerFurnitureMenuChalkboard, villagerFurnitureIronGardenTable, villagerFurnitureIronwoodKitchenette, villagerFurnitureCoffeeCup, villagerFurnitureEspressoMaker, villagerFurnitureLectureHallDesk, villagerFurnitureStovetopEspressoMaker, villagerFurniturePopcornMachine}}
	maelleGames         = villagersGames{[]VillagerGame{}} // TBD
	maelleGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	maelleInterest      = villagersInterest{villagerInterestFashion}
	maelleName          = villagersName{villagerNameMaelle}
	maelleNameColor     = villagersNameColor{villagerNameColor9B553A}
	maelleMusic         = villagersMusic{villagerMusicCafeKK}
	maellePersonality   = villagersPersonality{villagerPersonalitySnooty}
	maelleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	maelleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	maelleWallpaper     = villagersWallpaper{villagerWallpaperTreeLinedWall}
)
