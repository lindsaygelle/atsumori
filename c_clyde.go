package atsumori

import "time"

// Clyde is an Animal Crossing villager
var Clyde = villager{
	clydeAstrology,
	clydeBirthDay,
	clydeBirthMonth,
	clydeBubbleColor,
	clydeCategory,
	clydeClothes,
	clydeClothesColors,
	clydeFlooring,
	clydeFurniture,
	clydeGames,
	clydeGender,
	clydeInterest,
	clydeName,
	clydeNameColor,
	clydeMusic,
	clydePersonality,
	clydeSpecies,
	clydeStyle,
	clydeWallpaper}

var (
	clydeAstrology     = villagersAstrology{villagerAstrologyTaurus}
	clydeBirthDay      = villagersBirthDay{1}
	clydeBirthMonth    = villagersBirthMonth{time.May} // 5
	clydeBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	clydeCategory      = villagersCategory{villagerCategoryB}
	clydeClothes       = villagersClothes{villagerClothesMadrasPlaidShirt} // 8966
	clydeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorWhite}}
	clydeFlooring      = villagersFlooring{villagerFlooringBirchFlooring}
	clydeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCushion, villagerFurnitureCandyMachine, villagerFurniturePearBed, villagerFurniturePearWardrobe, villagerFurnitureDIYWorkbench, villagerFurnitureSystemKitchen, villagerFurnitureRefrigerator, villagerFurnitureWoodenLowTable, villagerFurnitureOrangeEndTable, villagerFurnitureCherrySpeakers, villagerFurnitureOrangeWallMountedClock}}
	clydeGames         = villagersGames{[]VillagerGame{}} // TBD
	clydeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	clydeInterest      = villagersInterest{villagerInterestPlay}
	clydeName          = villagersName{villagerNameClyde}
	clydeNameColor     = villagersNameColor{villagerNameColor8B5F57}
	clydeMusic         = villagersMusic{villagerMusicNeapolitan}
	clydePersonality   = villagersPersonality{villagerPersonalityLazy}
	clydeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	clydeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	clydeWallpaper     = villagersWallpaper{villagerWallpaperOrangeWall}
)
