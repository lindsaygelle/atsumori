package atsumori

import "time"

// Angus is an Animal Crossing villager
var Angus = villager{
	angusAstrology,
	angusBirthDay,
	angusBirthMonth,
	angusBubbleColor,
	angusCategory,
	angusClothes,
	angusClothesColors,
	angusFlooring,
	angusFurniture,
	angusGames,
	angusGender,
	angusInterest,
	angusName,
	angusNameColor,
	angusMusic,
	angusPersonality,
	angusSpecies,
	angusStyle,
	angusWallpaper}

var (
	angusAstrology     = villagersAstrology{villagerAstrologyTaurus}
	angusBirthDay      = villagersBirthDay{30}
	angusBirthMonth    = villagersBirthMonth{time.April} // 4
	angusBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	angusCategory      = villagersCategory{villagerCategoryB}
	angusClothes       = villagersClothes{villagerClothesFlameTee} // 9536
	angusClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlack}}
	angusFlooring      = villagersFlooring{villagerFlooringRedAndBlackVinylFlooring}
	angusFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDartboard, villagerFurnitureAntiqueBed, villagerFurnitureYucca, villagerFurnitureAntiqueBureau, villagerFurnitureBilliardTable, villagerFurnitureIronwoodLowTable, villagerFurniturePennant, villagerFurnitureHighEndStereo, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueMiniTable, villagerFurnitureBingoWheel}}
	angusGames         = villagersGames{[]VillagerGame{}} // TBD
	angusGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	angusInterest      = villagersInterest{villagerInterestFitness}
	angusName          = villagersName{villagerNameAngus}
	angusNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	angusMusic         = villagersMusic{} // K.K. Calypso
	angusPersonality   = villagersPersonality{villagerPersonalityCranky}
	angusSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBull}}
	angusStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	angusWallpaper     = villagersWallpaper{villagerWallpaperWildWoodWall}
)
