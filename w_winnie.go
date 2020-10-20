package atsumori

import "time"

// Winnie is an Animal Crossing villager.
var Winnie = villager{
	winnieAstrology,
	winnieBirthDay,
	winnieBirthMonth,
	winnieBubbleColor,
	winnieCategory,
	winnieClothes,
	winnieClothesColors,
	winnieFlooring,
	winnieFurniture,
	winnieGames,
	winnieGender,
	winnieInterest,
	winnieName,
	winnieNameColor,
	winnieMusic,
	winniePersonality,
	winnieSpecies,
	winnieStyle,
	winnieWallpaper}

var (
	winnieAstrology     = villagersAstrology{villagerAstrologyAquarius}
	winnieBirthDay      = villagersBirthDay{31}
	winnieBirthMonth    = villagersBirthMonth{time.January} // 1
	winnieBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	winnieCategory      = villagersCategory{villagerCategoryB}
	winnieClothes       = villagersClothes{villagerClothesFauxHairSweater} // 7648
	winnieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorGray}}
	winnieFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	winnieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronGardenChair, villagerFurnitureIronGardenChair, villagerFurnitureFirewood, villagerFurnitureCuteMusicPlayer, villagerFurnitureGardenBench, villagerFurnitureHammock, villagerFurnitureIronGardenTable, villagerFurnitureTeaSet, villagerFurnitureGardenFaucet, villagerFurnitureSwingingBench, villagerFurnitureBrickOven}}
	winnieGames         = villagersGames{[]VillagerGame{}} // TBD
	winnieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	winnieInterest      = villagersInterest{villagerInterestFashion}
	winnieName          = villagersName{villagerNameWinnie}
	winnieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	winnieMusic         = villagersMusic{villagerMusicKKCountry}
	winniePersonality   = villagersPersonality{villagerPersonalityPeppy}
	winnieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	winnieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	winnieWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
