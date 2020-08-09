package atsumori

import "time"

// Admiral is an Animal Crossing villager.
var Admiral Villager = villager{
	admiralAstrology,
	admiralBirthDay,
	admiralBirthMonth,
	admiralBubbleColor,
	admiralCategory,
	admiralClothes,
	admiralClothesColors,
	admiralFlooring,
	admiralFurniture,
	admiralGames,
	admiralGender,
	admiralInterest,
	admiralName,
	admiralNameColor,
	admiralMusic,
	admiralPersonality,
	admiralSpecies,
	admiralStyle,
	admiralWallpaper}

var (
	admiralAstrology     = villagersAstrology{villagerAstrologyAquarius}
	admiralBirthDay      = villagersBirthDay{27}
	admiralBirthMonth    = villagersBirthMonth{time.January}
	admiralBubbleColor   = villagersBubbleColor{villagerBubbleColor0CA54A}
	admiralCategory      = villagersCategory{villagerCategoryA}
	admiralClothes       = villagersClothes{villagerClothesHantenJacket}
	admiralClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorBlue}}
	admiralFlooring      = villagersFlooring{villagerFlooringTatami}
	admiralFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureClayFurnace, villagerFurnitureFirewood, villagerFurnitureFloorSeat, villagerFurnitureKotatsu, villagerFurniturePendulumClock, villagerFurniturePineBonsaiTree, villagerFurnitureRetroFan, villagerFurnitureRotaryPhone, villagerFurnitureTapeDeck, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenStool, villagerFurnitureZenCushion}}
	admiralGames         = villagersGames{[]VillagerGame{}} // TBD
	admiralGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	admiralInterest      = villagersInterest{villagerInterestNature}
	admiralName          = villagersName{villagerNameAdmiral}
	admiralNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	admiralMusic         = villagersMusic{}
	admiralPersonality   = villagersPersonality{villagerPersonalityCranky}
	admiralSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	admiralStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	admiralWallpaper     = villagersWallpaper{villagerWallpaperDirtClodWall}
)
