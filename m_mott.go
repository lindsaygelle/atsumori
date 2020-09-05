package atsumori

import "time"

// Mott is an Animal Crossing villager.
var Mott = villager{
	mottAstrology,
	mottBirthDay,
	mottBirthMonth,
	mottBubbleColor,
	mottCategory,
	mottClothes,
	mottClothesColors,
	mottFlooring,
	mottFurniture,
	mottGames,
	mottGender,
	mottInterest,
	mottName,
	mottNameColor,
	mottMusic,
	mottPersonality,
	mottSpecies,
	mottStyle,
	mottWallpaper}

var (
	mottAstrology     = villagersAstrology{villagerAstrologyCancer}
	mottBirthDay      = villagersBirthDay{10}
	mottBirthMonth    = villagersBirthMonth{time.July} // 7
	mottBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	mottCategory      = villagersCategory{villagerCategoryB}
	mottClothes       = villagersClothes{villagerClothesCollegeCardigan} // 4306
	mottClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorGreen}}
	mottFlooring      = villagersFlooring{villagerFlooringPineBoardFlooring}
	mottFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenChest, villagerFurnitureCuteMusicPlayer, villagerFurnitureWritingPoster, villagerFurnitureWoodenDoubleBed, villagerFurnitureWritingDesk, villagerFurnitureWoodenWasteBin, villagerFurnitureDIYWorkbench, villagerFurnitureWritingChair, villagerFurnitureWoodenEndTable, villagerFurnitureOldFashionedAlarmClock}}
	mottGames         = villagersGames{[]VillagerGame{}} // TBD
	mottGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	mottInterest      = villagersInterest{villagerInterestFitness}
	mottName          = villagersName{villagerNameMott}
	mottNameColor     = villagersNameColor{villagerNameColor8B5F57}
	mottMusic         = villagersMusic{villagerMusicKKFusion}
	mottPersonality   = villagersPersonality{villagerPersonalityJock}
	mottSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesLion}}
	mottStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleElegant}}
	mottWallpaper     = villagersWallpaper{villagerWallpaperGreenMoldedPanelWall}
)
