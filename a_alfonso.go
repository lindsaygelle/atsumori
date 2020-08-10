package atsumori

// Alfonso is an Animal Crossing villager.
var Alfonso Villager = villager{
	alfonsoAstrology,
	alfonsoBirthDay,
	alfonsoBirthMonth,
	alfonsoBubbleColor,
	alfonsoCategory,
	alfonsoClothes,
	alfonsoClothesColors,
	alfonsoFlooring,
	alfonsoFurniture,
	alfonsoGames,
	alfonsoGender,
	alfonsoInterest,
	alfonsoName,
	alfonsoNameColor,
	alfonsoMusic,
	alfonsoPersonality,
	alfonsoSpecies,
	alfonsoStyle,
	alfonsoWallpaper}

var (
	alfonsoAstrology     = villagersAstrology{}
	alfonsoBirthDay      = villagersBirthDay{}
	alfonsoBirthMonth    = villagersBirthMonth{}
	alfonsoBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	alfonsoCategory      = villagersCategory{villagerCategoryB}
	alfonsoClothes       = villagersClothes{villagerClothesSimpleParka}
	alfonsoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorBlue}}
	alfonsoFlooring      = villagersFlooring{villagerFlooringGreenHoneycombTile}
	alfonsoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureStandingToilet, villagerFurnitureThrowbackDinoScreen, villagerFurnitureThrowbackRocket, villagerFurnitureToyBox, villagerFurnitureTrainSet, villagerFurnitureWoodenBlockBed, villagerFurnitureWoodenBlockStereo, villagerFurnitureWoodenBlockWallClock, villagerFurnitureWritingChair, villagerFurnitureWritingDesk, villagerFurnitureWritingPoster}}
	alfonsoGames         = villagersGames{[]VillagerGame{}} // TBD
	alfonsoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	alfonsoInterest      = villagersInterest{villagerInterestPlay}
	alfonsoName          = villagersName{villagerNameAlfonso}
	alfonsoNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	alfonsoMusic         = villagersMusic{}
	alfonsoPersonality   = villagersPersonality{villagerPersonalityLazy}
	alfonsoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAlligator}}
	alfonsoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	alfonsoWallpaper     = villagersWallpaper{villagerWallpaperYellowPlayroomWall}
)
