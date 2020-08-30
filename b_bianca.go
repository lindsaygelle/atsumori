package atsumori

import "time"

// Bianca is an Animal Crossing villager.
var Bianca = villager{
	biancaAstrology,
	biancaBirthDay,
	biancaBirthMonth,
	biancaBubbleColor,
	biancaCategory,
	biancaClothes,
	biancaClothesColors,
	biancaFlooring,
	biancaFurniture,
	biancaGames,
	biancaGender,
	biancaInterest,
	biancaName,
	biancaNameColor,
	biancaMusic,
	biancaPersonality,
	biancaSpecies,
	biancaStyle,
	biancaWallpaper}

var (
	biancaAstrology     = villagersAstrology{villagerAstrologySagittarius}
	biancaBirthDay      = villagersBirthDay{13}
	biancaBirthMonth    = villagersBirthMonth{time.December} // 12
	biancaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	biancaCategory      = villagersCategory{villagerCategoryA}
	biancaClothes       = villagersClothes{villagerClothesFrontTieButtonDownShirt} // 8506
	biancaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorOrange}}
	biancaFlooring      = villagersFlooring{villagerFlooringLeopardPrintFlooring}
	biancaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRefrigerator, villagerFurnitureRattanBed, villagerFurnitureWoodenLowTable, villagerFurnitureWoodenWardrobe, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureDinerSofa, villagerFurnitureRattanEndTable, villagerFurnitureCuteTeaTable, villagerFurnitureIronwoodKitchenette, villagerFurnitureHumidifier, villagerFurnitureMixer, villagerFurnitureCuttingBoard, villagerFurnitureIronWallRack}}
	biancaGames         = villagersGames{[]VillagerGame{}} // TBD
	biancaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	biancaInterest      = villagersInterest{villagerInterestPlay}
	biancaName          = villagersName{villagerNameBianca}
	biancaNameColor     = villagersNameColor{villagerNameColor848484}
	biancaMusic         = villagersMusic{villagerMusicKKHouse}
	biancaPersonality   = villagersPersonality{villagerPersonalityPeppy}
	biancaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesTiger}}
	biancaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	biancaWallpaper     = villagersWallpaper{villagerWallpaperWhiteBrickWall}
)
