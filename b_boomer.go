package atsumori

import "time"

// Boomer is an Animal Crossing villager
var Boomer = villager{
	boomerAstrology,
	boomerBirthDay,
	boomerBirthMonth,
	boomerBubbleColor,
	boomerCategory,
	boomerClothes,
	boomerClothesColors,
	boomerFlooring,
	boomerFurniture,
	boomerGames,
	boomerGender,
	boomerInterest,
	boomerName,
	boomerNameColor,
	boomerMusic,
	boomerPersonality,
	boomerSpecies,
	boomerStyle,
	boomerWallpaper}

var (
	boomerAstrology     = villagersAstrology{villagerAstrologyAquarius}
	boomerBirthDay      = villagersBirthDay{7}
	boomerBirthMonth    = villagersBirthMonth{time.February} // 2
	boomerBubbleColor   = villagersBubbleColor{villagerBubbleColor798040}
	boomerCategory      = villagersCategory{villagerCategoryB}
	boomerClothes       = villagersClothes{} // 4207
	boomerClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorBeige}}
	boomerFlooring      = villagersFlooring{villagerFlooringIcebergFlooring}
	boomerFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFrozenPillar, villagerFurnitureFrozenPillar, villagerFurnitureFrozenPillar, villagerFurnitureFrozenPillar, villagerFurnitureTikiTorch, villagerFurnitureTikiTorch, villagerFurnitureSleigh, villagerFurnitureThreeTieredSnowperson, villagerFurnitureSimpleDIYWorkbench}}
	boomerGames         = villagersGames{[]VillagerGame{}} // TBD
	boomerGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	boomerInterest      = villagersInterest{villagerInterestFitness}
	boomerName          = villagersName{villagerNameBoomer}
	boomerNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	boomerMusic         = villagersMusic{} // Farewell
	boomerPersonality   = villagersPersonality{villagerPersonalityLazy}
	boomerSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	boomerStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	boomerWallpaper     = villagersWallpaper{villagerWallpaperIcebergWall}
)
