package atsumori

import "time"

// Savannah is an Animal Crossing villager.
var Savannah = villager{
	savannahAstrology,
	savannahBirthDay,
	savannahBirthMonth,
	savannahBubbleColor,
	savannahCategory,
	savannahClothes,
	savannahClothesColors,
	savannahFlooring,
	savannahFurniture,
	savannahGames,
	savannahGender,
	savannahInterest,
	savannahName,
	savannahNameColor,
	savannahMusic,
	savannahPersonality,
	savannahSpecies,
	savannahStyle,
	savannahWallpaper}

var (
	savannahAstrology     = villagersAstrology{villagerAstrologyAquarius}
	savannahBirthDay      = villagersBirthDay{25}
	savannahBirthMonth    = villagersBirthMonth{time.January} // 1
	savannahBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	savannahCategory      = villagersCategory{villagerCategoryB}
	savannahClothes       = villagersClothes{villagerClothesTopCoat} // 4574
	savannahClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorBlue}}
	savannahFlooring      = villagersFlooring{villagerFlooringZebraPrintFlooring}
	savannahFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureDoubleSofa, villagerFurniturePianoBench, villagerFurnitureUprightPiano, villagerFurnitureRattanArmchair, villagerFurnitureRattanVanity, villagerFurnitureRattanWardrobe, villagerFurnitureRattanEndTable, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureMacrameTapestry, villagerFurnitureRattanLowTable, villagerFurnitureAnthuriumPlant}}
	savannahGames         = villagersGames{[]VillagerGame{}} // TBD
	savannahGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	savannahInterest      = villagersInterest{villagerInterestMusic}
	savannahName          = villagersName{villagerNameSavannah}
	savannahNameColor     = villagersNameColor{villagerNameColor848484}
	savannahMusic         = villagersMusic{villagerMusicForestLife}
	savannahPersonality   = villagersPersonality{villagerPersonalityNormal}
	savannahSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	savannahStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	savannahWallpaper     = villagersWallpaper{villagerWallpaperBlueSubwayTileWall}
)
