package atsumori

import "time"

// Gruff is an Animal Crossing villager.
var Gruff = villager{
	gruffAstrology,
	gruffBirthDay,
	gruffBirthMonth,
	gruffBubbleColor,
	gruffCategory,
	gruffClothes,
	gruffClothesColors,
	gruffFlooring,
	gruffFurniture,
	gruffGames,
	gruffGender,
	gruffInterest,
	gruffName,
	gruffNameColor,
	gruffMusic,
	gruffPersonality,
	gruffSpecies,
	gruffStyle,
	gruffWallpaper}

var (
	gruffAstrology     = villagersAstrology{villagerAstrologyVirgo}
	gruffBirthDay      = villagersBirthDay{29}
	gruffBirthMonth    = villagersBirthMonth{time.August} // 8
	gruffBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	gruffCategory      = villagersCategory{villagerCategoryB}
	gruffClothes       = villagersClothes{villagerClothesOldSchoolJacket} // 7621
	gruffClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlack}}
	gruffFlooring      = villagersFlooring{villagerFlooringGrayVinylFlooring}
	gruffFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDartboard, villagerFurnitureJukebox, villagerFurnitureSilverMic, villagerFurnitureBilliardTable, villagerFurnitureDrumSet, villagerFurnitureDinerCounterTable, villagerFurnitureDinerNeonClock, villagerFurnitureDinerSofa}}
	gruffGames         = villagersGames{[]VillagerGame{}} // TBD
	gruffGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	gruffInterest      = villagersInterest{villagerInterestMusic}
	gruffName          = villagersName{villagerNameGruff}
	gruffNameColor     = villagersNameColor{villagerNameColor28665A}
	gruffMusic         = villagersMusic{villagerMusicKKRockabilly}
	gruffPersonality   = villagersPersonality{villagerPersonalityCranky}
	gruffSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGoat}}
	gruffStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	gruffWallpaper     = villagersWallpaper{villagerWallpaperRusticStoneWall}
)
