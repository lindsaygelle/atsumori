package atsumori

import "time"

// Flo is an Animal Crossing villager.
var Flo = villager{
	floAstrology,
	floBirthDay,
	floBirthMonth,
	floBubbleColor,
	floCategory,
	floClothes,
	floClothesColors,
	floFlooring,
	floFurniture,
	floGames,
	floGender,
	floInterest,
	floName,
	floNameColor,
	floMusic,
	floPersonality,
	floSpecies,
	floStyle,
	floWallpaper}

var (
	floAstrology     = villagersAstrology{villagerAstrologyVirgo}
	floBirthDay      = villagersBirthDay{2}
	floBirthMonth    = villagersBirthMonth{time.September} // 9
	floBubbleColor   = villagersBubbleColor{villagerBubbleColor6B75CE}
	floCategory      = villagersCategory{villagerCategoryA}
	floClothes       = villagersClothes{villagerClothesFolkShirt} // 4446
	floClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorPurple}}
	floFlooring      = villagersFlooring{villagerFlooringIceFlooring}
	floFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRoseBed, villagerFurnitureShowerBooth, villagerFurnitureFrozenPartition, villagerFurnitureFrozenPartition, villagerFurnitureRattanLowTable, villagerFurnitureToilet, villagerFurniturePortableRecordPlayer, villagerFurnitureBlueKitchenMat, villagerFurnitureIronWallRack, villagerFurnitureKitchenIsland, villagerFurnitureFrozenTreatSet, villagerFurniturePotRack, villagerFurnitureIronWallLamp, villagerFurnitureWallMountedTV50In}}
	floGames         = villagersGames{[]VillagerGame{}} // TBD
	floGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	floInterest      = villagersInterest{villagerInterestMusic}
	floName          = villagersName{villagerNameFlo}
	floNameColor     = villagersNameColor{villagerNameColor9AE8DF}
	floMusic         = villagersMusic{villagerMusicSpaceKK}
	floPersonality   = villagersPersonality{villagerPersonalityBigSister}
	floSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	floStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	floWallpaper     = villagersWallpaper{villagerWallpaperIceWall}
)
