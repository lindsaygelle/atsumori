package atsumori

import "time"

// Ribbot is an Animal Crossing villager.
var Ribbot = villager{
	ribbotAstrology,
	ribbotBirthDay,
	ribbotBirthMonth,
	ribbotBubbleColor,
	ribbotCategory,
	ribbotClothes,
	ribbotClothesColors,
	ribbotFlooring,
	ribbotFurniture,
	ribbotGames,
	ribbotGender,
	ribbotInterest,
	ribbotName,
	ribbotNameColor,
	ribbotMusic,
	ribbotPersonality,
	ribbotSpecies,
	ribbotStyle,
	ribbotWallpaper}

var (
	ribbotAstrology     = villagersAstrology{villagerAstrologyAquarius}
	ribbotBirthDay      = villagersBirthDay{13}
	ribbotBirthMonth    = villagersBirthMonth{time.February} // 2
	ribbotBubbleColor   = villagersBubbleColor{villagerBubbleColorBFBFBF}
	ribbotCategory      = villagersCategory{villagerCategoryB}
	ribbotClothes       = villagersClothes{villagerClothesSimpleParka} // 7974
	ribbotClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	ribbotFlooring      = villagersFlooring{villagerFlooringFutureTechFlooring}
	ribbotFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronCloset, villagerFurnitureIronWallLamp, villagerFurnitureServer, villagerFurnitureHighEndStereo, villagerFurnitureServer, villagerFurnitureIronwoodBed, villagerFurnitureIronWorktable, villagerFurnitureLaptop, villagerFurnitureElectronicsKit, villagerFurnitureElectronicsKit, villagerFurnitureIronShelf, villagerFurnitureToolCart, villagerFurnitureDesktopComputer, villagerFurnitureWallMountedToolBoard, villagerFurnitureBreaker}}
	ribbotGames         = villagersGames{[]VillagerGame{}} // TBD
	ribbotGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	ribbotInterest      = villagersInterest{villagerInterestFitness}
	ribbotName          = villagersName{villagerNameRibbot}
	ribbotNameColor     = villagersNameColor{villagerNameColor5E5E5E}
	ribbotMusic         = villagersMusic{villagerMusicKKTechnopop}
	ribbotPersonality   = villagersPersonality{villagerPersonalityJock}
	ribbotSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	ribbotStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	ribbotWallpaper     = villagersWallpaper{villagerWallpaperCircuitBoardWall}
)
