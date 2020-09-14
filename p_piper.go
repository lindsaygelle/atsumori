package atsumori

import "time"

// Piper is an Animal Crossing villager.
var Piper = villager{
	piperAstrology,
	piperBirthDay,
	piperBirthMonth,
	piperBubbleColor,
	piperCategory,
	piperClothes,
	piperClothesColors,
	piperFlooring,
	piperFurniture,
	piperGames,
	piperGender,
	piperInterest,
	piperName,
	piperNameColor,
	piperMusic,
	piperPersonality,
	piperSpecies,
	piperStyle,
	piperWallpaper}

var (
	piperAstrology     = villagersAstrology{villagerAstrologyAries}
	piperBirthDay      = villagersBirthDay{18}
	piperBirthMonth    = villagersBirthMonth{time.April} // 4
	piperBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	piperCategory      = villagersCategory{villagerCategoryA}
	piperClothes       = villagersClothes{villagerClothesLacyShirt} // 3811
	piperClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorWhite}}
	piperFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	piperFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShowerBooth, villagerFurnitureRattanArmchair, villagerFurnitureAntiqueBed, villagerFurnitureRattanEndTable, villagerFurnitureAntiqueVanity, villagerFurnitureWallMountedTV50In, villagerFurnitureTeaSet, villagerFurnitureAntiqueConsoleTable, villagerFurnitureFragranceDiffuser, villagerFurniturePortableRecordPlayer}}
	piperGames         = villagersGames{[]VillagerGame{}} // TBD
	piperGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	piperInterest      = villagersInterest{villagerInterestPlay}
	piperName          = villagersName{villagerNamePiper}
	piperNameColor     = villagersNameColor{villagerNameColor848484}
	piperMusic         = villagersMusic{villagerMusicKKSynth}
	piperPersonality   = villagersPersonality{villagerPersonalityPeppy}
	piperSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	piperStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	piperWallpaper     = villagersWallpaper{villagerWallpaperCityscapeWall}
)
