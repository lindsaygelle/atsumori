package atsumori

import "time"

// Bunnie is an Animal Crossing villager
var Bunnie = villager{
	bunnieAstrology,
	bunnieBirthDay,
	bunnieBirthMonth,
	bunnieBubbleColor,
	bunnieCategory,
	bunnieClothes,
	bunnieClothesColors,
	bunnieFlooring,
	bunnieFurniture,
	bunnieGames,
	bunnieGender,
	bunnieInterest,
	bunnieName,
	bunnieNameColor,
	bunnieMusic,
	bunniePersonality,
	bunnieSpecies,
	bunnieStyle,
	bunnieWallpaper}

var (
	bunnieAstrology     = villagersAstrology{villagerAstrologyTaurus}
	bunnieBirthDay      = villagersBirthDay{9}
	bunnieBirthMonth    = villagersBirthMonth{time.May} // 5
	bunnieBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	bunnieCategory      = villagersCategory{villagerCategoryB}
	bunnieClothes       = villagersClothes{} // 9504
	bunnieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorPink}}
	bunnieFlooring      = villagersFlooring{villagerFlooringWhitePaintFlooring}
	bunnieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJuicyAppleTV, villagerFurnitureGasRange, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureClothesCloset, villagerFurnitureAppleRug, villagerFurnitureWoodenChest, villagerFurnitureMomsCushion, villagerFurnitureMagazineRack, villagerFurnitureWoodenSimpleBed, villagerFurnitureMiniFridge, villagerFurnitureCuteMusicPlayer, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureCorkboard, villagerFurnitureDishDryingRack, villagerFurnitureCuteDIYTable, villagerFurnitureMomsPlayfulKitchenMat, villagerFurnitureWoodenLowTable, villagerFurnitureTeaSet}}
	bunnieGames         = villagersGames{[]VillagerGame{}} // TBD
	bunnieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	bunnieInterest      = villagersInterest{villagerInterestFashion}
	bunnieName          = villagersName{villagerNameBunnie}
	bunnieNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	bunnieMusic         = villagersMusic{} // Forest Life
	bunniePersonality   = villagersPersonality{villagerPersonalityPeppy}
	bunnieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	bunnieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	bunnieWallpaper     = villagersWallpaper{villagerWallpaperPurpleDottedWall}
)
