package atsumori

import "time"

// Margie is an Animal Crossing villager.
var Margie = villager{
	margieAstrology,
	margieBirthDay,
	margieBirthMonth,
	margieBubbleColor,
	margieCategory,
	margieClothes,
	margieClothesColors,
	margieFlooring,
	margieFurniture,
	margieGames,
	margieGender,
	margieInterest,
	margieName,
	margieNameColor,
	margieMusic,
	margiePersonality,
	margieSpecies,
	margieStyle,
	margieWallpaper}

var (
	margieAstrology     = villagersAstrology{villagerAstrologyAquarius}
	margieBirthDay      = villagersBirthDay{28}
	margieBirthMonth    = villagersBirthMonth{time.January} // 1
	margieBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	margieCategory      = villagersCategory{villagerCategoryB}
	margieClothes       = villagersClothes{villagerClothesSilkFloralPrintShirt} // 3248
	margieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorRed}}
	margieFlooring      = villagersFlooring{villagerFlooringBambooFlooring}
	margieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanVanity, villagerFurnitureImperialPartition, villagerFurnitureRattanBed, villagerFurnitureCoconutWallPlanter, villagerFurnitureRattanStool, villagerFurnitureRattanEndTable, villagerFurnitureImperialLowTable, villagerFurnitureRattanWasteBin, villagerFurnitureCushion, villagerFurnitureMonstera, villagerFurnitureIncenseBurner, villagerFurnitureTraditionalTeaSet, villagerFurnitureRattanLowTable, villagerFurniturePortableRecordPlayer}}
	margieGames         = villagersGames{[]VillagerGame{}} // TBD
	margieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	margieInterest      = villagersInterest{villagerInterestEducation}
	margieName          = villagersName{villagerNameMargie}
	margieNameColor     = villagersNameColor{villagerNameColor848484}
	margieMusic         = villagersMusic{villagerMusicKKMarathon}
	margiePersonality   = villagersPersonality{villagerPersonalityNormal}
	margieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesElephant}}
	margieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCute}}
	margieWallpaper     = villagersWallpaper{villagerWallpaperOliveDesertTileWall}
)
