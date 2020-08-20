package atsumori

import "time"

// Cheri is an Animal Crossing villager
var Cheri = villager{
	cheriAstrology,
	cheriBirthDay,
	cheriBirthMonth,
	cheriBubbleColor,
	cheriCategory,
	cheriClothes,
	cheriClothesColors,
	cheriFlooring,
	cheriFurniture,
	cheriGames,
	cheriGender,
	cheriInterest,
	cheriName,
	cheriNameColor,
	cheriMusic,
	cheriPersonality,
	cheriSpecies,
	cheriStyle,
	cheriWallpaper}

var (
	cheriAstrology     = villagersAstrology{villagerAstrologyPisces}
	cheriBirthDay      = villagersBirthDay{17}
	cheriBirthMonth    = villagersBirthMonth{time.March} // 3
	cheriBubbleColor   = villagersBubbleColor{villagerBubbleColorDB6161}
	cheriCategory      = villagersCategory{villagerCategoryB}
	cheriClothes       = villagersClothes{} // 8818
	cheriClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorLightBlue}}
	cheriFlooring      = villagersFlooring{villagerFlooringSimpleWhiteFlooring}
	cheriFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBookshelf, villagerFurnitureGasRange, villagerFurnitureYucca, villagerFurnitureDoubleSofa, villagerFurnitureClawFootTub, villagerFurniturePotRack, villagerFurnitureShowerSet, villagerFurnitureBathroomTowelRack, villagerFurnitureNaturalWoodenDeckRug, villagerFurnitureCuteBed, villagerFurnitureIronWorktable, villagerFurnitureTanklessToilet, villagerFurnitureCuteTeaTable, villagerFurnitureNaturalWoodenDeckRug, villagerFurnitureAnalogKitchenScale, villagerFurnitureCardboardBox, villagerFurnitureCuteMusicPlayer, villagerFurnitureMiniFridge, villagerFurnitureRevolvingSpiceRack}}
	cheriGames         = villagersGames{[]VillagerGame{}} // TBD
	cheriGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	cheriInterest      = villagersInterest{villagerInterestFashion}
	cheriName          = villagersName{villagerNameCheri}
	cheriNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	cheriMusic         = villagersMusic{} // K.K. Disco
	cheriPersonality   = villagersPersonality{villagerPersonalityPeppy}
	cheriSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	cheriStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	cheriWallpaper     = villagersWallpaper{villagerWallpaperRedDottedWall}
)
