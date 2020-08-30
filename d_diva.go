package atsumori

import "time"

// Diva is an Animal Crossing villager.
var Diva = villager{
	divaAstrology,
	divaBirthDay,
	divaBirthMonth,
	divaBubbleColor,
	divaCategory,
	divaClothes,
	divaClothesColors,
	divaFlooring,
	divaFurniture,
	divaGames,
	divaGender,
	divaInterest,
	divaName,
	divaNameColor,
	divaMusic,
	divaPersonality,
	divaSpecies,
	divaStyle,
	divaWallpaper}

var (
	divaAstrology     = villagersAstrology{villagerAstrologyLibra}
	divaBirthDay      = villagersBirthDay{2}
	divaBirthMonth    = villagersBirthMonth{time.October} // 10
	divaBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	divaCategory      = villagersCategory{villagerCategoryA}
	divaClothes       = villagersClothes{} // 3441
	divaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorLightBlue}}
	divaFlooring      = villagersFlooring{villagerFlooringRockyMountainFlooring}
	divaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCancerTable, villagerFurnitureCancerTable, villagerFurnitureSphinx, villagerFurnitureSphinx, villagerFurnitureGoldenCasket, villagerFurnitureGoldenSeat, villagerFurnitureGoldenCandlestick, villagerFurnitureWallMountedCandle, villagerFurnitureWallMountedCandle, villagerFurnitureWallMountedCandle}}
	divaGames         = villagersGames{[]VillagerGame{}} // TBD
	divaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	divaInterest      = villagersInterest{villagerInterestFitness}
	divaName          = villagersName{villagerNameDiva}
	divaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	divaMusic         = villagersMusic{villagerMusicKKOasis}
	divaPersonality   = villagersPersonality{villagerPersonalityBigSister}
	divaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	divaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	divaWallpaper     = villagersWallpaper{villagerWallpaperAncientWall}
)
