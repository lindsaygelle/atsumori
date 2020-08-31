package atsumori

import "time"

// Mac is an Animal Crossing villager.
var Mac = villager{
	macAstrology,
	macBirthDay,
	macBirthMonth,
	macBubbleColor,
	macCategory,
	macClothes,
	macClothesColors,
	macFlooring,
	macFurniture,
	macGames,
	macGender,
	macInterest,
	macName,
	macNameColor,
	macMusic,
	macPersonality,
	macSpecies,
	macStyle,
	macWallpaper}

var (
	macAstrology     = villagersAstrology{villagerAstrologyScorpio}
	macBirthDay      = villagersBirthDay{11}
	macBirthMonth    = villagersBirthMonth{time.November} // 11
	macBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	macCategory      = villagersCategory{villagerCategoryB}
	macClothes       = villagersClothes{villagerClothesRaglanShirt} // 2674
	macClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlack}}
	macFlooring      = villagersFlooring{villagerFlooringSandlot}
	macFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGardenFaucet, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureDoghouse, villagerFurniturePlasticPool, villagerFurniturePetBed, villagerFurnitureClothesline, villagerFurnitureHammock, villagerFurnitureHoseReel, villagerFurnitureOldFashionedWashtub, villagerFurnitureLogBench, villagerFurnitureLantern, villagerFurniturePortableRadio}}
	macGames         = villagersGames{[]VillagerGame{}} // TBD
	macGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	macInterest      = villagersInterest{villagerInterestFitness}
	macName          = villagersName{villagerNameMac}
	macNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	macMusic         = villagersMusic{villagerMusicMyPlace}
	macPersonality   = villagersPersonality{villagerPersonalityJock}
	macSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	macStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	macWallpaper     = villagersWallpaper{villagerWallpaperIvyWall}
)
