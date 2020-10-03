package atsumori

import "time"

// Walker is an Animal Crossing villager.
var Walker = villager{
	walkerAstrology,
	walkerBirthDay,
	walkerBirthMonth,
	walkerBubbleColor,
	walkerCategory,
	walkerClothes,
	walkerClothesColors,
	walkerFlooring,
	walkerFurniture,
	walkerGames,
	walkerGender,
	walkerInterest,
	walkerName,
	walkerNameColor,
	walkerMusic,
	walkerPersonality,
	walkerSpecies,
	walkerStyle,
	walkerWallpaper}

var (
	walkerAstrology     = villagersAstrology{villagerAstrologyGemini}
	walkerBirthDay      = villagersBirthDay{10}
	walkerBirthMonth    = villagersBirthMonth{time.June} // 6
	walkerBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	walkerCategory      = villagersCategory{villagerCategoryB}
	walkerClothes       = villagersClothes{villagerClothesFiveBallTee} // 3322
	walkerClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorRed}}
	walkerFlooring      = villagersFlooring{villagerFlooringColorfulTileFlooring}
	walkerFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWritingDesk, villagerFurnitureWritingChair, villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenSimpleBed, villagerFurnitureCuteMusicPlayer, villagerFurnitureMomsCushion, villagerFurnitureWoodenEndTable, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureWoodenBlockstool, villagerFurnitureToyBox, villagerFurnitureWoodenChest, villagerFurnitureWritingPoster, villagerFurnitureGlobe}}
	walkerGames         = villagersGames{[]VillagerGame{}} // TBD
	walkerGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	walkerInterest      = villagersInterest{villagerInterestPlay}
	walkerName          = villagersName{villagerNameWalker}
	walkerNameColor     = villagersNameColor{villagerNameColor848484}
	walkerMusic         = villagersMusic{villagerMusicTwoDaysAgo}
	walkerPersonality   = villagersPersonality{villagerPersonalityLazy}
	walkerSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	walkerStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	walkerWallpaper     = villagersWallpaper{villagerWallpaperYellowPlayroomWall}
)
