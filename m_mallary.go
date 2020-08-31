package atsumori

import "time"

// Mallary is an Animal Crossing villager.
var Mallary = villager{
	mallaryAstrology,
	mallaryBirthDay,
	mallaryBirthMonth,
	mallaryBubbleColor,
	mallaryCategory,
	mallaryClothes,
	mallaryClothesColors,
	mallaryFlooring,
	mallaryFurniture,
	mallaryGames,
	mallaryGender,
	mallaryInterest,
	mallaryName,
	mallaryNameColor,
	mallaryMusic,
	mallaryPersonality,
	mallarySpecies,
	mallaryStyle,
	mallaryWallpaper}

var (
	mallaryAstrology     = villagersAstrology{villagerAstrologyCancer}
	mallaryBirthDay      = villagersBirthDay{17}
	mallaryBirthMonth    = villagersBirthMonth{time.November} // 11
	mallaryBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	mallaryCategory      = villagersCategory{villagerCategoryB}
	mallaryClothes       = villagersClothes{villagerClothesStripedShirt} // 2655
	mallaryClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	mallaryFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	mallaryFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGardenFaucet, villagerFurnitureHoseReel, villagerFurnitureCuteMusicPlayer, villagerFurnitureFirePit, villagerFurnitureHangingTerrarium, villagerFurnitureGardenBench, villagerFurnitureCypressPlant, villagerFurnitureGardenLantern, villagerFurnitureGardenGnome, villagerFurnitureGardenWagon}}
	mallaryGames         = villagersGames{[]VillagerGame{}} // TBD
	mallaryGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	mallaryInterest      = villagersInterest{villagerInterestFashion}
	mallaryName          = villagersName{villagerNameMallary}
	mallaryNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	mallaryMusic         = villagersMusic{villagerMusicKKBossa}
	mallaryPersonality   = villagersPersonality{villagerPersonalitySnooty}
	mallarySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	mallaryStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	mallaryWallpaper     = villagersWallpaper{villagerWallpaperMossyGardenWall}
)
