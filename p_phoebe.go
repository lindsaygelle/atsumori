package atsumori

import "time"

// Phoebe is an Animal Crossing villager.
var Phoebe = villager{
	phoebeAstrology,
	phoebeBirthDay,
	phoebeBirthMonth,
	phoebeBubbleColor,
	phoebeCategory,
	phoebeClothes,
	phoebeClothesColors,
	phoebeFlooring,
	phoebeFurniture,
	phoebeGames,
	phoebeGender,
	phoebeInterest,
	phoebeName,
	phoebeNameColor,
	phoebeMusic,
	phoebePersonality,
	phoebeSpecies,
	phoebeStyle,
	phoebeWallpaper}

var (
	phoebeAstrology     = villagersAstrology{villagerAstrologyTaurus}
	phoebeBirthDay      = villagersBirthDay{22}
	phoebeBirthMonth    = villagersBirthMonth{time.April} // 4
	phoebeBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	phoebeCategory      = villagersCategory{villagerCategoryA}
	phoebeClothes       = villagersClothes{villagerClothesFrontTieTee} // 4320
	phoebeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorRed}}
	phoebeFlooring      = villagersFlooring{villagerFlooringLavaFlooring}
	phoebeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBonfire, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureTikiTorch, villagerFurnitureCampingCot, villagerFurnitureLogStool, villagerFurnitureCuteMusicPlayer, villagerFurnitureLogBench, villagerFurnitureMug}}
	phoebeGames         = villagersGames{[]VillagerGame{}} // TBD
	phoebeGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	phoebeInterest      = villagersInterest{villagerInterestFitness}
	phoebeName          = villagersName{villagerNamePhoebe}
	phoebeNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	phoebeMusic         = villagersMusic{villagerMusicKKFlamenco}
	phoebePersonality   = villagersPersonality{villagerPersonalityBigSister}
	phoebeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	phoebeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCool}}
	phoebeWallpaper     = villagersWallpaper{villagerWallpaperMagmaCavernWall}
)
