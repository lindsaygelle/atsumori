package atsumori

import "time"

// Mira is an Animal Crossing villager.
var Mira = villager{
	miraAstrology,
	miraBirthDay,
	miraBirthMonth,
	miraBubbleColor,
	miraCategory,
	miraClothes,
	miraClothesColors,
	miraFlooring,
	miraFurniture,
	miraGames,
	miraGender,
	miraInterest,
	miraName,
	miraNameColor,
	miraMusic,
	miraPersonality,
	miraSpecies,
	miraStyle,
	miraWallpaper}

var (
	miraAstrology     = villagersAstrology{villagerAstrologyCancer}
	miraBirthDay      = villagersBirthDay{6}
	miraBirthMonth    = villagersBirthMonth{time.July} // 7
	miraBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF80D}
	miraCategory      = villagersCategory{villagerCategoryA}
	miraClothes       = villagersClothes{} // 9483
	miraClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorYellow}}
	miraFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	miraFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFoldingChair, villagerFurnitureFoldingChair, villagerFurnitureFoldingChair, villagerFurnitureHedgeStandee, villagerFurnitureCassettePlayer, villagerFurnitureSimpleMediumAvocadoMat, villagerFurnitureTreeStandee, villagerFurnitureTreeStandee, villagerFurnitureGrassStandee, villagerFurnitureTreeStandee, villagerFurnitureHedgeStandee, villagerFurnitureThrowbackDinoScreen, villagerFurnitureThrowbackDinoScreen, villagerFurnitureSimpleMediumAvocadoMat, villagerFurnitureGrassStandee, villagerFurnitureHedgeStandee, villagerFurnitureHedgeStandee, villagerFurnitureThrowbackDinoScreen}}
	miraGames         = villagersGames{[]VillagerGame{}} // TBD
	miraGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	miraInterest      = villagersInterest{villagerInterestFitness}
	miraName          = villagersName{villagerNameMira}
	miraNameColor     = villagersNameColor{villagerNameColor9B8986}
	miraMusic         = villagersMusic{villagerMusicGoKKRider}
	miraPersonality   = villagersPersonality{villagerPersonalityBigSister}
	miraSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	miraStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	miraWallpaper     = villagersWallpaper{villagerWallpaperHeavyCurtainWall}
)
