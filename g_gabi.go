package atsumori

import "time"

// Gabi is an Animal Crossing villager
var Gabi = villager{
	gabiAstrology,
	gabiBirthDay,
	gabiBirthMonth,
	gabiBubbleColor,
	gabiCategory,
	gabiClothes,
	gabiClothesColors,
	gabiFlooring,
	gabiFurniture,
	gabiGames,
	gabiGender,
	gabiInterest,
	gabiName,
	gabiNameColor,
	gabiMusic,
	gabiPersonality,
	gabiSpecies,
	gabiStyle,
	gabiWallpaper}

var (
	gabiAstrology     = villagersAstrology{villagerAstrologySagittarius}
	gabiBirthDay      = villagersBirthDay{16}
	gabiBirthMonth    = villagersBirthMonth{time.December} // 12
	gabiBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	gabiCategory      = villagersCategory{villagerCategoryB}
	gabiClothes       = villagersClothes{villagerClothesGinghamPicnicShirt} // 6811
	gabiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorRed}}
	gabiFlooring      = villagersFlooring{villagerFlooringOrangeRetroFlooring}
	gabiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePotRack, villagerFurnitureShowerSet, villagerFurnitureBathroomTowelRack, villagerFurnitureImperialPartition, villagerFurnitureClawFootTub, villagerFurnitureWoodenTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureAutomaticWasher, villagerFurnitureTanklessToilet, villagerFurnitureRefrigerator, villagerFurniturePlainSink, villagerFurnitureIronwoodKitchenette, villagerFurnitureCuttingBoard}}
	gabiGames         = villagersGames{[]VillagerGame{}} // TBD
	gabiGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	gabiInterest      = villagersInterest{villagerInterestFashion}
	gabiName          = villagersName{villagerNameGabi}
	gabiNameColor     = villagersNameColor{villagerNameColor848484}
	gabiMusic         = villagersMusic{villagerMusicNeapolitan}
	gabiPersonality   = villagersPersonality{villagerPersonalityPeppy}
	gabiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	gabiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleGorgeous}}
	gabiWallpaper     = villagersWallpaper{villagerWallpaperBeadedCurtainWall}
)
