package atsumori

import "time"

// Flurry is an Animal Crossing villager.
var Flurry = villager{
	flurryAstrology,
	flurryBirthDay,
	flurryBirthMonth,
	flurryBubbleColor,
	flurryCategory,
	flurryClothes,
	flurryClothesColors,
	flurryFlooring,
	flurryFurniture,
	flurryGames,
	flurryGender,
	flurryInterest,
	flurryName,
	flurryNameColor,
	flurryMusic,
	flurryPersonality,
	flurrySpecies,
	flurryStyle,
	flurryWallpaper}

var (
	flurryAstrology     = villagersAstrology{villagerAstrologyAquarius}
	flurryBirthDay      = villagersBirthDay{30}
	flurryBirthMonth    = villagersBirthMonth{time.January} // 1
	flurryBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	flurryCategory      = villagersCategory{villagerCategoryA}
	flurryClothes       = villagersClothes{villagerClothesFlowerSweater} // 3597
	flurryClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorPink}}
	flurryFlooring      = villagersFlooring{villagerFlooringArabesqueFlooring}
	flurryFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePianoBench, villagerFurnitureFloorLamp, villagerFurnitureWoodenChest, villagerFurnitureMiniFridge, villagerFurniturePotRack, villagerFurnitureAnthuriumPlant, villagerFurnitureDinnerware, villagerFurnitureGasRange, villagerFurnitureWoodenLowTable, villagerFurnitureMiniDIYWorkbench, villagerFurnitureVentilationFan, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenSimpleBed, villagerFurnitureUprightPiano, villagerFurnitureHamsterCage}}
	flurryGames         = villagersGames{[]VillagerGame{}} // TBD
	flurryGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	flurryInterest      = villagersInterest{villagerInterestNature}
	flurryName          = villagersName{villagerNameFlurry}
	flurryNameColor     = villagersNameColor{villagerNameColor848484}
	flurryMusic         = villagersMusic{villagerMusicHypnoKK}
	flurryPersonality   = villagersPersonality{villagerPersonalityNormal}
	flurrySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHamster}}
	flurryStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCute}}
	flurryWallpaper     = villagersWallpaper{villagerWallpaperBeigeBlossomingWall}
)
