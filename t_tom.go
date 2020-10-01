package atsumori

import "time"

// Tom is an Animal Crossing villager.
var Tom = villager{
	tomAstrology,
	tomBirthDay,
	tomBirthMonth,
	tomBubbleColor,
	tomCategory,
	tomClothes,
	tomClothesColors,
	tomFlooring,
	tomFurniture,
	tomGames,
	tomGender,
	tomInterest,
	tomName,
	tomNameColor,
	tomMusic,
	tomPersonality,
	tomSpecies,
	tomStyle,
	tomWallpaper}

var (
	tomAstrology     = villagersAstrology{villagerAstrologySagittarius}
	tomBirthDay      = villagersBirthDay{10}
	tomBirthMonth    = villagersBirthMonth{time.December} // 12
	tomBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	tomCategory      = villagersCategory{villagerCategoryB}
	tomClothes       = villagersClothes{villagerClothesBulldogJacket} // 5980
	tomClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	tomFlooring      = villagersFlooring{villagerFlooringPaintballFlooring}
	tomFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackRaceCarBed, villagerFurnitureDirectorsChair, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureBreaker, villagerFurniturePortableRadio, villagerFurnitureDIYWorkbench, villagerFurnitureHoseReel, villagerFurnitureTinBucket, villagerFurnitureTireStack, villagerFurnitureToolCart, villagerFurnitureWallMountedToolBoard, villagerFurnitureBroomAndDustpan, villagerFurnitureToolbox}}
	tomGames         = villagersGames{[]VillagerGame{}} // TBD
	tomGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	tomInterest      = villagersInterest{villagerInterestEducation}
	tomName          = villagersName{villagerNameTom}
	tomNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	tomMusic         = villagersMusic{villagerMusicKKRock}
	tomPersonality   = villagersPersonality{villagerPersonalityCranky}
	tomSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	tomStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	tomWallpaper     = villagersWallpaper{villagerWallpaperShutterWall}
)
