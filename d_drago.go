package atsumori

import "time"

// Drago is an Animal Crossing villager.
var Drago = villager{
	dragoAstrology,
	dragoBirthDay,
	dragoBirthMonth,
	dragoBubbleColor,
	dragoCategory,
	dragoClothes,
	dragoClothesColors,
	dragoFlooring,
	dragoFurniture,
	dragoGames,
	dragoGender,
	dragoInterest,
	dragoName,
	dragoNameColor,
	dragoMusic,
	dragoPersonality,
	dragoSpecies,
	dragoStyle,
	dragoWallpaper}

var (
	dragoAstrology     = villagersAstrology{villagerAstrologyAquarius}
	dragoBirthDay      = villagersBirthDay{12}
	dragoBirthMonth    = villagersBirthMonth{time.February} // 2
	dragoBubbleColor   = villagersBubbleColor{villagerBubbleColor78DD62}
	dragoCategory      = villagersCategory{villagerCategoryA}
	dragoClothes       = villagersClothes{} // 3462
	dragoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorRed}}
	dragoFlooring      = villagersFlooring{villagerFlooringImperialTile}
	dragoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureImperialPartition, villagerFurnitureCherryBlossomBranches, villagerFurnitureCherryBlossomBranches, villagerFurnitureImperialDiningChair, villagerFurnitureImperialDiningTable, villagerFurnitureImperialDiningLantern, villagerFurnitureImperialDiningLantern, villagerFurnitureImperialChest, villagerFurniturePhonograph, villagerFurnitureImperialDecorativeShelves, villagerFurnitureImperialChest, villagerFurnitureSeaHorseModel}}
	dragoGames         = villagersGames{[]VillagerGame{}} // TBD
	dragoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	dragoInterest      = villagersInterest{villagerInterestNature}
	dragoName          = villagersName{villagerNameDrago}
	dragoNameColor     = villagersNameColor{villagerNameColor28665A}
	dragoMusic         = villagersMusic{villagerMusicImperialKK}
	dragoPersonality   = villagersPersonality{villagerPersonalityLazy}
	dragoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAlligator}}
	dragoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	dragoWallpaper     = villagersWallpaper{villagerWallpaperImperialWall}
)
