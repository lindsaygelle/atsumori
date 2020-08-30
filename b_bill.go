package atsumori

import "time"

// Bill is an Animal Crossing villager.
var Bill = villager{
	billAstrology,
	billBirthDay,
	billBirthMonth,
	billBubbleColor,
	billCategory,
	billClothes,
	billClothesColors,
	billFlooring,
	billFurniture,
	billGames,
	billGender,
	billInterest,
	billName,
	billNameColor,
	billMusic,
	billPersonality,
	billSpecies,
	billStyle,
	billWallpaper}

var (
	billAstrology     = villagersAstrology{villagerAstrologyAquarius}
	billBirthDay      = villagersBirthDay{1}
	billBirthMonth    = villagersBirthMonth{time.February} // 2
	billBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	billCategory      = villagersCategory{villagerCategoryA}
	billClothes       = villagersClothes{villagerClothesNo2Shirt} // 2727
	billClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorPurple}}
	billFlooring      = villagersFlooring{villagerFlooringBambooFlooring}
	billFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureClayFurnace, villagerFurnitureImperialPartition, villagerFurniturePotRack, villagerFurnitureImperialChest, villagerFurnitureRedCarpet, villagerFurniturePhonograph, villagerFurnitureCherryBlossomBranches, villagerFurnitureCherryBlossomPondStone, villagerFurnitureImperialDecorativeShelves, villagerFurnitureBambooBench, villagerFurnitureSteamerBasketSet, villagerFurnitureSteamerBasketSet, villagerFurnitureImperialDiningTable, villagerFurnitureImperialDiningChair, villagerFurnitureImperialDiningChair, villagerFurnitureImperialDiningLantern}}
	billGames         = villagersGames{[]VillagerGame{}} // TBD
	billGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	billInterest      = villagersInterest{villagerInterestPlay}
	billName          = villagersName{villagerNameBill}
	billNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	billMusic         = villagersMusic{villagerMusicImperialKK}
	billPersonality   = villagersPersonality{villagerPersonalityJock}
	billSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	billStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleActive}}
	billWallpaper     = villagersWallpaper{villagerWallpaperBambooScreenWall}
)
