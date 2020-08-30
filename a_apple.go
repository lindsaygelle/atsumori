package atsumori

import "time"

// Apple is an Animal Crossing villager.
var Apple = villager{
	appleAstrology,
	appleBirthDay,
	appleBirthMonth,
	appleBubbleColor,
	appleCategory,
	appleClothes,
	appleClothesColors,
	appleFlooring,
	appleFurniture,
	appleGames,
	appleGender,
	appleInterest,
	appleName,
	appleNameColor,
	appleMusic,
	applePersonality,
	appleSpecies,
	appleStyle,
	appleWallpaper}

var (
	appleAstrology     = villagersAstrology{villagerAstrologyVirgo}
	appleBirthDay      = villagersBirthDay{24}
	appleBirthMonth    = villagersBirthMonth{time.September} // 9
	appleBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	appleCategory      = villagersCategory{villagerCategoryA}
	appleClothes       = villagersClothes{villagerClothesMarbleDotsTee} // 8296
	appleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorRed}}
	appleFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	appleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCushion, villagerFurnitureGasRange, villagerFurnitureWoodenWasteBin, villagerFurnitureWoodenLowTable, villagerFurnitureMiniDIYWorkbench, villagerFurniturePotRack, villagerFurnitureWoodenChest, villagerFurnitureCorkboard, villagerFurnitureAppleRug, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenSimpleBed, villagerFurnitureAppleChair, villagerFurnitureCuteWallMountedClock, villagerFurnitureJuicyAppleTV, villagerFurnitureMiniFridge, villagerFurnitureDishDryingRack}}
	appleGames         = villagersGames{[]VillagerGame{}} // TBD
	appleGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	appleInterest      = villagersInterest{villagerInterestPlay}
	appleName          = villagersName{villagerNameApple}
	appleNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	appleMusic         = villagersMusic{villagerMusicILoveYou}
	applePersonality   = villagersPersonality{villagerPersonalityPeppy}
	appleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHamster}}
	appleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	appleWallpaper     = villagersWallpaper{villagerWallpaperAppleWall}
)
