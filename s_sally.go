package atsumori

import "time"

// Sally is an Animal Crossing villager.
var Sally = villager{
	sallyAstrology,
	sallyBirthDay,
	sallyBirthMonth,
	sallyBubbleColor,
	sallyCategory,
	sallyClothes,
	sallyClothesColors,
	sallyFlooring,
	sallyFurniture,
	sallyGames,
	sallyGender,
	sallyInterest,
	sallyName,
	sallyNameColor,
	sallyMusic,
	sallyPersonality,
	sallySpecies,
	sallyStyle,
	sallyWallpaper}

var (
	sallyAstrology     = villagersAstrology{villagerAstrologyGemini}
	sallyBirthDay      = villagersBirthDay{19}
	sallyBirthMonth    = villagersBirthMonth{time.June} // 6
	sallyBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	sallyCategory      = villagersCategory{villagerCategoryB}
	sallyClothes       = villagersClothes{villagerClothesSnowySweater} // 3631
	sallyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorBeige}}
	sallyFlooring      = villagersFlooring{villagerFlooringDarkParquetFlooring}
	sallyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenSimpleBed, villagerFurnitureGasRange, villagerFurnitureCushion, villagerFurnitureWoodenLowTable, villagerFurnitureMomsCandleSet, villagerFurnitureYellowArgyleRug, villagerFurnitureWoodenMiniTable, villagerFurnitureMiniFridge, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureTreesBountyLittleTree, villagerFurniturePotRack, villagerFurnitureTreesBountyMobile, villagerFurnitureCuttingBoard, villagerFurnitureIronwoodClock, villagerFurnitureWoodenChest, villagerFurnitureWoodenEndTable, villagerFurniturePortableRecordPlayer, villagerFurnitureAccessoriesStand, villagerFurnitureYellowArgyleRug, villagerFurnitureTreesBountyLamp}}
	sallyGames         = villagersGames{[]VillagerGame{}} // TBD
	sallyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	sallyInterest      = villagersInterest{villagerInterestMusic}
	sallyName          = villagersName{villagerNameSally}
	sallyNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	sallyMusic         = villagersMusic{villagerMusicOnlyMe}
	sallyPersonality   = villagersPersonality{villagerPersonalityNormal}
	sallySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	sallyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	sallyWallpaper     = villagersWallpaper{villagerWallpaperYellowIntricateWall}
)
