package atsumori

import "time"

// Butch is an Animal Crossing villager
var Butch = villager{
	butchAstrology,
	butchBirthDay,
	butchBirthMonth,
	butchBubbleColor,
	butchCategory,
	butchClothes,
	butchClothesColors,
	butchFlooring,
	butchFurniture,
	butchGames,
	butchGender,
	butchInterest,
	butchName,
	butchNameColor,
	butchMusic,
	butchPersonality,
	butchSpecies,
	butchStyle,
	butchWallpaper}

var (
	butchAstrology     = villagersAstrology{villagerAstrologyScorpio}
	butchBirthDay      = villagersBirthDay{1}
	butchBirthMonth    = villagersBirthMonth{time.November} // 11
	butchBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	butchCategory      = villagersCategory{villagerCategoryB}
	butchClothes       = villagersClothes{villagerClothesArgyleVest} // 3262
	butchClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorGray}}
	butchFlooring      = villagersFlooring{villagerFlooringHighwayFlooring}
	butchFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackRaceCarBed, villagerFurnitureOilBarrel, villagerFurnitureCone, villagerFurniturePhoneBox, villagerFurniturePlainWoodenShopSign, villagerFurnitureConstructionSign, villagerFurnitureCassettePlayer, villagerFurnitureThrowbackRaceCarBed, villagerFurnitureRetroGasPump, villagerFurniturePlasticCanister}}
	butchGames         = villagersGames{[]VillagerGame{}} // TBD
	butchGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	butchInterest      = villagersInterest{villagerInterestMusic}
	butchName          = villagersName{villagerNameButch}
	butchNameColor     = villagersNameColor{villagerNameColor874C25}
	butchMusic         = villagersMusic{} // K.K. Western
	butchPersonality   = villagersPersonality{villagerPersonalityCranky}
	butchSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	butchStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	butchWallpaper     = villagersWallpaper{villagerWallpaperWesternVista}
)
