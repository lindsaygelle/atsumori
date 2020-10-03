package atsumori

import "time"

// Vladimir is an Animal Crossing villager.
var Vladimir = villager{
	vladimirAstrology,
	vladimirBirthDay,
	vladimirBirthMonth,
	vladimirBubbleColor,
	vladimirCategory,
	vladimirClothes,
	vladimirClothesColors,
	vladimirFlooring,
	vladimirFurniture,
	vladimirGames,
	vladimirGender,
	vladimirInterest,
	vladimirName,
	vladimirNameColor,
	vladimirMusic,
	vladimirPersonality,
	vladimirSpecies,
	vladimirStyle,
	vladimirWallpaper}

var (
	vladimirAstrology     = villagersAstrology{villagerAstrologyLeo}
	vladimirBirthDay      = villagersBirthDay{2}
	vladimirBirthMonth    = villagersBirthMonth{time.August} // 8
	vladimirBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	vladimirCategory      = villagersCategory{villagerCategoryB}
	vladimirClothes       = villagersClothes{villagerClothesStripedShirt} // 4148
	vladimirClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorColorful}}
	vladimirFlooring      = villagersFlooring{villagerFlooringRamshackleFlooring}
	vladimirFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureTapeDeck, villagerFurnitureCampingCot, villagerFurnitureGarbagePail, villagerFurnitureManholeCover, villagerFurnitureIronFrame, villagerFurnitureCone, villagerFurnitureIronFrame, villagerFurnitureDrinkMachine, villagerFurnitureWallMountedToolBoard, villagerFurnitureCone}}
	vladimirGames         = villagersGames{[]VillagerGame{}} // TBD
	vladimirGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	vladimirInterest      = villagersInterest{villagerInterestPlay}
	vladimirName          = villagersName{villagerNameVladimir}
	vladimirNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	vladimirMusic         = villagersMusic{villagerMusicAgentKK}
	vladimirPersonality   = villagersPersonality{villagerPersonalityCranky}
	vladimirSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	vladimirStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	vladimirWallpaper     = villagersWallpaper{villagerWallpaperChainLinkFence}
)
