package atsumori

import "time"

// Norma is an Animal Crossing villager.
var Norma = villager{
	normaAstrology,
	normaBirthDay,
	normaBirthMonth,
	normaBubbleColor,
	normaCategory,
	normaClothes,
	normaClothesColors,
	normaFlooring,
	normaFurniture,
	normaGames,
	normaGender,
	normaInterest,
	normaName,
	normaNameColor,
	normaMusic,
	normaPersonality,
	normaSpecies,
	normaStyle,
	normaWallpaper}

var (
	normaAstrology     = villagersAstrology{villagerAstrologyVirgo}
	normaBirthDay      = villagersBirthDay{20}
	normaBirthMonth    = villagersBirthMonth{time.September} // 9
	normaBubbleColor   = villagersBubbleColor{villagerBubbleColorF2BDC7}
	normaCategory      = villagersCategory{villagerCategoryA}
	normaClothes       = villagersClothes{} // 4399
	normaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorLightBlue}}
	normaFlooring      = villagersFlooring{villagerFlooringDaisyMeadow}
	normaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFirewood, villagerFurnitureHayBed, villagerFurnitureLogRoundTable, villagerFurnitureButterChurn, villagerFurnitureBrickOven, villagerFurnitureLogBench, villagerFurniturePicnicBasket, villagerFurnitureBarrel, villagerFurnitureCuttingBoard}}
	normaGames         = villagersGames{[]VillagerGame{}} // TBD
	normaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	normaInterest      = villagersInterest{villagerInterestNature}
	normaName          = villagersName{villagerNameNorma}
	normaNameColor     = villagersNameColor{villagerNameColor634B4B}
	normaMusic         = villagersMusic{villagerMusicMountainSong}
	normaPersonality   = villagersPersonality{villagerPersonalityNormal}
	normaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCow}}
	normaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	normaWallpaper     = villagersWallpaper{villagerWallpaperMeadowVista}
)
