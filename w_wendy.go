package atsumori

import "time"

// Wendy is an Animal Crossing villager.
var Wendy = villager{
	wendyAstrology,
	wendyBirthDay,
	wendyBirthMonth,
	wendyBubbleColor,
	wendyCategory,
	wendyClothes,
	wendyClothesColors,
	wendyFlooring,
	wendyFurniture,
	wendyGames,
	wendyGender,
	wendyInterest,
	wendyName,
	wendyNameColor,
	wendyMusic,
	wendyPersonality,
	wendySpecies,
	wendyStyle,
	wendyWallpaper}

var (
	wendyAstrology     = villagersAstrology{villagerAstrologyAries}
	wendyBirthDay      = villagersBirthDay{15}
	wendyBirthMonth    = villagersBirthMonth{time.August} // 8
	wendyBubbleColor   = villagersBubbleColor{villagerBubbleColor55CFFF}
	wendyCategory      = villagersCategory{villagerCategoryB}
	wendyClothes       = villagersClothes{} // 3287
	wendyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorGreen}}
	wendyFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	wendyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRefrigerator, villagerFurnitureCushion, villagerFurnitureMagazine, villagerFurnitureRoundSpaceHeater, villagerFurnitureWoodenEndTable, villagerFurnitureSnowflakeRug, villagerFurnitureRetroRadiator, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureGasRange, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureWoodenDoubleBed, villagerFurnitureSimpleKettle, villagerFurnitureWoodenChest, villagerFurnitureSnowGlobe, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenLowTable, villagerFurnitureTeaSet}}
	wendyGames         = villagersGames{[]VillagerGame{}} // TBD
	wendyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	wendyInterest      = villagersInterest{villagerInterestFashion}
	wendyName          = villagersName{villagerNameWendy}
	wendyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	wendyMusic         = villagersMusic{villagerMusicSpaceKK}
	wendyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	wendySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	wendyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	wendyWallpaper     = villagersWallpaper{villagerWallpaperSnowflakeWall}
)
