package atsumori

import "time"

// Sheldon is an Animal Crossing villager.
var Sheldon = villager{
	sheldonAstrology,
	sheldonBirthDay,
	sheldonBirthMonth,
	sheldonBubbleColor,
	sheldonCategory,
	sheldonClothes,
	sheldonClothesColors,
	sheldonFlooring,
	sheldonFurniture,
	sheldonGames,
	sheldonGender,
	sheldonInterest,
	sheldonName,
	sheldonNameColor,
	sheldonMusic,
	sheldonPersonality,
	sheldonSpecies,
	sheldonStyle,
	sheldonWallpaper}

var (
	sheldonAstrology     = villagersAstrology{villagerAstrologyPisces}
	sheldonBirthDay      = villagersBirthDay{26}
	sheldonBirthMonth    = villagersBirthMonth{time.February} // 2
	sheldonBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	sheldonCategory      = villagersCategory{villagerCategoryA}
	sheldonClothes       = villagersClothes{villagerClothesTigerJacket} // 3253
	sheldonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorYellow}}
	sheldonFlooring      = villagersFlooring{villagerFlooringForestFlooring}
	sheldonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureLeafCampfire, villagerFurnitureMushPartition, villagerFurnitureMushLowStool, villagerFurnitureMushParasol, villagerFurnitureMushLamp, villagerFurnitureLogGardenLounge, villagerFurnitureMushLamp, villagerFurnitureMushTable, villagerFurnitureMug, villagerFurnitureMushLog, villagerFurnitureTapeDeck}}
	sheldonGames         = villagersGames{[]VillagerGame{}} // TBD
	sheldonGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	sheldonInterest      = villagersInterest{villagerInterestPlay}
	sheldonName          = villagersName{villagerNameSheldon}
	sheldonNameColor     = villagersNameColor{villagerNameColor874C25}
	sheldonMusic         = villagersMusic{villagerMusicKKCountry}
	sheldonPersonality   = villagersPersonality{villagerPersonalityJock}
	sheldonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	sheldonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	sheldonWallpaper     = villagersWallpaper{villagerWallpaperForestWall}
)
