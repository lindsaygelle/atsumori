package atsumori

import "time"

// Boone is an Animal Crossing villager
var Boone = villager{
	booneAstrology,
	booneBirthDay,
	booneBirthMonth,
	booneBubbleColor,
	booneCategory,
	booneClothes,
	booneClothesColors,
	booneFlooring,
	booneFurniture,
	booneGames,
	booneGender,
	booneInterest,
	booneName,
	booneNameColor,
	booneMusic,
	boonePersonality,
	booneSpecies,
	booneStyle,
	booneWallpaper}

var (
	booneAstrology     = villagersAstrology{villagerAstrologyVirgo}
	booneBirthDay      = villagersBirthDay{12}
	booneBirthMonth    = villagersBirthMonth{time.September} // 9
	booneBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	booneCategory      = villagersCategory{villagerCategoryB}
	booneClothes       = villagersClothes{villagerClothesCyclingShirt} // 3258
	booneClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorRed}}
	booneFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	booneFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodBurningStove, villagerFurnitureTapestry, villagerFurnitureLogRoundTable, villagerFurnitureMacrameTapestry, villagerFurnitureLogExtraLongSofa, villagerFurnitureLogBed, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureLogDecorativeShelves, villagerFurnitureWallMountedToolBoard, villagerFurnitureOvalEntranceMat, villagerFurnitureLogWallMountedClock}}
	booneGames         = villagersGames{[]VillagerGame{}} // TBD
	booneGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	booneInterest      = villagersInterest{villagerInterestFitness}
	booneName          = villagersName{villagerNameBoone}
	booneNameColor     = villagersNameColor{villagerNameColor874C25}
	booneMusic         = villagersMusic{villagerMusicMountainSong}
	boonePersonality   = villagersPersonality{villagerPersonalityJock}
	booneSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGorilla}}
	booneStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	booneWallpaper     = villagersWallpaper{villagerWallpaperCabinWall}
)
