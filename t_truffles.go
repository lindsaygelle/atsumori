package atsumori

import "time"

// Truffles is an Animal Crossing villager.
var Truffles = villager{
	trufflesAstrology,
	trufflesBirthDay,
	trufflesBirthMonth,
	trufflesBubbleColor,
	trufflesCategory,
	trufflesClothes,
	trufflesClothesColors,
	trufflesFlooring,
	trufflesFurniture,
	trufflesGames,
	trufflesGender,
	trufflesInterest,
	trufflesName,
	trufflesNameColor,
	trufflesMusic,
	trufflesPersonality,
	trufflesSpecies,
	trufflesStyle,
	trufflesWallpaper}

var (
	trufflesAstrology     = villagersAstrology{villagerAstrologyLeo}
	trufflesBirthDay      = villagersBirthDay{28}
	trufflesBirthMonth    = villagersBirthMonth{time.July} // 7
	trufflesBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	trufflesCategory      = villagersCategory{villagerCategoryB}
	trufflesClothes       = villagersClothes{villagerClothesTeeParkaCombo} // 4423
	trufflesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorRed}}
	trufflesFlooring      = villagersFlooring{villagerFlooringBrownHoneycombTile}
	trufflesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRefrigerator, villagerFurnitureCuteMusicPlayer, villagerFurnitureWallMountedTV50In, villagerFurnitureFloorLamp, villagerFurnitureWoodenDoubleBed, villagerFurnitureSystemKitchen, villagerFurnitureCuteSofa, villagerFurnitureAntiqueTable, villagerFurnitureSimpleMediumOrangeMat, villagerFurnitureRiceCooker}}
	trufflesGames         = villagersGames{[]VillagerGame{}} // TBD
	trufflesGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	trufflesInterest      = villagersInterest{villagerInterestFashion}
	trufflesName          = villagersName{villagerNameTruffles}
	trufflesNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	trufflesMusic         = villagersMusic{villagerMusicNeapolitan}
	trufflesPersonality   = villagersPersonality{villagerPersonalityPeppy}
	trufflesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	trufflesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleGorgeous}}
	trufflesWallpaper     = villagersWallpaper{villagerWallpaperModWall}
)
