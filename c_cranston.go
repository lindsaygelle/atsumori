package atsumori

import "time"

// Cranston is an Animal Crossing villager.
var Cranston = villager{
	cranstonAstrology,
	cranstonBirthDay,
	cranstonBirthMonth,
	cranstonBubbleColor,
	cranstonCategory,
	cranstonClothes,
	cranstonClothesColors,
	cranstonFlooring,
	cranstonFurniture,
	cranstonGames,
	cranstonGender,
	cranstonInterest,
	cranstonName,
	cranstonNameColor,
	cranstonMusic,
	cranstonPersonality,
	cranstonSpecies,
	cranstonStyle,
	cranstonWallpaper}

var (
	cranstonAstrology     = villagersAstrology{villagerAstrologyLibra}
	cranstonBirthDay      = villagersBirthDay{23}
	cranstonBirthMonth    = villagersBirthMonth{time.September} // 9
	cranstonBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	cranstonCategory      = villagersCategory{villagerCategoryA}
	cranstonClothes       = villagersClothes{villagerClothesOversizedShawlOvershirt} // 4154
	cranstonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorBrown}}
	cranstonFlooring      = villagersFlooring{villagerFlooringStoneTile}
	cranstonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureImperialPartition, villagerFurnitureImperialBed, villagerFurnitureImperialLowTable, villagerFurnitureAnthuriumPlant, villagerFurniturePaperLantern, villagerFurnitureRedCarpet, villagerFurnitureRedCarpet, villagerFurnitureImperialChest, villagerFurnitureWallMountedTV50In, villagerFurniturePortableRecordPlayer, villagerFurnitureIncenseBurner, villagerFurnitureHangingScroll, villagerFurnitureImperialDecorativeShelves, villagerFurnitureImperialDiningLantern}}
	cranstonGames         = villagersGames{[]VillagerGame{}} // TBD
	cranstonGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	cranstonInterest      = villagersInterest{villagerInterestNature}
	cranstonName          = villagersName{villagerNameCranston}
	cranstonNameColor     = villagersNameColor{villagerNameColor848484}
	cranstonMusic         = villagersMusic{villagerMusicKKFaire}
	cranstonPersonality   = villagersPersonality{villagerPersonalityLazy}
	cranstonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	cranstonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	cranstonWallpaper     = villagersWallpaper{villagerWallpaperWhiteSubwayTileWall}
)
