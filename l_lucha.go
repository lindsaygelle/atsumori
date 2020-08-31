package atsumori

import "time"

// Lucha is an Animal Crossing villager.
var Lucha = villager{
	luchaAstrology,
	luchaBirthDay,
	luchaBirthMonth,
	luchaBubbleColor,
	luchaCategory,
	luchaClothes,
	luchaClothesColors,
	luchaFlooring,
	luchaFurniture,
	luchaGames,
	luchaGender,
	luchaInterest,
	luchaName,
	luchaNameColor,
	luchaMusic,
	luchaPersonality,
	luchaSpecies,
	luchaStyle,
	luchaWallpaper}

var (
	luchaAstrology     = villagersAstrology{villagerAstrologySagittarius}
	luchaBirthDay      = villagersBirthDay{12}
	luchaBirthMonth    = villagersBirthMonth{time.December} // 12
	luchaBubbleColor   = villagersBubbleColor{villagerBubbleColorFF4040}
	luchaCategory      = villagersCategory{villagerCategoryA}
	luchaClothes       = villagersClothes{villagerClothesAthleticJacket} // 7199
	luchaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	luchaFlooring      = villagersFlooring{villagerFlooringBoxingRingMat}
	luchaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurniturePunchingBag, villagerFurnitureRedCorner, villagerFurnitureNeutralCorner, villagerFurnitureSpeedBag, villagerFurnitureToolCart, villagerFurnitureCassettePlayer}}
	luchaGames         = villagersGames{[]VillagerGame{}} // TBD
	luchaGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	luchaInterest      = villagersInterest{villagerInterestFitness}
	luchaName          = villagersName{villagerNameLucha}
	luchaNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	luchaMusic         = villagersMusic{villagerMusicKKWestern}
	luchaPersonality   = villagersPersonality{villagerPersonalitySmug}
	luchaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	luchaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	luchaWallpaper     = villagersWallpaper{villagerWallpaperRingsideSeating}
)
