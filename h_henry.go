package atsumori

import "time"

// Henry is an Animal Crossing villager
var Henry = villager{
	henryAstrology,
	henryBirthDay,
	henryBirthMonth,
	henryBubbleColor,
	henryCategory,
	henryClothes,
	henryClothesColors,
	henryFlooring,
	henryFurniture,
	henryGames,
	henryGender,
	henryInterest,
	henryName,
	henryNameColor,
	henryMusic,
	henryPersonality,
	henrySpecies,
	henryStyle,
	henryWallpaper}

var (
	henryAstrology     = villagersAstrology{villagerAstrologyVirgo}
	henryBirthDay      = villagersBirthDay{21}
	henryBirthMonth    = villagersBirthMonth{time.September} // 9
	henryBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	henryCategory      = villagersCategory{villagerCategoryA}
	henryClothes       = villagersClothes{villagerClothesDenimJacket} // 3209
	henryClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorBlue}}
	henryFlooring      = villagersFlooring{villagerFlooringModernWoodFlooring}
	henryFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureMonstera, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureMagneticKnifeRack, villagerFurnitureIvorySmallRoundMat, villagerFurnitureIronwoodChair, villagerFurniturePortableRecordPlayer, villagerFurnitureIronwoodTable, villagerFurnitureCoffeeCup, villagerFurnitureCoffeeGrinder, villagerFurnitureGasRange, villagerFurnitureIronwoodCupboard, villagerFurnitureIronwoodBed, villagerFurnitureIronwoodKitchenette, villagerFurniturePotRack, villagerFurnitureSimpleKettle, villagerFurnitureIronWallRack, villagerFurnitureMicrowave, villagerFurnitureEspressoMaker, villagerFurniturePopUpToaster}}
	henryGames         = villagersGames{[]VillagerGame{}} // TBD
	henryGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	henryInterest      = villagersInterest{villagerInterestMusic}
	henryName          = villagersName{villagerNameHenry}
	henryNameColor     = villagersNameColor{villagerNameColor28665A}
	henryMusic         = villagersMusic{villagerMusicKKSoul}
	henryPersonality   = villagersPersonality{villagerPersonalitySmug}
	henrySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	henryStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	henryWallpaper     = villagersWallpaper{villagerWallpaperRusticStoneWall}
)
