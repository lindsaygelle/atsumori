package atsumori

import "time"

// Apollo is an Animal Crossing villager.
var Apollo = villager{
	apolloAstrology,
	apolloBirthDay,
	apolloBirthMonth,
	apolloBubbleColor,
	apolloCategory,
	apolloClothes,
	apolloClothesColors,
	apolloFlooring,
	apolloFurniture,
	apolloGames,
	apolloGender,
	apolloInterest,
	apolloName,
	apolloNameColor,
	apolloMusic,
	apolloPersonality,
	apolloSpecies,
	apolloStyle,
	apolloWallpaper}

var (
	apolloAstrology     = villagersAstrology{villagerAstrologyCancer}
	apolloBirthDay      = villagersBirthDay{4}
	apolloBirthMonth    = villagersBirthMonth{time.July} // 7
	apolloBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	apolloCategory      = villagersCategory{villagerCategoryB}
	apolloClothes       = villagersClothes{villagerClothesFlightJacket} // 8564
	apolloClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorBlack}}
	apolloFlooring      = villagersFlooring{villagerFlooringSlateFlooring}
	apolloFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureDinerSofa, villagerFurnitureDinerMiniTable, villagerFurnitureBilliardTable, villagerFurnitureWallMountedTV50In, villagerFurnitureHighEndStereo, villagerFurnitureFanPalm, villagerFurnitureDinerCounterTable, villagerFurnitureFoldingFloorLamp, villagerFurnitureMagazine}}
	apolloGames         = villagersGames{[]VillagerGame{}} // TBD
	apolloGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	apolloInterest      = villagersInterest{villagerInterestMusic}
	apolloName          = villagersName{villagerNameApollo}
	apolloNameColor     = villagersNameColor{villagerNameColor848484}
	apolloMusic         = villagersMusic{villagerMusicKKRock}
	apolloPersonality   = villagersPersonality{villagerPersonalityCranky}
	apolloSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	apolloStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	apolloWallpaper     = villagersWallpaper{villagerWallpaperShutterWall}
)
