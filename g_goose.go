package atsumori

import "time"

// Goose is an Animal Crossing villager
var Goose = villager{
	gooseAstrology,
	gooseBirthDay,
	gooseBirthMonth,
	gooseBubbleColor,
	gooseCategory,
	gooseClothes,
	gooseClothesColors,
	gooseFlooring,
	gooseFurniture,
	gooseGames,
	gooseGender,
	gooseInterest,
	gooseName,
	gooseNameColor,
	gooseMusic,
	goosePersonality,
	gooseSpecies,
	gooseStyle,
	gooseWallpaper}

var (
	gooseAstrology     = villagersAstrology{villagerAstrologyLibra}
	gooseBirthDay      = villagersBirthDay{4}
	gooseBirthMonth    = villagersBirthMonth{time.October} // 10
	gooseBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	gooseCategory      = villagersCategory{villagerCategoryB}
	gooseClothes       = villagersClothes{villagerClothesBoldAlohaShirt} // 8206
	gooseClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorWhite}}
	gooseFlooring      = villagersFlooring{villagerFlooringArtsyParquetFlooring}
	gooseFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenTable, villagerFurnitureWoodenChair, villagerFurnitureMiniDIYWorkbench, villagerFurnitureWoodenSimpleBed, villagerFurnitureGasRange, villagerFurnitureRefrigerator, villagerFurnitureWoodenEndTable, villagerFurnitureTapeDeck, villagerFurnitureWoodenChest, villagerFurniturePennant, villagerFurnitureMiniCactusSet}}
	gooseGames         = villagersGames{[]VillagerGame{}} // TBD
	gooseGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	gooseInterest      = villagersInterest{villagerInterestFitness}
	gooseName          = villagersName{villagerNameGoose}
	gooseNameColor     = villagersNameColor{villagerNameColor848484}
	gooseMusic         = villagersMusic{villagerMusicKKCountry}
	goosePersonality   = villagersPersonality{villagerPersonalityJock}
	gooseSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	gooseStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	gooseWallpaper     = villagersWallpaper{villagerWallpaperGreenPaintedWoodWall}
)
