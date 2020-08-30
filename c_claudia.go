package atsumori

import "time"

// Claudia is an Animal Crossing villager.
var Claudia = villager{
	claudiaAstrology,
	claudiaBirthDay,
	claudiaBirthMonth,
	claudiaBubbleColor,
	claudiaCategory,
	claudiaClothes,
	claudiaClothesColors,
	claudiaFlooring,
	claudiaFurniture,
	claudiaGames,
	claudiaGender,
	claudiaInterest,
	claudiaName,
	claudiaNameColor,
	claudiaMusic,
	claudiaPersonality,
	claudiaSpecies,
	claudiaStyle,
	claudiaWallpaper}

var (
	claudiaAstrology     = villagersAstrology{villagerAstrologyScorpio}
	claudiaBirthDay      = villagersBirthDay{22}
	claudiaBirthMonth    = villagersBirthMonth{time.November} // 11
	claudiaBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	claudiaCategory      = villagersCategory{villagerCategoryA}
	claudiaClothes       = villagersClothes{} // 4737
	claudiaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorWhite}}
	claudiaFlooring      = villagersFlooring{villagerFlooringSimpleRedFlooring}
	claudiaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBed, villagerFurnitureDoubleSofa, villagerFurnitureHiFiStereo, villagerFurnitureOpenFrameKitchen, villagerFurnitureImperialPartition, villagerFurnitureWhirlpoolBath, villagerFurnitureImperialPartition, villagerFurnitureBathroomTowelRack, villagerFurnitureAntiqueConsoleTable, villagerFurnitureInfusedWaterDispenser, villagerFurnitureFruitBasket, villagerFurnitureShowerSet}}
	claudiaGames         = villagersGames{[]VillagerGame{}} // TBD
	claudiaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	claudiaInterest      = villagersInterest{villagerInterestMusic}
	claudiaName          = villagersName{villagerNameClaudia}
	claudiaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	claudiaMusic         = villagersMusic{villagerMusicKKSynth}
	claudiaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	claudiaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesTiger}}
	claudiaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	claudiaWallpaper     = villagersWallpaper{villagerWallpaperCityscapeWall}
)
