package atsumori

import "time"

// Goldie is an Animal Crossing villager.
var Goldie = villager{
	goldieAstrology,
	goldieBirthDay,
	goldieBirthMonth,
	goldieBubbleColor,
	goldieCategory,
	goldieClothes,
	goldieClothesColors,
	goldieFlooring,
	goldieFurniture,
	goldieGames,
	goldieGender,
	goldieInterest,
	goldieName,
	goldieNameColor,
	goldieMusic,
	goldiePersonality,
	goldieSpecies,
	goldieStyle,
	goldieWallpaper}

var (
	goldieAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	goldieBirthDay      = villagersBirthDay{27}
	goldieBirthMonth    = villagersBirthMonth{time.December} // 12
	goldieBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	goldieCategory      = villagersCategory{villagerCategoryB}
	goldieClothes       = villagersClothes{} // 2677
	goldieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorOrange}}
	goldieFlooring      = villagersFlooring{villagerFlooringLightParquetFlooring}
	goldieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurniturePianoBench, villagerFurnitureIvorySmallRoundMat, villagerFurnitureMiniDIYWorkbench, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureWoodenBlockChest, villagerFurnitureWoodenBlockBed, villagerFurnitureTableLamp, villagerFurnitureMacrameTapestry, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenBlockTable, villagerFurnitureUprightPiano}}
	goldieGames         = villagersGames{[]VillagerGame{}} // TBD
	goldieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	goldieInterest      = villagersInterest{villagerInterestNature}
	goldieName          = villagersName{villagerNameGoldie}
	goldieNameColor     = villagersNameColor{villagerNameColor879B96}
	goldieMusic         = villagersMusic{villagerMusicKKBossa}
	goldiePersonality   = villagersPersonality{villagerPersonalityNormal}
	goldieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	goldieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	goldieWallpaper     = villagersWallpaper{villagerWallpaperModernWoodWall}
)
