package atsumori

import "time"

// Buck is an Animal Crossing villager.
var Buck = villager{
	buckAstrology,
	buckBirthDay,
	buckBirthMonth,
	buckBubbleColor,
	buckCategory,
	buckClothes,
	buckClothesColors,
	buckFlooring,
	buckFurniture,
	buckGames,
	buckGender,
	buckInterest,
	buckName,
	buckNameColor,
	buckMusic,
	buckPersonality,
	buckSpecies,
	buckStyle,
	buckWallpaper}

var (
	buckAstrology     = villagersAstrology{villagerAstrologyAries}
	buckBirthDay      = villagersBirthDay{4}
	buckBirthMonth    = villagersBirthMonth{time.April} // 4
	buckBubbleColor   = villagersBubbleColor{villagerBubbleColor798040}
	buckCategory      = villagersCategory{villagerCategoryB}
	buckClothes       = villagersClothes{villagerClothesSweatshirt} // 3671
	buckClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorBrown}}
	buckFlooring      = villagersFlooring{villagerFlooringGreenVinylFlooring}
	buckFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBlockStereo, villagerFurnitureCardboardBox, villagerFurnitureCushion, villagerFurnitureThrowbackWallClock, villagerFurnitureTrainSet, villagerFurnitureWoodenLowTable, villagerFurnitureWoodenBookshelf, villagerFurnitureThrowbackContainer, villagerFurnitureLCDTV50In, villagerFurnitureMug, villagerFurnitureRetroFan}}
	buckGames         = villagersGames{[]VillagerGame{}} // TBD
	buckGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	buckInterest      = villagersInterest{villagerInterestFitness}
	buckName          = villagersName{villagerNameBuck}
	buckNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	buckMusic         = villagersMusic{villagerMusicMyPlace}
	buckPersonality   = villagersPersonality{villagerPersonalityJock}
	buckSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	buckStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	buckWallpaper     = villagersWallpaper{villagerWallpaperPastelDottedWall}
)
