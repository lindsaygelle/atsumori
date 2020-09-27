package atsumori

import "time"

// Sly is an Animal Crossing villager.
var Sly = villager{
	slyAstrology,
	slyBirthDay,
	slyBirthMonth,
	slyBubbleColor,
	slyCategory,
	slyClothes,
	slyClothesColors,
	slyFlooring,
	slyFurniture,
	slyGames,
	slyGender,
	slyInterest,
	slyName,
	slyNameColor,
	slyMusic,
	slyPersonality,
	slySpecies,
	slyStyle,
	slyWallpaper}

var (
	slyAstrology     = villagersAstrology{villagerAstrologyScorpio}
	slyBirthDay      = villagersBirthDay{15}
	slyBirthMonth    = villagersBirthMonth{time.November} // 11
	slyBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	slyCategory      = villagersCategory{villagerCategoryA}
	slyClothes       = villagersClothes{villagerClothesCamoTee} // 3257
	slyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorGreen}}
	slyFlooring      = villagersFlooring{villagerFlooringJungleFlooring}
	slyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureThrowbackDinoScreen, villagerFurnitureCampingCot, villagerFurnitureSimpleDIYWorkbench, villagerFurnitureLogStool, villagerFurnitureHedgeStandee, villagerFurnitureHedgeStandee, villagerFurnitureLeafCampfire, villagerFurnitureToyCentipede, villagerFurnitureGrassStandee, villagerFurnitureGreenLeafPile, villagerFurnitureGrassStandee, villagerFurnitureTreeStandee, villagerFurnitureTreeStandee}}
	slyGames         = villagersGames{[]VillagerGame{}} // TBD
	slyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	slyInterest      = villagersInterest{villagerInterestPlay}
	slyName          = villagersName{villagerNameSly}
	slyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	slyMusic         = villagersMusic{villagerMusicKKSafari}
	slyPersonality   = villagersPersonality{villagerPersonalityJock}
	slySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAlligator}}
	slyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	slyWallpaper     = villagersWallpaper{villagerWallpaperJungleWall}
)
