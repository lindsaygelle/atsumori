package atsumori

import "time"

// Tammy is an Animal Crossing villager.
var Tammy = villager{
	tammyAstrology,
	tammyBirthDay,
	tammyBirthMonth,
	tammyBubbleColor,
	tammyCategory,
	tammyClothes,
	tammyClothesColors,
	tammyFlooring,
	tammyFurniture,
	tammyGames,
	tammyGender,
	tammyInterest,
	tammyName,
	tammyNameColor,
	tammyMusic,
	tammyPersonality,
	tammySpecies,
	tammyStyle,
	tammyWallpaper}

var (
	tammyAstrology     = villagersAstrology{villagerAstrologyCancer}
	tammyBirthDay      = villagersBirthDay{23}
	tammyBirthMonth    = villagersBirthMonth{time.June} // 6
	tammyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF80D}
	tammyCategory      = villagersCategory{villagerCategoryA}
	tammyClothes       = villagersClothes{villagerClothesHawkJacket} // 4343
	tammyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorPurple}}
	tammyFlooring      = villagersFlooring{villagerFlooringGiraffePrintFlooring}
	tammyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteSofa, villagerFurnitureCuteMusicPlayer, villagerFurnitureCuteBed, villagerFurnitureCuteDIYTable, villagerFurnitureFanPalm, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureMagazineRack, villagerFurnitureWoodenWasteBin, villagerFurnitureWallMountedTV50In, villagerFurnitureRetroRadiator, villagerFurnitureSimpleSmallOrangeMat, villagerFurnitureYellowMessageMat}}
	tammyGames         = villagersGames{[]VillagerGame{}} // TBD
	tammyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tammyInterest      = villagersInterest{villagerInterestPlay}
	tammyName          = villagersName{villagerNameTammy}
	tammyNameColor     = villagersNameColor{villagerNameColor9B8986}
	tammyMusic         = villagersMusic{villagerMusicDJKK}
	tammyPersonality   = villagersPersonality{villagerPersonalityBigSister}
	tammySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	tammyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	tammyWallpaper     = villagersWallpaper{villagerWallpaperYellowIntricateWall}
)
