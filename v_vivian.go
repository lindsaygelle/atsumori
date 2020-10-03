package atsumori

import "time"

// Vivian is an Animal Crossing villager.
var Vivian = villager{
	vivianAstrology,
	vivianBirthDay,
	vivianBirthMonth,
	vivianBubbleColor,
	vivianCategory,
	vivianClothes,
	vivianClothesColors,
	vivianFlooring,
	vivianFurniture,
	vivianGames,
	vivianGender,
	vivianInterest,
	vivianName,
	vivianNameColor,
	vivianMusic,
	vivianPersonality,
	vivianSpecies,
	vivianStyle,
	vivianWallpaper}

var (
	vivianAstrology     = villagersAstrology{villagerAstrologyAquarius}
	vivianBirthDay      = villagersBirthDay{26}
	vivianBirthMonth    = villagersBirthMonth{time.January} // 1
	vivianBubbleColor   = villagersBubbleColor{villagerBubbleColorE4B887}
	vivianCategory      = villagersCategory{villagerCategoryA}
	vivianClothes       = villagersClothes{} // 4594
	vivianClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorPurple}}
	vivianFlooring      = villagersFlooring{villagerFlooringDarkWoodPatternFlooring}
	vivianFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShowerBooth, villagerFurnitureIronwoodKitchenette, villagerFurnitureImperialPartition, villagerFurnitureRattanBed, villagerFurnitureRattanArmchair, villagerFurnitureWallMountedTV20In, villagerFurnitureMagneticKnifeRack, villagerFurnitureRattanEndTable, villagerFurniturePortableRecordPlayer, villagerFurnitureRattanLowTable, villagerFurniturePotRack, villagerFurnitureNailArtSet, villagerFurnitureAntiqueConsoleTable, villagerFurnitureAccessoriesStand, villagerFurnitureDeskMirror, villagerFurnitureFloralSwag}}
	vivianGames         = villagersGames{[]VillagerGame{}} // TBD
	vivianGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	vivianInterest      = villagersInterest{villagerInterestEducation}
	vivianName          = villagersName{villagerNameVivian}
	vivianNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	vivianMusic         = villagersMusic{villagerMusicKKCruisin}
	vivianPersonality   = villagersPersonality{villagerPersonalitySnooty}
	vivianSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	vivianStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	vivianWallpaper     = villagersWallpaper{villagerWallpaperWhiteBrickWall}
)
