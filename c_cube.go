package atsumori

import "time"

// Cube is an Animal Crossing villager.
var Cube = villager{
	cubeAstrology,
	cubeBirthDay,
	cubeBirthMonth,
	cubeBubbleColor,
	cubeCategory,
	cubeClothes,
	cubeClothesColors,
	cubeFlooring,
	cubeFurniture,
	cubeGames,
	cubeGender,
	cubeInterest,
	cubeName,
	cubeNameColor,
	cubeMusic,
	cubePersonality,
	cubeSpecies,
	cubeStyle,
	cubeWallpaper}

var (
	cubeAstrology     = villagersAstrology{villagerAstrologyAquarius}
	cubeBirthDay      = villagersBirthDay{29}
	cubeBirthMonth    = villagersBirthMonth{time.January} // 1
	cubeBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	cubeCategory      = villagersCategory{villagerCategoryB}
	cubeClothes       = villagersClothes{villagerClothesSimpleDotsTee} // 8203
	cubeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorColorful}}
	cubeFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	cubeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBlockBed, villagerFurnitureWoodenBlockChair, villagerFurnitureRockingHorse, villagerFurnitureWoodenBlockStereo, villagerFurnitureWoodenBlockTable, villagerFurnitureFrozenTreatSet, villagerFurnitureSnowflakeRug, villagerFurnitureWoodenBlockstool, villagerFurnitureTabletopFestiveTree, villagerFurnitureWoodenBlockChest, villagerFurnitureSnowGlobe, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenBlockBench}}
	cubeGames         = villagersGames{[]VillagerGame{}} // TBD
	cubeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	cubeInterest      = villagersInterest{villagerInterestPlay}
	cubeName          = villagersName{villagerNameCube}
	cubeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	cubeMusic         = villagersMusic{villagerMusicFarewell}
	cubePersonality   = villagersPersonality{villagerPersonalityLazy}
	cubeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	cubeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	cubeWallpaper     = villagersWallpaper{villagerWallpaperSnowflakeWall}
)
