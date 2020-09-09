package atsumori

import "time"

// OHare is an Animal Crossing villager.
var OHare = villager{
	oHareAstrology,
	oHareBirthDay,
	oHareBirthMonth,
	oHareBubbleColor,
	oHareCategory,
	oHareClothes,
	oHareClothesColors,
	oHareFlooring,
	oHareFurniture,
	oHareGames,
	oHareGender,
	oHareInterest,
	oHareName,
	oHareNameColor,
	oHareMusic,
	oHarePersonality,
	oHareSpecies,
	oHareStyle,
	oHareWallpaper}

var (
	oHareAstrology     = villagersAstrology{villagerAstrologyLeo}
	oHareBirthDay      = villagersBirthDay{24}
	oHareBirthMonth    = villagersBirthMonth{time.July} // 7
	oHareBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	oHareCategory      = villagersCategory{villagerCategoryB}
	oHareClothes       = villagersClothes{villagerClothesPineappleAlohaShirt} // 3237
	oHareClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	oHareFlooring      = villagersFlooring{villagerFlooringWaterFlooring}
	oHareFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLifeRing, villagerFurnitureWaveBreaker, villagerFurniturePalmTreeLamp, villagerFurnitureBeachTowel, villagerFurnitureBeachBall, villagerFurnitureCoconutJuice, villagerFurnitureCassettePlayer, villagerFurnitureOutdoorsyFishingRod, villagerFurnitureTropicalRug, villagerFurnitureGreatWhiteSharkModel, villagerFurnitureCoolerBox, villagerFurnitureTropicalRug}}
	oHareGames         = villagersGames{[]VillagerGame{}} // TBD
	oHareGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	oHareInterest      = villagersInterest{villagerInterestNature}
	oHareName          = villagersName{villagerNameOHare}
	oHareNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	oHareMusic         = villagersMusic{villagerMusicKKIsland}
	oHarePersonality   = villagersPersonality{villagerPersonalitySmug}
	oHareSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	oHareStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	oHareWallpaper     = villagersWallpaper{villagerWallpaperOceanHorizonWall}
)
