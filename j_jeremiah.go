package atsumori

import "time"

// Jeremiah is an Animal Crossing villager
var Jeremiah = villager{
	jeremiahAstrology,
	jeremiahBirthDay,
	jeremiahBirthMonth,
	jeremiahBubbleColor,
	jeremiahCategory,
	jeremiahClothes,
	jeremiahClothesColors,
	jeremiahFlooring,
	jeremiahFurniture,
	jeremiahGames,
	jeremiahGender,
	jeremiahInterest,
	jeremiahName,
	jeremiahNameColor,
	jeremiahMusic,
	jeremiahPersonality,
	jeremiahSpecies,
	jeremiahStyle,
	jeremiahWallpaper}

var (
	jeremiahAstrology     = villagersAstrology{villagerAstrologyCancer}
	jeremiahBirthDay      = villagersBirthDay{8}
	jeremiahBirthMonth    = villagersBirthMonth{time.July} // 7
	jeremiahBubbleColor   = villagersBubbleColor{villagerBubbleColor55CFFF}
	jeremiahCategory      = villagersCategory{villagerCategoryB}
	jeremiahClothes       = villagersClothes{villagerClothesGinghamPicnicShirt} // 8974
	jeremiahClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorYellow}}
	jeremiahFlooring      = villagersFlooring{villagerFlooringBlueDotFlooring}
	jeremiahFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCushion, villagerFurnitureWoodenBlockstool, villagerFurnitureToyBox, villagerFurnitureElephantSlide, villagerFurnitureWoodenChest, villagerFurnitureCuteMusicPlayer, villagerFurnitureCuteDIYTable, villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenBlockWallClock, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenLowTable, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureUnfinishedPuzzle}}
	jeremiahGames         = villagersGames{[]VillagerGame{}} // TBD
	jeremiahGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	jeremiahInterest      = villagersInterest{villagerInterestPlay}
	jeremiahName          = villagersName{villagerNameJeremiah}
	jeremiahNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	jeremiahMusic         = villagersMusic{villagerMusicHypnoKK}
	jeremiahPersonality   = villagersPersonality{villagerPersonalityLazy}
	jeremiahSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	jeremiahStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	jeremiahWallpaper     = villagersWallpaper{villagerWallpaperBluePlayroomWall}
)
