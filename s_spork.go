package atsumori

import "time"

// Spork is an Animal Crossing villager.
var Spork = villager{
	sporkAstrology,
	sporkBirthDay,
	sporkBirthMonth,
	sporkBubbleColor,
	sporkCategory,
	sporkClothes,
	sporkClothesColors,
	sporkFlooring,
	sporkFurniture,
	sporkGames,
	sporkGender,
	sporkInterest,
	sporkName,
	sporkNameColor,
	sporkMusic,
	sporkPersonality,
	sporkSpecies,
	sporkStyle,
	sporkWallpaper}

var (
	sporkAstrology     = villagersAstrology{villagerAstrologyVirgo}
	sporkBirthDay      = villagersBirthDay{3}
	sporkBirthMonth    = villagersBirthMonth{time.September} // 9
	sporkBubbleColor   = villagersBubbleColor{villagerBubbleColorE4B887}
	sporkCategory      = villagersCategory{villagerCategoryB}
	sporkClothes       = villagersClothes{villagerClothesLayeredShirt} // 8550
	sporkClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorColorful}}
	sporkFlooring      = villagersFlooring{villagerFlooringSwampFlooring}
	sporkFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureOldFashionedWashtub, villagerFurnitureLogStool, villagerFurnitureFirewood, villagerFurnitureLogRoundTable, villagerFurnitureLantern, villagerFurnitureBrickOven, villagerFurniturePicnicBasket, villagerFurniturePortableRadio, villagerFurnitureSleepingBag, villagerFurniturePondStone}}
	sporkGames         = villagersGames{[]VillagerGame{}} // TBD
	sporkGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	sporkInterest      = villagersInterest{villagerInterestPlay}
	sporkName          = villagersName{villagerNameSpork}
	sporkNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	sporkMusic         = villagersMusic{villagerMusicForestLife}
	sporkPersonality   = villagersPersonality{villagerPersonalityLazy}
	sporkSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	sporkStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	sporkWallpaper     = villagersWallpaper{villagerWallpaperWoodlandWall}
)
