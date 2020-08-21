package atsumori

import "time"

// Bea is an Animal Crossing villager
var Bea = villager{
	beaAstrology,
	beaBirthDay,
	beaBirthMonth,
	beaBubbleColor,
	beaCategory,
	beaClothes,
	beaClothesColors,
	beaFlooring,
	beaFurniture,
	beaGames,
	beaGender,
	beaInterest,
	beaName,
	beaNameColor,
	beaMusic,
	beaPersonality,
	beaSpecies,
	beaStyle,
	beaWallpaper}

var (
	beaAstrology     = villagersAstrology{villagerAstrologyLibra}
	beaBirthDay      = villagersBirthDay{15}
	beaBirthMonth    = villagersBirthMonth{time.October} // 10
	beaBubbleColor   = villagersBubbleColor{villagerBubbleColorA87850}
	beaCategory      = villagersCategory{villagerCategoryA}
	beaClothes       = villagersClothes{} // 7784 Dress
	beaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorGreen}}
	beaFlooring      = villagersFlooring{villagerFlooringCorkFlooring}
	beaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanArmchair, villagerFurnitureLectureHallDesk, villagerFurnitureMonstera, villagerFurnitureMenuChalkboard, villagerFurniturePotRack, villagerFurnitureLectureHallDesk, villagerFurnitureLaptop, villagerFurnitureCuteMusicPlayer, villagerFurnitureDinerCounterTable, villagerFurnitureSoupKettle, villagerFurnitureSoupKettle, villagerFurnitureDinerCounterTable, villagerFurnitureSoupKettle, villagerFurnitureSoupKettle, villagerFurnitureOpenFrameKitchen, villagerFurnitureIronEntranceMat, villagerFurnitureDinerCounterTable, villagerFurnitureEspressoMaker, villagerFurnitureInfusedWaterDispenser, villagerFurnitureSimpleKettle}}
	beaGames         = villagersGames{[]VillagerGame{}} // TBD
	beaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	beaInterest      = villagersInterest{villagerInterestNature}
	beaName          = villagersName{villagerNameBea}
	beaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	beaMusic         = villagersMusic{villagerMusicDrivin}
	beaPersonality   = villagersPersonality{villagerPersonalityNormal}
	beaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	beaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	beaWallpaper     = villagersWallpaper{villagerWallpaperBlackboardWall}
)
