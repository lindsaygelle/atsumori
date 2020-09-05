package atsumori

import "time"

// Murphy is an Animal Crossing villager.
var Murphy = villager{
	murphyAstrology,
	murphyBirthDay,
	murphyBirthMonth,
	murphyBubbleColor,
	murphyCategory,
	murphyClothes,
	murphyClothesColors,
	murphyFlooring,
	murphyFurniture,
	murphyGames,
	murphyGender,
	murphyInterest,
	murphyName,
	murphyNameColor,
	murphyMusic,
	murphyPersonality,
	murphySpecies,
	murphyStyle,
	murphyWallpaper}

var (
	murphyAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	murphyBirthDay      = villagersBirthDay{29}
	murphyBirthMonth    = villagersBirthMonth{time.December} // 12
	murphyBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	murphyCategory      = villagersCategory{villagerCategoryA}
	murphyClothes       = villagersClothes{villagerClothesPuffyVest} // 8323
	murphyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	murphyFlooring      = villagersFlooring{villagerFlooringJungleFlooring}
	murphyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWildLogBench, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureTikiTorch, villagerFurnitureLogStool, villagerFurnitureTapeDeck, villagerFurniturePondStone, villagerFurnitureLogRoundTable, villagerFurnitureLantern, villagerFurnitureLogGardenLounge}}
	murphyGames         = villagersGames{[]VillagerGame{}} // TBD
	murphyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	murphyInterest      = villagersInterest{villagerInterestEducation}
	murphyName          = villagersName{villagerNameMurphy}
	murphyNameColor     = villagersNameColor{villagerNameColor28665A}
	murphyMusic         = villagersMusic{villagerMusicKKSafari}
	murphyPersonality   = villagersPersonality{villagerPersonalityCranky}
	murphySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	murphyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	murphyWallpaper     = villagersWallpaper{villagerWallpaperWoodlandWall}
)
