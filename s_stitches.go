package atsumori

import "time"

// Stitches is an Animal Crossing villager.
var Stitches = villager{
	stitchesAstrology,
	stitchesBirthDay,
	stitchesBirthMonth,
	stitchesBubbleColor,
	stitchesCategory,
	stitchesClothes,
	stitchesClothesColors,
	stitchesFlooring,
	stitchesFurniture,
	stitchesGames,
	stitchesGender,
	stitchesInterest,
	stitchesName,
	stitchesNameColor,
	stitchesMusic,
	stitchesPersonality,
	stitchesSpecies,
	stitchesStyle,
	stitchesWallpaper}

var (
	stitchesAstrology     = villagersAstrology{villagerAstrologyAquarius}
	stitchesBirthDay      = villagersBirthDay{10}
	stitchesBirthMonth    = villagersBirthMonth{time.February} // 2
	stitchesBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	stitchesCategory      = villagersCategory{villagerCategoryB}
	stitchesClothes       = villagersClothes{villagerClothesStarryTank} // 9535
	stitchesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorWhite}}
	stitchesFlooring      = villagersFlooring{villagerFlooringPatchworkTileFlooring}
	stitchesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureMomsCushion, villagerFurnitureWoodenBlockToy, villagerFurnitureToyBox, villagerFurnitureWoodenSimpleBed, villagerFurnitureBabyChair, villagerFurnitureMomsPlushie, villagerFurnitureWoodenChest, villagerFurnitureMomsLivelyKitchenMat, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenBlockstool, villagerFurnitureMomsPenStand, villagerFurnitureWoodenMiniTable, villagerFurnitureMomsTissueBox, villagerFurnitureWoodenLowTable, villagerFurnitureModelingClay, villagerFurnitureWritingPoster, villagerFurnitureMomsEmbroidery, villagerFurnitureYellowCheckedRug, villagerFurnitureBabyBear}}
	stitchesGames         = villagersGames{[]VillagerGame{}} // TBD
	stitchesGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	stitchesInterest      = villagersInterest{villagerInterestPlay}
	stitchesName          = villagersName{villagerNameStitches}
	stitchesNameColor     = villagersNameColor{villagerNameColor874C25}
	stitchesMusic         = villagersMusic{villagerMusicILoveYou}
	stitchesPersonality   = villagersPersonality{villagerPersonalityLazy}
	stitchesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	stitchesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	stitchesWallpaper     = villagersWallpaper{villagerWallpaperBlueQuiltWall}
)
