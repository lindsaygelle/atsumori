package atsumori

import "time"

// Boris is an Animal Crossing villager
var Boris = villager{
	borisAstrology,
	borisBirthDay,
	borisBirthMonth,
	borisBubbleColor,
	borisCategory,
	borisClothes,
	borisClothesColors,
	borisFlooring,
	borisFurniture,
	borisGames,
	borisGender,
	borisInterest,
	borisName,
	borisNameColor,
	borisMusic,
	borisPersonality,
	borisSpecies,
	borisStyle,
	borisWallpaper}

var (
	borisAstrology     = villagersAstrology{villagerAstrologyScorpio}
	borisBirthDay      = villagersBirthDay{6}
	borisBirthMonth    = villagersBirthMonth{time.November} // 11
	borisBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	borisCategory      = villagersCategory{villagerCategoryA}
	borisClothes       = villagersClothes{} // 3606
	borisClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlack}}
	borisFlooring      = villagersFlooring{villagerFlooringRockyMountainFlooring}
	borisFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAquariusUrn, villagerFurnitureRedCarpet, villagerFurnitureRedCarpet, villagerFurnitureCancerTable, villagerFurniturePhonograph, villagerFurnitureGoldenCasket, villagerFurnitureGoldenSeat, villagerFurnitureSphinx, villagerFurnitureSphinx, villagerFurnitureCancerTable, villagerFurnitureCancerTable, villagerFurnitureGoldenCandlestick, villagerFurnitureGoldenCandlestick, villagerFurniturePiscesLamp, villagerFurnitureSagittariusArrow}}
	borisGames         = villagersGames{[]VillagerGame{}} // TBD
	borisGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	borisInterest      = villagersInterest{villagerInterestNature}
	borisName          = villagersName{villagerNameBoris}
	borisNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	borisMusic         = villagersMusic{} // K.K. Oasis
	borisPersonality   = villagersPersonality{villagerPersonalityCranky}
	borisSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	borisStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	borisWallpaper     = villagersWallpaper{villagerWallpaperAncientWall}
)
