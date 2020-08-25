package atsumori

import "time"

// Grizzly is an Animal Crossing villager
var Grizzly = villager{
	grizzlyAstrology,
	grizzlyBirthDay,
	grizzlyBirthMonth,
	grizzlyBubbleColor,
	grizzlyCategory,
	grizzlyClothes,
	grizzlyClothesColors,
	grizzlyFlooring,
	grizzlyFurniture,
	grizzlyGames,
	grizzlyGender,
	grizzlyInterest,
	grizzlyName,
	grizzlyNameColor,
	grizzlyMusic,
	grizzlyPersonality,
	grizzlySpecies,
	grizzlyStyle,
	grizzlyWallpaper}

var (
	grizzlyAstrology     = villagersAstrology{villagerAstrologyLeo}
	grizzlyBirthDay      = villagersBirthDay{31}
	grizzlyBirthMonth    = villagersBirthMonth{time.July} // 7
	grizzlyBubbleColor   = villagersBubbleColor{villagerBubbleColor7C6559}
	grizzlyCategory      = villagersCategory{villagerCategoryB}
	grizzlyClothes       = villagersClothes{villagerClothesFlannelShirt} // 3225
	grizzlyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlack}}
	grizzlyFlooring      = villagersFlooring{villagerFlooringSaharahsDesert}
	grizzlyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCampfireCookware, villagerFurnitureLogStool, villagerFurnitureBonfire, villagerFurnitureLantern, villagerFurnitureSleepingBag, villagerFurnitureLogBench, villagerFurnitureCampStove}}
	grizzlyGames         = villagersGames{[]VillagerGame{}} // TBD
	grizzlyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	grizzlyInterest      = villagersInterest{villagerInterestEducation}
	grizzlyName          = villagersName{villagerNameGrizzly}
	grizzlyNameColor     = villagersNameColor{villagerNameColorEAC113}
	grizzlyMusic         = villagersMusic{villagerMusicKKCondor}
	grizzlyPersonality   = villagersPersonality{villagerPersonalityCranky}
	grizzlySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	grizzlyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	grizzlyWallpaper     = villagersWallpaper{villagerWallpaperStarrySkyWall}
)
