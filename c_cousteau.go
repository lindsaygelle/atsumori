package atsumori

import "time"

// Cousteau is an Animal Crossing villager
var Cousteau = villager{
	cousteauAstrology,
	cousteauBirthDay,
	cousteauBirthMonth,
	cousteauBubbleColor,
	cousteauCategory,
	cousteauClothes,
	cousteauClothesColors,
	cousteauFlooring,
	cousteauFurniture,
	cousteauGames,
	cousteauGender,
	cousteauInterest,
	cousteauName,
	cousteauNameColor,
	cousteauMusic,
	cousteauPersonality,
	cousteauSpecies,
	cousteauStyle,
	cousteauWallpaper}

var (
	cousteauAstrology     = villagersAstrology{villagerAstrologySagittarius}
	cousteauBirthDay      = villagersBirthDay{17}
	cousteauBirthMonth    = villagersBirthMonth{time.December} // 12
	cousteauBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	cousteauCategory      = villagersCategory{villagerCategoryB}
	cousteauClothes       = villagersClothes{villagerClothesSilkShirt} // 2672
	cousteauClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorRed}}
	cousteauFlooring      = villagersFlooring{villagerFlooringImperialTile}
	cousteauFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBambooStool, villagerFurnitureBambooStool, villagerFurnitureImperialDiningTable, villagerFurnitureImperialDiningLantern, villagerFurnitureImperialPartition, villagerFurnitureImperialDiningChair, villagerFurnitureImperialLowTable, villagerFurnitureImperialDecorativeShelves, villagerFurnitureGoldfish, villagerFurnitureImperialDiningChair, villagerFurnitureImperialChest, villagerFurniturePhonograph, villagerFurnitureHangingScroll}}
	cousteauGames         = villagersGames{[]VillagerGame{}} // TBD
	cousteauGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	cousteauInterest      = villagersInterest{villagerInterestFitness}
	cousteauName          = villagersName{villagerNameCousteau}
	cousteauNameColor     = villagersNameColor{villagerNameColor9B553A}
	cousteauMusic         = villagersMusic{} // Imperial K.K.
	cousteauPersonality   = villagersPersonality{villagerPersonalityJock}
	cousteauSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	cousteauStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	cousteauWallpaper     = villagersWallpaper{villagerWallpaperExquisiteWall}
)
