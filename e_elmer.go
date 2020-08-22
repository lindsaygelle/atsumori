package atsumori

import "time"

// Elmer is an Animal Crossing villager
var Elmer = villager{
	elmerAstrology,
	elmerBirthDay,
	elmerBirthMonth,
	elmerBubbleColor,
	elmerCategory,
	elmerClothes,
	elmerClothesColors,
	elmerFlooring,
	elmerFurniture,
	elmerGames,
	elmerGender,
	elmerInterest,
	elmerName,
	elmerNameColor,
	elmerMusic,
	elmerPersonality,
	elmerSpecies,
	elmerStyle,
	elmerWallpaper}

var (
	elmerAstrology     = villagersAstrology{villagerAstrologyLibra}
	elmerBirthDay      = villagersBirthDay{5}
	elmerBirthMonth    = villagersBirthMonth{time.October} // 10
	elmerBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	elmerCategory      = villagersCategory{villagerCategoryB}
	elmerClothes       = villagersClothes{villagerClothesBoaFleece} // 3705
	elmerClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	elmerFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	elmerFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLogStool, villagerFurnitureFirewood, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureHayBed, villagerFurnitureWoodenBucket, villagerFurnitureGardenFaucet, villagerFurnitureWildLogBench, villagerFurnitureLogRoundTable, villagerFurnitureClothesline, villagerFurnitureCassettePlayer}}
	elmerGames         = villagersGames{[]VillagerGame{}} // TBD
	elmerGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	elmerInterest      = villagersInterest{villagerInterestPlay}
	elmerName          = villagersName{villagerNameElmer}
	elmerNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	elmerMusic         = villagersMusic{villagerMusicKKCountry}
	elmerPersonality   = villagersPersonality{villagerPersonalityLazy}
	elmerSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	elmerStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	elmerWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
