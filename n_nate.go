package atsumori

import "time"

// Nate is an Animal Crossing villager.
var Nate = villager{
	nateAstrology,
	nateBirthDay,
	nateBirthMonth,
	nateBubbleColor,
	nateCategory,
	nateClothes,
	nateClothesColors,
	nateFlooring,
	nateFurniture,
	nateGames,
	nateGender,
	nateInterest,
	nateName,
	nateNameColor,
	nateMusic,
	natePersonality,
	nateSpecies,
	nateStyle,
	nateWallpaper}

var (
	nateAstrology     = villagersAstrology{villagerAstrologyLeo}
	nateBirthDay      = villagersBirthDay{16}
	nateBirthMonth    = villagersBirthMonth{time.August} // 8
	nateBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	nateCategory      = villagersCategory{villagerCategoryB}
	nateClothes       = villagersClothes{villagerClothesReindeerSweater} // 4566
	nateClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorGreen}}
	nateFlooring      = villagersFlooring{villagerFlooringDaisyMeadow}
	nateFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureStackOfBooks, villagerFurnitureBrickOven, villagerFurnitureTinyLibrary, villagerFurnitureLogStool, villagerFurnitureLogGardenLounge, villagerFurnitureLogRoundTable, villagerFurnitureLogStool, villagerFurnitureBarrel, villagerFurniturePortableRadio, villagerFurnitureWildLogBench}}
	nateGames         = villagersGames{[]VillagerGame{}} // TBD
	nateGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	nateInterest      = villagersInterest{villagerInterestPlay}
	nateName          = villagersName{villagerNameNate}
	nateNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	nateMusic         = villagersMusic{villagerMusicForestLife}
	natePersonality   = villagersPersonality{villagerPersonalityLazy}
	nateSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	nateStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	nateWallpaper     = villagersWallpaper{villagerWallpaperWoodlandWall}
)
