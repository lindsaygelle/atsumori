package atsumori

import "time"

// Freckles is an Animal Crossing villager
var Freckles = villager{
	frecklesAstrology,
	frecklesBirthDay,
	frecklesBirthMonth,
	frecklesBubbleColor,
	frecklesCategory,
	frecklesClothes,
	frecklesClothesColors,
	frecklesFlooring,
	frecklesFurniture,
	frecklesGames,
	frecklesGender,
	frecklesInterest,
	frecklesName,
	frecklesNameColor,
	frecklesMusic,
	frecklesPersonality,
	frecklesSpecies,
	frecklesStyle,
	frecklesWallpaper}

var (
	frecklesAstrology     = villagersAstrology{villagerAstrologyPisces}
	frecklesBirthDay      = villagersBirthDay{19}
	frecklesBirthMonth    = villagersBirthMonth{time.February} // 2
	frecklesBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	frecklesCategory      = villagersCategory{villagerCategoryB}
	frecklesClothes       = villagersClothes{} // 8820
	frecklesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorColorful}}
	frecklesFlooring      = villagersFlooring{villagerFlooringBambooFlooring}
	frecklesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureUtilitySink, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureZenCushion, villagerFurnitureCardboardBox, villagerFurnitureCoolerBox, villagerFurnitureTuna, villagerFurnitureFishingRodStand, villagerFurnitureTeaTable, villagerFurnitureFishingBoatFlag, villagerFurniturePortableRadio, villagerFurnitureTatamiMat, villagerFurnitureFishingBoatFlag, villagerFurnitureOpenFrameKitchen, villagerFurnitureKettle, villagerFurnitureTatamiMat, villagerFurnitureFishPrint, villagerFurnitureFishingBoatFlag, villagerFurnitureFishingBoatFlag, villagerFurnitureMagneticKnifeRack}}
	frecklesGames         = villagersGames{[]VillagerGame{}} // TBD
	frecklesGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	frecklesInterest      = villagersInterest{villagerInterestFashion}
	frecklesName          = villagersName{villagerNameFreckles}
	frecklesNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	frecklesMusic         = villagersMusic{villagerMusicMarineSong2001}
	frecklesPersonality   = villagersPersonality{villagerPersonalityPeppy}
	frecklesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	frecklesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	frecklesWallpaper     = villagersWallpaper{villagerWallpaperLatticeWall}
)
