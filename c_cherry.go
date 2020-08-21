package atsumori

import "time"

// Cherry is an Animal Crossing villager
var Cherry = villager{
	cherryAstrology,
	cherryBirthDay,
	cherryBirthMonth,
	cherryBubbleColor,
	cherryCategory,
	cherryClothes,
	cherryClothesColors,
	cherryFlooring,
	cherryFurniture,
	cherryGames,
	cherryGender,
	cherryInterest,
	cherryName,
	cherryNameColor,
	cherryMusic,
	cherryPersonality,
	cherrySpecies,
	cherryStyle,
	cherryWallpaper}

var (
	cherryAstrology     = villagersAstrology{villagerAstrologyTaurus}
	cherryBirthDay      = villagersBirthDay{11}
	cherryBirthMonth    = villagersBirthMonth{time.May} // 5
	cherryBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	cherryCategory      = villagersCategory{villagerCategoryA}
	cherryClothes       = villagersClothes{villagerClothesSpiderWebTee} // 8207
	cherryClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorPurple}}
	cherryFlooring      = villagersFlooring{villagerFlooringSkullPrintFlooring}
	cherryFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRoseBed, villagerFurnitureWhirlpoolBath, villagerFurnitureBeachTowel, villagerFurnitureOpenFrameKitchen, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureAntiqueMiniTable, villagerFurnitureFragranceSticks, villagerFurnitureBathroomTowelRack, villagerFurnitureMumCushion, villagerFurnitureShowerSet, villagerFurnitureMumCushion, villagerFurnitureBlackWoodenDeckRug}}
	cherryGames         = villagersGames{[]VillagerGame{}} // TBD
	cherryGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	cherryInterest      = villagersInterest{villagerInterestMusic}
	cherryName          = villagersName{villagerNameCherry}
	cherryNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	cherryMusic         = villagersMusic{villagerMusicKKDB}
	cherryPersonality   = villagersPersonality{villagerPersonalityBigSister}
	cherrySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	cherryStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleElegant}}
	cherryWallpaper     = villagersWallpaper{villagerWallpaperCityscapeWall}
)
