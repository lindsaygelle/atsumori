package atsumori

import "time"

// Patty is an Animal Crossing villager.
var Patty = villager{
	pattyAstrology,
	pattyBirthDay,
	pattyBirthMonth,
	pattyBubbleColor,
	pattyCategory,
	pattyClothes,
	pattyClothesColors,
	pattyFlooring,
	pattyFurniture,
	pattyGames,
	pattyGender,
	pattyInterest,
	pattyName,
	pattyNameColor,
	pattyMusic,
	pattyPersonality,
	pattySpecies,
	pattyStyle,
	pattyWallpaper}

var (
	pattyAstrology     = villagersAstrology{villagerAstrologyTaurus}
	pattyBirthDay      = villagersBirthDay{10}
	pattyBirthMonth    = villagersBirthMonth{time.May} // 5
	pattyBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	pattyCategory      = villagersCategory{villagerCategoryB}
	pattyClothes       = villagersClothes{} // 3288
	pattyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorRed}}
	pattyFlooring      = villagersFlooring{villagerFlooringWildflowerMeadow}
	pattyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBirdcage, villagerFurnitureHammock, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureNaturalGardenTable, villagerFurnitureClotheslinePole, villagerFurnitureNaturalGardenChair, villagerFurnitureCypressPlant}}
	pattyGames         = villagersGames{[]VillagerGame{}} // TBD
	pattyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pattyInterest      = villagersInterest{villagerInterestFashion}
	pattyName          = villagersName{villagerNamePatty}
	pattyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	pattyMusic         = villagersMusic{villagerMusicKKCalypso}
	pattyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	pattySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCow}}
	pattyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	pattyWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
