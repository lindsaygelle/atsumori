package atsumori

import "time"

// Olaf is an Animal Crossing villager.
var Olaf = villager{
	olafAstrology,
	olafBirthDay,
	olafBirthMonth,
	olafBubbleColor,
	olafCategory,
	olafClothes,
	olafClothesColors,
	olafFlooring,
	olafFurniture,
	olafGames,
	olafGender,
	olafInterest,
	olafName,
	olafNameColor,
	olafMusic,
	olafPersonality,
	olafSpecies,
	olafStyle,
	olafWallpaper}

var (
	olafAstrology     = villagersAstrology{villagerAstrologyTaurus}
	olafBirthDay      = villagersBirthDay{19}
	olafBirthMonth    = villagersBirthMonth{time.May} // 5
	olafBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	olafCategory      = villagersCategory{villagerCategoryA}
	olafClothes       = villagersClothes{} // 3613
	olafClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlack}}
	olafFlooring      = villagersFlooring{villagerFlooringSimpleRedFlooring}
	olafFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBed, villagerFurnitureAntiqueConsoleTable, villagerFurnitureStudioWallSpotlight, villagerFurnitureWhirlpoolBath, villagerFurnitureBathroomTowelRack, villagerFurnitureAntiqueChair, villagerFurnitureAntiqueClock, villagerFurnitureBilliardTable, villagerFurnitureDoubleSofa}}
	olafGames         = villagersGames{[]VillagerGame{}} // TBD
	olafGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	olafInterest      = villagersInterest{villagerInterestEducation}
	olafName          = villagersName{villagerNameOlaf}
	olafNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	olafMusic         = villagersMusic{villagerMusicKKMilonga}
	olafPersonality   = villagersPersonality{villagerPersonalitySmug}
	olafSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAnteater}}
	olafStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	olafWallpaper     = villagersWallpaper{villagerWallpaperBlackBrickWall}
)
