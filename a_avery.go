package atsumori

import (
	"time"
)

// Avery is an Animal Crossing villager.
var Avery = villager{
	averyAstrology,
	averyBirthDay,
	averyBirthMonth,
	averyBubbleColor,
	averyCategory,
	averyClothes,
	averyClothesColors,
	averyFlooring,
	averyFurniture,
	averyGames,
	averyGender,
	averyInterest,
	averyName,
	averyNameColor,
	averyMusic,
	averyPersonality,
	averySpecies,
	averyStyle,
	averyWallpaper}

var (
	averyAstrology     = villagersAstrology{villagerAstrologyPisces}
	averyBirthDay      = villagersBirthDay{22}
	averyBirthMonth    = villagersBirthMonth{time.February} // 2
	averyBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	averyCategory      = villagersCategory{villagerCategoryB}
	averyClothes       = villagersClothes{villagerClothesOversizedShawlOvershirt} // 4154
	averyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorBrown}}
	averyFlooring      = villagersFlooring{villagerFlooringRockyMountainFlooring}
	averyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureTikiTorch, villagerFurnitureTikiTorch, villagerFurnitureBonfire, villagerFurnitureStoneStool, villagerFurniturePondStone, villagerFurnitureBambooSpeaker, villagerFurnitureShantyMat, villagerFurnitureStoneTable, villagerFurnitureUnglazedDishSet, villagerFurnitureClassicPitcher}}
	averyGames         = villagersGames{[]VillagerGame{}} // TBD
	averyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	averyInterest      = villagersInterest{villagerInterestMusic}
	averyName          = villagersName{villagerNameAvery}
	averyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	averyMusic         = villagersMusic{villagerMusicKKCondor}
	averyPersonality   = villagersPersonality{villagerPersonalityCranky}
	averySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	averyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleGorgeous}}
	averyWallpaper     = villagersWallpaper{villagerWallpaperWesternVista}
)
