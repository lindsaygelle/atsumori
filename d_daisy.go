package atsumori

import "time"

// Daisy is an Animal Crossing villager.
var Daisy = villager{
	daisyAstrology,
	daisyBirthDay,
	daisyBirthMonth,
	daisyBubbleColor,
	daisyCategory,
	daisyClothes,
	daisyClothesColors,
	daisyFlooring,
	daisyFurniture,
	daisyGames,
	daisyGender,
	daisyInterest,
	daisyName,
	daisyNameColor,
	daisyMusic,
	daisyPersonality,
	daisySpecies,
	daisyStyle,
	daisyWallpaper}

var (
	daisyAstrology     = villagersAstrology{villagerAstrologyScorpio}
	daisyBirthDay      = villagersBirthDay{16}
	daisyBirthMonth    = villagersBirthMonth{time.November} // 11
	daisyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	daisyCategory      = villagersCategory{villagerCategoryB}
	daisyClothes       = villagersClothes{villagerClothesColorfulStripedSweater} // 4613
	daisyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorBlue}}
	daisyFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	daisyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBookshelf, villagerFurnitureYucca, villagerFurnitureWritingDesk, villagerFurnitureWritingChair, villagerFurnitureWoodenChest, villagerFurnitureCuteMusicPlayer, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureWoodenLowTable, villagerFurnitureTeaSet, villagerFurnitureWoodenDoubleBed, villagerFurnitureLCDTV50In, villagerFurnitureSnowflakeRug}}
	daisyGames         = villagersGames{[]VillagerGame{}} // TBD
	daisyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	daisyInterest      = villagersInterest{villagerInterestEducation}
	daisyName          = villagersName{villagerNameDaisy}
	daisyNameColor     = villagersNameColor{villagerNameColor879B96}
	daisyMusic         = villagersMusic{villagerMusicForestLife}
	daisyPersonality   = villagersPersonality{villagerPersonalityNormal}
	daisySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	daisyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	daisyWallpaper     = villagersWallpaper{villagerWallpaperBlueBlossomingWall}
)
