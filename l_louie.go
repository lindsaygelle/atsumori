package atsumori

import "time"

// Louie is an Animal Crossing villager.
var Louie = villager{
	louieAstrology,
	louieBirthDay,
	louieBirthMonth,
	louieBubbleColor,
	louieCategory,
	louieClothes,
	louieClothesColors,
	louieFlooring,
	louieFurniture,
	louieGames,
	louieGender,
	louieInterest,
	louieName,
	louieNameColor,
	louieMusic,
	louiePersonality,
	louieSpecies,
	louieStyle,
	louieWallpaper}

var (
	louieAstrology     = villagersAstrology{villagerAstrologyAries}
	louieBirthDay      = villagersBirthDay{26}
	louieBirthMonth    = villagersBirthMonth{time.March} // 3
	louieBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	louieCategory      = villagersCategory{villagerCategoryA}
	louieClothes       = villagersClothes{villagerClothesMuscleTank} // 4267
	louieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorRed}}
	louieFlooring      = villagersFlooring{villagerFlooringColorfulTileFlooring}
	louieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShowerBooth, villagerFurnitureWeightBench, villagerFurnitureToolCart, villagerFurnitureWaterCooler, villagerFurnitureUprightLocker, villagerFurniturePullUpBarStand, villagerFurnitureExerciseBike, villagerFurniturePunchingBag, villagerFurnitureChangingRoom, villagerFurnitureDigitalScale, villagerFurnitureRattanEndTable, villagerFurnitureTapeDeck, villagerFurnitureStadiometer, villagerFurnitureWallMountedTV50In, villagerFurnitureFormalPaper, villagerFurnitureNaturalWoodenDeckRug, villagerFurnitureNaturalWoodenDeckRug, villagerFurnitureDoubleSidedWallClock, villagerFurnitureProteinShakerBottle}}
	louieGames         = villagersGames{[]VillagerGame{}} // TBD
	louieGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	louieInterest      = villagersInterest{villagerInterestFitness}
	louieName          = villagersName{villagerNameLouie}
	louieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	louieMusic         = villagersMusic{villagerMusicKKFusion}
	louiePersonality   = villagersPersonality{villagerPersonalityJock}
	louieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGorilla}}
	louieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	louieWallpaper     = villagersWallpaper{villagerWallpaperWhitePerforatedBoardWall}
)
