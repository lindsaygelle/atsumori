package atsumori

import "time"

// Coach is an Animal Crossing villager
var Coach = villager{
	coachAstrology,
	coachBirthDay,
	coachBirthMonth,
	coachBubbleColor,
	coachCategory,
	coachClothes,
	coachClothesColors,
	coachFlooring,
	coachFurniture,
	coachGames,
	coachGender,
	coachInterest,
	coachName,
	coachNameColor,
	coachMusic,
	coachPersonality,
	coachSpecies,
	coachStyle,
	coachWallpaper}

var (
	coachAstrology     = villagersAstrology{villagerAstrologyTaurus}
	coachBirthDay      = villagersBirthDay{29}
	coachBirthMonth    = villagersBirthMonth{time.April} // 4
	coachBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	coachCategory      = villagersCategory{villagerCategoryA}
	coachClothes       = villagersClothes{villagerClothesRelayTank} // 3246
	coachClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorRed}}
	coachFlooring      = villagersFlooring{villagerFlooringRacetrackFlooring}
	coachFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBall, villagerFurnitureAirCirculator, villagerFurnitureKettlebell, villagerFurnitureClimbingWall, villagerFurnitureClimbingWall, villagerFurnitureCampingCot, villagerFurnitureKettlebell, villagerFurnitureDirectorsChair, villagerFurnitureBasketballHoop, villagerFurnitureClimbingWall, villagerFurnitureOutdoorTable, villagerFurnitureProteinShakerBottle, villagerFurnitureHandyWaterCooler}}
	coachGames         = villagersGames{[]VillagerGame{}} // TBD
	coachGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	coachInterest      = villagersInterest{villagerInterestFitness}
	coachName          = villagersName{villagerNameCoach}
	coachNameColor     = villagersNameColor{villagerNameColor9B553A}
	coachMusic         = villagersMusic{} // Mr. K.K.
	coachPersonality   = villagersPersonality{villagerPersonalityJock}
	coachSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBull}}
	coachStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleActive}}
	coachWallpaper     = villagersWallpaper{villagerWallpaperChainLinkFence}
)
