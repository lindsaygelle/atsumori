package atsumori

import "time"

// Ava is an Animal Crossing villager
var Ava = villager{
	avaAstrology,
	avaBirthDay,
	avaBirthMonth,
	avaBubbleColor,
	avaCategory,
	avaClothes,
	avaClothesColors,
	avaFlooring,
	avaFurniture,
	avaGames,
	avaGender,
	avaInterest,
	avaName,
	avaNameColor,
	avaMusic,
	avaPersonality,
	avaSpecies,
	avaStyle,
	avaWallpaper}

var (
	avaAstrology     = villagersAstrology{villagerAstrologyTaurus}
	avaBirthDay      = villagersBirthDay{28}
	avaBirthMonth    = villagersBirthMonth{time.April} // 4
	avaBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	avaCategory      = villagersCategory{villagerCategoryB}
	avaClothes       = villagersClothes{} // 4731
	avaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorGray}}
	avaFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	avaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenWasteBin, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenDoubleBed, villagerFurnitureWoodenChest, villagerFurnitureSystemKitchen, villagerFurnitureWallMountedTV50In, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureAutomaticWasher, villagerFurnitureDryingRack, villagerFurnitureWoodenLowTable, villagerFurniturePinkRoseRug, villagerFurnitureIroningSet}}
	avaGames         = villagersGames{[]VillagerGame{}} // TBD
	avaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	avaInterest      = villagersInterest{villagerInterestMusic}
	avaName          = villagersName{villagerNameAva}
	avaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	avaMusic         = villagersMusic{villagerMusicKKLoveSong}
	avaPersonality   = villagersPersonality{villagerPersonalityNormal}
	avaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	avaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCute}}
	avaWallpaper     = villagersWallpaper{villagerWallpaperWhiteRoseWall}
)
