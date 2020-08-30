package atsumori

import "time"

// Leonardo is an Animal Crossing villager.
var Leonardo = villager{
	leonardoAstrology,
	leonardoBirthDay,
	leonardoBirthMonth,
	leonardoBubbleColor,
	leonardoCategory,
	leonardoClothes,
	leonardoClothesColors,
	leonardoFlooring,
	leonardoFurniture,
	leonardoGames,
	leonardoGender,
	leonardoInterest,
	leonardoName,
	leonardoNameColor,
	leonardoMusic,
	leonardoPersonality,
	leonardoSpecies,
	leonardoStyle,
	leonardoWallpaper}

var (
	leonardoAstrology     = villagersAstrology{villagerAstrologyTaurus}
	leonardoBirthDay      = villagersBirthDay{15}
	leonardoBirthMonth    = villagersBirthMonth{time.May} // 5
	leonardoBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	leonardoCategory      = villagersCategory{villagerCategoryB}
	leonardoClothes       = villagersClothes{villagerClothesHawkJacket} // 4343
	leonardoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlue}}
	leonardoFlooring      = villagersFlooring{villagerFlooringLeopardPrintFlooring}
	leonardoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRecordBox, villagerFurnitureDinerMiniTable, villagerFurnitureDinerCounterChair, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureIronwoodBed, villagerFurnitureToolShelf, villagerFurnitureRockGuitar, villagerFurnitureEffectsRack, villagerFurniturePedalBoard, villagerFurnitureToolCart, villagerFurnitureCardboardBox, villagerFurnitureToolbox, villagerFurnitureLaptop, villagerFurnitureRetroStereo}}
	leonardoGames         = villagersGames{[]VillagerGame{}} // TBD
	leonardoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	leonardoInterest      = villagersInterest{villagerInterestFitness}
	leonardoName          = villagersName{villagerNameLeonardo}
	leonardoNameColor     = villagersNameColor{villagerNameColor9B553A}
	leonardoMusic         = villagersMusic{villagerMusicKKMetal}
	leonardoPersonality   = villagersPersonality{villagerPersonalityJock}
	leonardoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesTiger}}
	leonardoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleGorgeous}}
	leonardoWallpaper     = villagersWallpaper{villagerWallpaperSkullWall}
)
