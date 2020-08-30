package atsumori

import "time"

// KidCat is an Animal Crossing villager.
var KidCat = villager{
	kidCatAstrology,
	kidCatBirthDay,
	kidCatBirthMonth,
	kidCatBubbleColor,
	kidCatCategory,
	kidCatClothes,
	kidCatClothesColors,
	kidCatFlooring,
	kidCatFurniture,
	kidCatGames,
	kidCatGender,
	kidCatInterest,
	kidCatName,
	kidCatNameColor,
	kidCatMusic,
	kidCatPersonality,
	kidCatSpecies,
	kidCatStyle,
	kidCatWallpaper}

var (
	kidCatAstrology     = villagersAstrology{villagerAstrologyLeo}
	kidCatBirthDay      = villagersBirthDay{1}
	kidCatBirthMonth    = villagersBirthMonth{time.August} // 8
	kidCatBubbleColor   = villagersBubbleColor{villagerBubbleColorFF4040}
	kidCatCategory      = villagersCategory{villagerCategoryB}
	kidCatClothes       = villagersClothes{villagerClothesNo1Shirt} // 6888
	kidCatClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorRed}}
	kidCatFlooring      = villagersFlooring{villagerFlooringSandlot}
	kidCatFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurnitureGarbagePail, villagerFurniturePullUpBarStand, villagerFurnitureOilBarrel, villagerFurnitureManholeCover, villagerFurnitureMountainBike, villagerFurnitureBasketballHoop, villagerFurnitureCardboardBox, villagerFurnitureTapeDeck}}
	kidCatGames         = villagersGames{[]VillagerGame{}} // TBD
	kidCatGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	kidCatInterest      = villagersInterest{villagerInterestFitness}
	kidCatName          = villagersName{villagerNameKidCat}
	kidCatNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	kidCatMusic         = villagersMusic{villagerMusicGoKKRider}
	kidCatPersonality   = villagersPersonality{villagerPersonalityJock}
	kidCatSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	kidCatStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	kidCatWallpaper     = villagersWallpaper{villagerWallpaperChainLinkFence}
)
