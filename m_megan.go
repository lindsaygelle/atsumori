package atsumori

import "time"

// Megan is an Animal Crossing villager.
var Megan = villager{
	meganAstrology,
	meganBirthDay,
	meganBirthMonth,
	meganBubbleColor,
	meganCategory,
	meganClothes,
	meganClothesColors,
	meganFlooring,
	meganFurniture,
	meganGames,
	meganGender,
	meganInterest,
	meganName,
	meganNameColor,
	meganMusic,
	meganPersonality,
	meganSpecies,
	meganStyle,
	meganWallpaper}

var (
	meganAstrology     = villagersAstrology{villagerAstrologyPisces}
	meganBirthDay      = villagersBirthDay{13}
	meganBirthMonth    = villagersBirthMonth{time.March} // 3
	meganBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	meganCategory      = villagersCategory{villagerCategoryA}
	meganClothes       = villagersClothes{} // 4406
	meganClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorLightBlue}}
	meganFlooring      = villagersFlooring{villagerFlooringDarkHerringboneFlooring}
	meganFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenSimpleBed, villagerFurnitureGasRange, villagerFurnitureBeekeepersHive, villagerFurnitureCacaoTree, villagerFurnitureInflatableSofa, villagerFurnitureMiniDIYWorkbench, villagerFurnitureLongBathtub, villagerFurnitureIronWallRack, villagerFurniturePotRack, villagerFurniturePlainSink, villagerFurnitureWoodenEndTable, villagerFurnitureBathroomTowelRack, villagerFurnitureWoodenTableMirror, villagerFurnitureMiniFridge, villagerFurnitureDishDryingRack}}
	meganGames         = villagersGames{[]VillagerGame{}} // TBD
	meganGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	meganInterest      = villagersInterest{villagerInterestNature}
	meganName          = villagersName{villagerNameMegan}
	meganNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	meganMusic         = villagersMusic{villagerMusicForestLife}
	meganPersonality   = villagersPersonality{villagerPersonalityNormal}
	meganSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	meganStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	meganWallpaper     = villagersWallpaper{villagerWallpaperHoneycombWall}
)
