package atsumori

import "time"

// Pudge is an Animal Crossing villager.
var Pudge = villager{
	pudgeAstrology,
	pudgeBirthDay,
	pudgeBirthMonth,
	pudgeBubbleColor,
	pudgeCategory,
	pudgeClothes,
	pudgeClothesColors,
	pudgeFlooring,
	pudgeFurniture,
	pudgeGames,
	pudgeGender,
	pudgeInterest,
	pudgeName,
	pudgeNameColor,
	pudgeMusic,
	pudgePersonality,
	pudgeSpecies,
	pudgeStyle,
	pudgeWallpaper}

var (
	pudgeAstrology     = villagersAstrology{villagerAstrologyGemini}
	pudgeBirthDay      = villagersBirthDay{11}
	pudgeBirthMonth    = villagersBirthMonth{time.June} // 6
	pudgeBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	pudgeCategory      = villagersCategory{villagerCategoryB}
	pudgeClothes       = villagersClothes{villagerClothesLetterJacket} // 8306
	pudgeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBlue}}
	pudgeFlooring      = villagersFlooring{villagerFlooringColorfulPuzzleFlooring}
	pudgeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBlockStereo, villagerFurnitureWoodenBlockChest, villagerFurnitureWoodenBlockBed, villagerFurnitureElephantSlide, villagerFurnitureTrainSet, villagerFurnitureBabyChair, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureWoodenBlockToy, villagerFurnitureWoodenBlockstool, villagerFurnitureToyBox, villagerFurnitureWoodenBlockWallClock, villagerFurnitureRockingHorse}}
	pudgeGames         = villagersGames{[]VillagerGame{}} // TBD
	pudgeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	pudgeInterest      = villagersInterest{villagerInterestPlay}
	pudgeName          = villagersName{villagerNamePudge}
	pudgeNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	pudgeMusic         = villagersMusic{villagerMusicKKStroll}
	pudgePersonality   = villagersPersonality{villagerPersonalityLazy}
	pudgeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	pudgeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	pudgeWallpaper     = villagersWallpaper{villagerWallpaperGreenPlayroomWall}
)
