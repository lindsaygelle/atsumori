package atsumori

import "time"

// Nan is an Animal Crossing villager.
var Nan = villager{
	nanAstrology,
	nanBirthDay,
	nanBirthMonth,
	nanBubbleColor,
	nanCategory,
	nanClothes,
	nanClothesColors,
	nanFlooring,
	nanFurniture,
	nanGames,
	nanGender,
	nanInterest,
	nanName,
	nanNameColor,
	nanMusic,
	nanPersonality,
	nanSpecies,
	nanStyle,
	nanWallpaper}

var (
	nanAstrology     = villagersAstrology{villagerAstrologyVirgo}
	nanBirthDay      = villagersBirthDay{24}
	nanBirthMonth    = villagersBirthMonth{time.August} // 8
	nanBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	nanCategory      = villagersCategory{villagerCategoryB}
	nanClothes       = villagersClothes{} // 2677
	nanClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorYellow}}
	nanFlooring      = villagersFlooring{villagerFlooringWhiteIronParquetFlooring}
	nanFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGrandPiano, villagerFurniturePianoBench, villagerFurnitureKitchenIsland, villagerFurnitureDoubleSofa, villagerFurnitureRattanEndTable, villagerFurnitureRattanLowTable, villagerFurniturePortableRecordPlayer, villagerFurnitureTissueBox, villagerFurnitureChevresPhoto, villagerFurnitureRattanBed, villagerFurnitureRattanVanity}}
	nanGames         = villagersGames{[]VillagerGame{}} // TBD
	nanGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	nanInterest      = villagersInterest{villagerInterestNature}
	nanName          = villagersName{villagerNameNan}
	nanNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	nanMusic         = villagersMusic{villagerMusicKKEtude}
	nanPersonality   = villagersPersonality{villagerPersonalityNormal}
	nanSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGoat}}
	nanStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	nanWallpaper     = villagersWallpaper{villagerWallpaperBlackBotanicalTileWall}
)
