package atsumori

import "time"

// Lucky is an Animal Crossing villager.
var Lucky = villager{
	luckyAstrology,
	luckyBirthDay,
	luckyBirthMonth,
	luckyBubbleColor,
	luckyCategory,
	luckyClothes,
	luckyClothesColors,
	luckyFlooring,
	luckyFurniture,
	luckyGames,
	luckyGender,
	luckyInterest,
	luckyName,
	luckyNameColor,
	luckyMusic,
	luckyPersonality,
	luckySpecies,
	luckyStyle,
	luckyWallpaper}

var (
	luckyAstrology     = villagersAstrology{villagerAstrologyScorpio}
	luckyBirthDay      = villagersBirthDay{4}
	luckyBirthMonth    = villagersBirthMonth{time.November} // 11
	luckyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	luckyCategory      = villagersCategory{villagerCategoryB}
	luckyClothes       = villagersClothes{villagerClothesOpenCollarShirt} // 4715
	luckyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorWhite}}
	luckyFlooring      = villagersFlooring{villagerFlooringSwampFlooring}
	luckyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureTinBucket, villagerFurnitureTikiTorch, villagerFurnitureTikiTorch, villagerFurnitureWesternStyleStone, villagerFurnitureWesternStyleStone, villagerFurnitureWesternStyleStone, villagerFurnitureRedLeafPile, villagerFurnitureSkeleton, villagerFurnitureThrowbackSkullRadio}}
	luckyGames         = villagersGames{[]VillagerGame{}} // TBD
	luckyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	luckyInterest      = villagersInterest{villagerInterestPlay}
	luckyName          = villagersName{villagerNameLucky}
	luckyNameColor     = villagersNameColor{villagerNameColor848484}
	luckyMusic         = villagersMusic{villagerMusicKKDirge}
	luckyPersonality   = villagersPersonality{villagerPersonalityLazy}
	luckySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	luckyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	luckyWallpaper     = villagersWallpaper{villagerWallpaperRamshackleWall}
)
