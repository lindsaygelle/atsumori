package atsumori

import "time"

// Fauna is an Animal Crossing villager.
var Fauna = villager{
	faunaAstrology,
	faunaBirthDay,
	faunaBirthMonth,
	faunaBubbleColor,
	faunaCategory,
	faunaClothes,
	faunaClothesColors,
	faunaFlooring,
	faunaFurniture,
	faunaGames,
	faunaGender,
	faunaInterest,
	faunaName,
	faunaNameColor,
	faunaMusic,
	faunaPersonality,
	faunaSpecies,
	faunaStyle,
	faunaWallpaper}

var (
	faunaAstrology     = villagersAstrology{villagerAstrologyAries}
	faunaBirthDay      = villagersBirthDay{26}
	faunaBirthMonth    = villagersBirthMonth{time.March} // 3
	faunaBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	faunaCategory      = villagersCategory{villagerCategoryA}
	faunaClothes       = villagersClothes{} // 4409
	faunaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorWhite}}
	faunaFlooring      = villagersFlooring{villagerFlooringMintDotFlooring}
	faunaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBlockBed, villagerFurnitureWoodenBlockBookshelf, villagerFurniturePotRack, villagerFurnitureGasRange, villagerFurnitureWoodenBlockChair, villagerFurnitureWoodenBlockChest, villagerFurniturePortableRecordPlayer, villagerFurnitureWoodenBlockTable, villagerFurnitureMiniFridge, villagerFurnitureCuttingBoard, villagerFurnitureCuckooClock, villagerFurnitureVintageTVTray, villagerFurnitureTreesBountyLamp, villagerFurnitureWoodenBlockstool, villagerFurnitureBookStands}}
	faunaGames         = villagersGames{[]VillagerGame{}} // TBD
	faunaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	faunaInterest      = villagersInterest{villagerInterestNature}
	faunaName          = villagersName{villagerNameFauna}
	faunaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	faunaMusic         = villagersMusic{villagerMusicKKStroll}
	faunaPersonality   = villagersPersonality{villagerPersonalityNormal}
	faunaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	faunaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	faunaWallpaper     = villagersWallpaper{villagerWallpaperRedBrickWall}
)
