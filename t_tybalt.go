package atsumori

import "time"

// Tybalt is an Animal Crossing villager.
var Tybalt = villager{
	tybaltAstrology,
	tybaltBirthDay,
	tybaltBirthMonth,
	tybaltBubbleColor,
	tybaltCategory,
	tybaltClothes,
	tybaltClothesColors,
	tybaltFlooring,
	tybaltFurniture,
	tybaltGames,
	tybaltGender,
	tybaltInterest,
	tybaltName,
	tybaltNameColor,
	tybaltMusic,
	tybaltPersonality,
	tybaltSpecies,
	tybaltStyle,
	tybaltWallpaper}

var (
	tybaltAstrology     = villagersAstrology{villagerAstrologyLeo}
	tybaltBirthDay      = villagersBirthDay{19}
	tybaltBirthMonth    = villagersBirthMonth{time.August} // 8
	tybaltBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	tybaltCategory      = villagersCategory{villagerCategoryA}
	tybaltClothes       = villagersClothes{villagerClothesSimpleParka} // 3056
	tybaltClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorPurple}}
	tybaltFlooring      = villagersFlooring{villagerFlooringTigerPrintFlooring}
	tybaltFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurnitureElectricKickScooter, villagerFurnitureDinerCounterChair, villagerFurniturePullUpBarStand, villagerFurnitureDinerCounterTable, villagerFurnitureDIYWorkbench, villagerFurnitureBunkBed, villagerFurnitureToolCart, villagerFurnitureDinerSofa, villagerFurnitureCuteMusicPlayer, villagerFurnitureWallMountedToolBoard, villagerFurnitureThrowbackWallClock}}
	tybaltGames         = villagersGames{[]VillagerGame{}} // TBD
	tybaltGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	tybaltInterest      = villagersInterest{villagerInterestPlay}
	tybaltName          = villagersName{villagerNameTybalt}
	tybaltNameColor     = villagersNameColor{villagerNameColor9B553A}
	tybaltMusic         = villagersMusic{villagerMusicKKGroove}
	tybaltPersonality   = villagersPersonality{villagerPersonalityJock}
	tybaltSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesTiger}}
	tybaltStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	tybaltWallpaper     = villagersWallpaper{villagerWallpaperBluePaintWall}
)
