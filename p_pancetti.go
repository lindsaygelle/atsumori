package atsumori

import "time"

// Pancetti is an Animal Crossing villager.
var Pancetti = villager{
	pancettiAstrology,
	pancettiBirthDay,
	pancettiBirthMonth,
	pancettiBubbleColor,
	pancettiCategory,
	pancettiClothes,
	pancettiClothesColors,
	pancettiFlooring,
	pancettiFurniture,
	pancettiGames,
	pancettiGender,
	pancettiInterest,
	pancettiName,
	pancettiNameColor,
	pancettiMusic,
	pancettiPersonality,
	pancettiSpecies,
	pancettiStyle,
	pancettiWallpaper}

var (
	pancettiAstrology     = villagersAstrology{villagerAstrologyScorpio}
	pancettiBirthDay      = villagersBirthDay{14}
	pancettiBirthMonth    = villagersBirthMonth{time.November} // 11
	pancettiBubbleColor   = villagersBubbleColor{villagerBubbleColorE2856B}
	pancettiCategory      = villagersCategory{villagerCategoryA}
	pancettiClothes       = villagersClothes{} // 4406
	pancettiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorYellow}}
	pancettiFlooring      = villagersFlooring{villagerFlooringYellowFloralFlooring}
	pancettiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureWoodenChair, villagerFurnitureWoodenChair, villagerFurnitureFanPalm, villagerFurnitureLilyRecordPlayer, villagerFurnitureCoconutWallPlanter, villagerFurnitureWoodenTable, villagerFurniturePopUpToaster, villagerFurnitureTeaSet, villagerFurnitureIronwoodCupboard, villagerFurnitureEspressoMaker, villagerFurnitureAnalogKitchenScale, villagerFurnitureWallMountedTV50In, villagerFurnitureKitchenIsland, villagerFurnitureStandMixer}}
	pancettiGames         = villagersGames{[]VillagerGame{}} // TBD
	pancettiGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pancettiInterest      = villagersInterest{villagerInterestMusic}
	pancettiName          = villagersName{villagerNamePancetti}
	pancettiNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	pancettiMusic         = villagersMusic{villagerMusicKKCruisin}
	pancettiPersonality   = villagersPersonality{villagerPersonalitySnooty}
	pancettiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	pancettiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleGorgeous}}
	pancettiWallpaper     = villagersWallpaper{villagerWallpaperWhiteBotanicalTileWall}
)
