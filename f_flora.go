package atsumori

import "time"

// Flora is an Animal Crossing villager
var Flora = villager{
	floraAstrology,
	floraBirthDay,
	floraBirthMonth,
	floraBubbleColor,
	floraCategory,
	floraClothes,
	floraClothesColors,
	floraFlooring,
	floraFurniture,
	floraGames,
	floraGender,
	floraInterest,
	floraName,
	floraNameColor,
	floraMusic,
	floraPersonality,
	floraSpecies,
	floraStyle,
	floraWallpaper}

var (
	floraAstrology     = villagersAstrology{villagerAstrologyAquarius}
	floraBirthDay      = villagersBirthDay{9}
	floraBirthMonth    = villagersBirthMonth{time.February} // 2
	floraBubbleColor   = villagersBubbleColor{villagerBubbleColorFF6183}
	floraCategory      = villagersCategory{villagerCategoryA}
	floraClothes       = villagersClothes{} // 2686
	floraClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorPink}}
	floraFlooring      = villagersFlooring{villagerFlooringOasisFlooring}
	floraFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWildLogBench, villagerFurnitureMrFlamingo, villagerFurnitureMrsFlamingo, villagerFurnitureMrFlamingo, villagerFurnitureMrsFlamingo, villagerFurnitureSleepingBag, villagerFurnitureLogBench, villagerFurniturePortableRadio, villagerFurnitureCampfire}}
	floraGames         = villagersGames{[]VillagerGame{}} // TBD
	floraGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	floraInterest      = villagersInterest{villagerInterestPlay}
	floraName          = villagersName{villagerNameFlora}
	floraNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	floraMusic         = villagersMusic{villagerMusicKKCondor}
	floraPersonality   = villagersPersonality{villagerPersonalityPeppy}
	floraSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	floraStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleActive}}
	floraWallpaper     = villagersWallpaper{villagerWallpaperMeadowVista}
)
