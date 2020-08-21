package atsumori

import "time"

// Cashmere is an Animal Crossing villager
var Cashmere = villager{
	cashmereAstrology,
	cashmereBirthDay,
	cashmereBirthMonth,
	cashmereBubbleColor,
	cashmereCategory,
	cashmereClothes,
	cashmereClothesColors,
	cashmereFlooring,
	cashmereFurniture,
	cashmereGames,
	cashmereGender,
	cashmereInterest,
	cashmereName,
	cashmereNameColor,
	cashmereMusic,
	cashmerePersonality,
	cashmereSpecies,
	cashmereStyle,
	cashmereWallpaper}

var (
	cashmereAstrology     = villagersAstrology{villagerAstrologyAries}
	cashmereBirthDay      = villagersBirthDay{2}
	cashmereBirthMonth    = villagersBirthMonth{time.April} // 4
	cashmereBubbleColor   = villagersBubbleColor{villagerBubbleColorBFF2FF}
	cashmereCategory      = villagersCategory{villagerCategoryA}
	cashmereClothes       = villagersClothes{villagerClothesSleevelessTunic} // 4599
	cashmereClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBeige}}
	cashmereFlooring      = villagersFlooring{villagerFlooringSlateFlooring}
	cashmereFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFirewood, villagerFurnitureWoodBurningStove, villagerFurnitureWoodenSimpleBed, villagerFurnitureSpinningWheel, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureWoodenStool, villagerFurnitureWoodenMiniTable, villagerFurniturePortableRecordPlayer, villagerFurnitureSimpleMediumPurpleMat, villagerFurnitureWoodenEndTable, villagerFurnitureGlassHolderWithCandle, villagerFurnitureLoom, villagerFurnitureMacrameTapestry}}
	cashmereGames         = villagersGames{[]VillagerGame{}} // TBD
	cashmereGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	cashmereInterest      = villagersInterest{villagerInterestFashion}
	cashmereName          = villagersName{villagerNameCashmere}
	cashmereNameColor     = villagersNameColor{villagerNameColor85A2AF}
	cashmereMusic         = villagersMusic{villagerMusicFarewell}
	cashmerePersonality   = villagersPersonality{villagerPersonalitySnooty}
	cashmereSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	cashmereStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	cashmereWallpaper     = villagersWallpaper{villagerWallpaperRusticStoneWall}
)
