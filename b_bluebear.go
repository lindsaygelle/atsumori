package atsumori

import "time"

// Bluebear is an Animal Crossing villager
var Bluebear = villager{
	bluebearAstrology,
	bluebearBirthDay,
	bluebearBirthMonth,
	bluebearBubbleColor,
	bluebearCategory,
	bluebearClothes,
	bluebearClothesColors,
	bluebearFlooring,
	bluebearFurniture,
	bluebearGames,
	bluebearGender,
	bluebearInterest,
	bluebearName,
	bluebearNameColor,
	bluebearMusic,
	bluebearPersonality,
	bluebearSpecies,
	bluebearStyle,
	bluebearWallpaper}

var (
	bluebearAstrology     = villagersAstrology{villagerAstrologyCancer}
	bluebearBirthDay      = villagersBirthDay{24}
	bluebearBirthMonth    = villagersBirthMonth{time.June} // 6
	bluebearBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	bluebearCategory      = villagersCategory{villagerCategoryB}
	bluebearClothes       = villagersClothes{villagerClothesAranKnitSweater} // 7672
	bluebearClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorBlue}}
	bluebearFlooring      = villagersFlooring{villagerFlooringBlueDotFlooring}
	bluebearFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteDIYTable, villagerFurnitureCuteSofa, villagerFurnitureCuteWardrobe, villagerFurnitureCuteChair, villagerFurnitureCuteBed, villagerFurnitureCuteTeaTable, villagerFurnitureCuteWallMountedClock, villagerFurnitureCuteMusicPlayer, villagerFurnitureCuteFloorLamp, villagerFurnitureCuteVanity}}
	bluebearGames         = villagersGames{[]VillagerGame{}} // TBD
	bluebearGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	bluebearInterest      = villagersInterest{villagerInterestFashion}
	bluebearName          = villagersName{villagerNameBluebear}
	bluebearNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bluebearMusic         = villagersMusic{} // Only Me
	bluebearPersonality   = villagersPersonality{villagerPersonalityPeppy}
	bluebearSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	bluebearStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	bluebearWallpaper     = villagersWallpaper{villagerWallpaperBlueSimpleClothWall}
)
