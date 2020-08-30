package atsumori

import "time"

// Dizzy is an Animal Crossing villager.
var Dizzy = villager{
	dizzyAstrology,
	dizzyBirthDay,
	dizzyBirthMonth,
	dizzyBubbleColor,
	dizzyCategory,
	dizzyClothes,
	dizzyClothesColors,
	dizzyFlooring,
	dizzyFurniture,
	dizzyGames,
	dizzyGender,
	dizzyInterest,
	dizzyName,
	dizzyNameColor,
	dizzyMusic,
	dizzyPersonality,
	dizzySpecies,
	dizzyStyle,
	dizzyWallpaper}

var (
	dizzyAstrology     = villagersAstrology{villagerAstrologyCancer}
	dizzyBirthDay      = villagersBirthDay{14}
	dizzyBirthMonth    = villagersBirthMonth{time.July} // 7
	dizzyBubbleColor   = villagersBubbleColor{villagerBubbleColor7FA9FF}
	dizzyCategory      = villagersCategory{villagerCategoryB}
	dizzyClothes       = villagersClothes{villagerClothesBoneTee} // 8200
	dizzyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorBlue}}
	dizzyFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	dizzyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureElephantSlide, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureToyBox, villagerFurnitureWoodenBlockstool, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureWritingPoster, villagerFurnitureTrainSet, villagerFurnitureWoodenBlockStereo, villagerFurnitureWoodenBlockBed, villagerFurnitureWoodenBlockChest}}
	dizzyGames         = villagersGames{[]VillagerGame{}} // TBD
	dizzyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	dizzyInterest      = villagersInterest{villagerInterestPlay}
	dizzyName          = villagersName{villagerNameDizzy}
	dizzyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	dizzyMusic         = villagersMusic{villagerMusicTwoDaysAgo}
	dizzyPersonality   = villagersPersonality{villagerPersonalityLazy}
	dizzySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	dizzyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	dizzyWallpaper     = villagersWallpaper{villagerWallpaperYellowPlayroomWall}
)
