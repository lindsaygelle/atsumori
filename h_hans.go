package atsumori

import "time"

// Hans is an Animal Crossing villager.
var Hans = villager{
	hansAstrology,
	hansBirthDay,
	hansBirthMonth,
	hansBubbleColor,
	hansCategory,
	hansClothes,
	hansClothesColors,
	hansFlooring,
	hansFurniture,
	hansGames,
	hansGender,
	hansInterest,
	hansName,
	hansNameColor,
	hansMusic,
	hansPersonality,
	hansSpecies,
	hansStyle,
	hansWallpaper}

var (
	hansAstrology     = villagersAstrology{villagerAstrologySagittarius}
	hansBirthDay      = villagersBirthDay{5}
	hansBirthMonth    = villagersBirthMonth{time.December} // 12
	hansBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	hansCategory      = villagersCategory{villagerCategoryA}
	hansClothes       = villagersClothes{villagerClothesDownSkiJacket} // 4735
	hansClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorBlue}}
	hansFlooring      = villagersFlooring{villagerFlooringSkiSlopeFlooring}
	hansFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCampfire, villagerFurnitureFirewood, villagerFurnitureSleigh, villagerFurnitureTreeStandee, villagerFurnitureTreeStandee, villagerFurnitureWildLogBench, villagerFurnitureTreeStandee, villagerFurnitureSleepingBag, villagerFurnitureBonfire, villagerFurnitureLogBench, villagerFurnitureGrassStandee, villagerFurnitureGrassStandee, villagerFurnitureCampStove, villagerFurniturePortableRadio}}
	hansGames         = villagersGames{[]VillagerGame{}} // TBD
	hansGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	hansInterest      = villagersInterest{villagerInterestFitness}
	hansName          = villagersName{villagerNameHans}
	hansNameColor     = villagersNameColor{villagerNameColor848484}
	hansMusic         = villagersMusic{villagerMusicKKSonata}
	hansPersonality   = villagersPersonality{villagerPersonalitySmug}
	hansSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGorilla}}
	hansStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	hansWallpaper     = villagersWallpaper{villagerWallpaperSkiSlopeWall}
)
