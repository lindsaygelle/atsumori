package atsumori

import "time"

// Tank is an Animal Crossing villager.
var Tank = villager{
	tankAstrology,
	tankBirthDay,
	tankBirthMonth,
	tankBubbleColor,
	tankCategory,
	tankClothes,
	tankClothesColors,
	tankFlooring,
	tankFurniture,
	tankGames,
	tankGender,
	tankInterest,
	tankName,
	tankNameColor,
	tankMusic,
	tankPersonality,
	tankSpecies,
	tankStyle,
	tankWallpaper}

var (
	tankAstrology     = villagersAstrology{villagerAstrologyTaurus}
	tankBirthDay      = villagersBirthDay{6}
	tankBirthMonth    = villagersBirthMonth{time.May} // 5
	tankBubbleColor   = villagersBubbleColor{villagerBubbleColor8BCDEA}
	tankCategory      = villagersCategory{villagerCategoryB}
	tankClothes       = villagersClothes{villagerClothesNo1Shirt} // 6888
	tankClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorGreen}}
	tankFlooring      = villagersFlooring{villagerFlooringSandlot}
	tankFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBonsaiShelf, villagerFurniturePondStone, villagerFurnitureStoneStool, villagerFurnitureStoneTable, villagerFurniturePortableRadio, villagerFurnitureCypressBathtub, villagerFurnitureFloatingBiotopePlanter}}
	tankGames         = villagersGames{[]VillagerGame{}} // TBD
	tankGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	tankInterest      = villagersInterest{villagerInterestFitness}
	tankName          = villagersName{villagerNameTank}
	tankNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	tankMusic         = villagersMusic{villagerMusicKKLament}
	tankPersonality   = villagersPersonality{villagerPersonalityJock}
	tankSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRhino}}
	tankStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	tankWallpaper     = villagersWallpaper{villagerWallpaperBambooGroveWall}
)
