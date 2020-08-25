package atsumori

import "time"

// Hamlet is an Animal Crossing villager
var Hamlet = villager{
	hamletAstrology,
	hamletBirthDay,
	hamletBirthMonth,
	hamletBubbleColor,
	hamletCategory,
	hamletClothes,
	hamletClothesColors,
	hamletFlooring,
	hamletFurniture,
	hamletGames,
	hamletGender,
	hamletInterest,
	hamletName,
	hamletNameColor,
	hamletMusic,
	hamletPersonality,
	hamletSpecies,
	hamletStyle,
	hamletWallpaper}

var (
	hamletAstrology     = villagersAstrology{villagerAstrologyGemini}
	hamletBirthDay      = villagersBirthDay{30}
	hamletBirthMonth    = villagersBirthMonth{time.May} // 5
	hamletBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	hamletCategory      = villagersCategory{villagerCategoryA}
	hamletClothes       = villagersClothes{villagerClothesBigStarTee} // 6822
	hamletClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlue}}
	hamletFlooring      = villagersFlooring{villagerFlooringColorfulPuzzleFlooring}
	hamletFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackRaceCarBed, villagerFurnitureWoodenBlockstool, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureWoodenBlockStereo, villagerFurnitureThrowbackMittChair, villagerFurnitureDIYWorkbench, villagerFurnitureWritingPoster, villagerFurnitureThrowbackContainer, villagerFurnitureToyBox, villagerFurnitureThrowbackHatTable, villagerFurnitureWoodenBlockWallClock, villagerFurnitureModelingClay}}
	hamletGames         = villagersGames{[]VillagerGame{}} // TBD
	hamletGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	hamletInterest      = villagersInterest{villagerInterestPlay}
	hamletName          = villagersName{villagerNameHamlet}
	hamletNameColor     = villagersNameColor{villagerNameColor9B553A}
	hamletMusic         = villagersMusic{villagerMusicKKReggae}
	hamletPersonality   = villagersPersonality{villagerPersonalityJock}
	hamletSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHamster}}
	hamletStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	hamletWallpaper     = villagersWallpaper{villagerWallpaperBluePlayroomWall}
)
