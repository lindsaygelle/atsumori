package atsumori

import "time"

// Static is an Animal Crossing villager.
var Static = villager{
	staticAstrology,
	staticBirthDay,
	staticBirthMonth,
	staticBubbleColor,
	staticCategory,
	staticClothes,
	staticClothesColors,
	staticFlooring,
	staticFurniture,
	staticGames,
	staticGender,
	staticInterest,
	staticName,
	staticNameColor,
	staticMusic,
	staticPersonality,
	staticSpecies,
	staticStyle,
	staticWallpaper}

var (
	staticAstrology     = villagersAstrology{villagerAstrologyCancer}
	staticBirthDay      = villagersBirthDay{9}
	staticBirthMonth    = villagersBirthMonth{time.July} // 7
	staticBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	staticCategory      = villagersCategory{villagerCategoryB}
	staticClothes       = villagersClothes{villagerClothesDragonJacket} // 3256
	staticClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorYellow}}
	staticFlooring      = villagersFlooring{villagerFlooringSkullPrintFlooring}
	staticFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDIYWorkbench, villagerFurnitureHighEndStereo, villagerFurnitureRattanEndTable, villagerFurnitureElectricGuitar, villagerFurnitureRattanArmchair, villagerFurnitureRattanBed, villagerFurnitureElectricKickScooter, villagerFurnitureRockGuitar, villagerFurnitureWallMountedToolBoard, villagerFurnitureIronwoodClock, villagerFurnitureIronWallLamp, villagerFurnitureToolCart, villagerFurnitureElectronicsKit}}
	staticGames         = villagersGames{[]VillagerGame{}} // TBD
	staticGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	staticInterest      = villagersInterest{villagerInterestMusic}
	staticName          = villagersName{villagerNameStatic}
	staticNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	staticMusic         = villagersMusic{villagerMusicSurfinKK}
	staticPersonality   = villagersPersonality{villagerPersonalityCranky}
	staticSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	staticStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	staticWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
