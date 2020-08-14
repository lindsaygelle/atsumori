package atsumori

import "time"

// Bettina is an Animal Crossing villager
var Bettina = villager{
	bettinaAstrology,
	bettinaBirthDay,
	bettinaBirthMonth,
	bettinaBubbleColor,
	bettinaCategory,
	bettinaClothes,
	bettinaClothesColors,
	bettinaFlooring,
	bettinaFurniture,
	bettinaGames,
	bettinaGender,
	bettinaInterest,
	bettinaName,
	bettinaNameColor,
	bettinaMusic,
	bettinaPersonality,
	bettinaSpecies,
	bettinaStyle,
	bettinaWallpaper}

var (
	bettinaAstrology     = villagersAstrology{villagerAstrologyGemini}
	bettinaBirthDay      = villagersBirthDay{12}
	bettinaBirthMonth    = villagersBirthMonth{time.June} // 6
	bettinaBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	bettinaCategory      = villagersCategory{villagerCategoryB}
	bettinaClothes       = villagersClothes{} // 3177
	bettinaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorRed}}
	bettinaFlooring      = villagersFlooring{villagerFlooringBirchFlooring}
	bettinaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDinerCounterChair, villagerFurnitureDinerCounterChair, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureUtilitySink, villagerFurnitureIronHangerStand, villagerFurnitureStandardUmbrellaStand, villagerFurnitureFreezer, villagerFurniturePotRack, villagerFurnitureAntiqueConsoleTable, villagerFurniturePhonograph, villagerFurnitureKitchenIsland, villagerFurnitureCoffeeCup, villagerFurnitureIronEntranceMat, villagerFurnitureIronwoodCupboard, villagerFurnitureEspressoMaker, villagerFurnitureCoffeeGrinder}}
	bettinaGames         = villagersGames{[]VillagerGame{}} // TBD
	bettinaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	bettinaInterest      = villagersInterest{villagerInterestEducation}
	bettinaName          = villagersName{villagerNameBettina}
	bettinaNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	bettinaMusic         = villagersMusic{} // Animal City
	bettinaPersonality   = villagersPersonality{villagerPersonalityNormal}
	bettinaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	bettinaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	bettinaWallpaper     = villagersWallpaper{villagerWallpaperBlackboardWall}
)
