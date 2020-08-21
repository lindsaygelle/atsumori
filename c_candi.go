package atsumori

import "time"

// Candi is an Animal Crossing villager
var Candi = villager{
	candiAstrology,
	candiBirthDay,
	candiBirthMonth,
	candiBubbleColor,
	candiCategory,
	candiClothes,
	candiClothesColors,
	candiFlooring,
	candiFurniture,
	candiGames,
	candiGender,
	candiInterest,
	candiName,
	candiNameColor,
	candiMusic,
	candiPersonality,
	candiSpecies,
	candiStyle,
	candiWallpaper}

var (
	candiAstrology     = villagersAstrology{villagerAstrologyAries}
	candiBirthDay      = villagersBirthDay{13}
	candiBirthMonth    = villagersBirthMonth{time.April} // 4
	candiBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	candiCategory      = villagersCategory{villagerCategoryA}
	candiClothes       = villagersClothes{} // 8593
	candiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorYellow}}
	candiFlooring      = villagersFlooring{villagerFlooringMintDotFlooring}
	candiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenChest, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenDoubleBed, villagerFurnitureCuteSofa, villagerFurnitureCuteDIYTable, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureMumCushion, villagerFurnitureCuteMusicPlayer, villagerFurnitureFloralSwag, villagerFurnitureWoodenWasteBin, villagerFurnitureWoodenLowTable}}
	candiGames         = villagersGames{[]VillagerGame{}} // TBD
	candiGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	candiInterest      = villagersInterest{villagerInterestPlay}
	candiName          = villagersName{villagerNameCandi}
	candiNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	candiMusic         = villagersMusic{villagerMusicNeapolitan}
	candiPersonality   = villagersPersonality{villagerPersonalityPeppy}
	candiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	candiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	candiWallpaper     = villagersWallpaper{villagerWallpaperWhitePaintedWoodWall}
)
