package atsumori

import "time"

// Zucker is an Animal Crossing villager.
var Zucker = villager{
	zuckerAstrology,
	zuckerBirthDay,
	zuckerBirthMonth,
	zuckerBubbleColor,
	zuckerCategory,
	zuckerClothes,
	zuckerClothesColors,
	zuckerFlooring,
	zuckerFurniture,
	zuckerGames,
	zuckerGender,
	zuckerInterest,
	zuckerName,
	zuckerNameColor,
	zuckerMusic,
	zuckerPersonality,
	zuckerSpecies,
	zuckerStyle,
	zuckerWallpaper}

var (
	zuckerAstrology     = villagersAstrology{villagerAstrologyPisces} // villagerAstrology
	zuckerBirthDay      = villagersBirthDay{8}
	zuckerBirthMonth    = villagersBirthMonth{time.March} // 3
	zuckerBubbleColor   = villagersBubbleColor{villagerBubbleColorE4B887}
	zuckerCategory      = villagersCategory{villagerCategoryA}
	zuckerClothes       = villagersClothes{villagerClothesHappiTee} // 3254
	zuckerClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorYellow}}
	zuckerFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	zuckerFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLogBench, villagerFurnitureCottonCandyStall, villagerFurnitureOutdoorGenerator, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurnitureCoolerBox, villagerFurniturePartyGarland, villagerFurnitureOutdoorTable, villagerFurnitureStall, villagerFurniturePortableRadio, villagerFurnitureStandMixer, villagerFurnitureCuttingBoard, villagerFurnitureSteamerBasketSet, villagerFurnitureCardboardBox, villagerFurnitureCardboardBox}}
	zuckerGames         = villagersGames{[]VillagerGame{}} // TBD
	zuckerGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	zuckerInterest      = villagersInterest{villagerInterestNature}
	zuckerName          = villagersName{villagerNameZucker}
	zuckerNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	zuckerMusic         = villagersMusic{villagerMusicSpringBlossoms} // Spring Blossoms
	zuckerPersonality   = villagersPersonality{villagerPersonalityLazy}
	zuckerSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOctopus}}
	zuckerStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	zuckerWallpaper     = villagersWallpaper{villagerWallpaperChainLinkFence}
)
