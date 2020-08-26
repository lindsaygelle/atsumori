package atsumori

import "time"

// Hopkins is an Animal Crossing villager
var Hopkins = villager{
	hopkinsAstrology,
	hopkinsBirthDay,
	hopkinsBirthMonth,
	hopkinsBubbleColor,
	hopkinsCategory,
	hopkinsClothes,
	hopkinsClothesColors,
	hopkinsFlooring,
	hopkinsFurniture,
	hopkinsGames,
	hopkinsGender,
	hopkinsInterest,
	hopkinsName,
	hopkinsNameColor,
	hopkinsMusic,
	hopkinsPersonality,
	hopkinsSpecies,
	hopkinsStyle,
	hopkinsWallpaper}

var (
	hopkinsAstrology     = villagersAstrology{villagerAstrologyPisces}
	hopkinsBirthDay      = villagersBirthDay{11}
	hopkinsBirthMonth    = villagersBirthMonth{time.March} // 3
	hopkinsBubbleColor   = villagersBubbleColor{villagerBubbleColor55CFFF}
	hopkinsCategory      = villagersCategory{villagerCategoryA}
	hopkinsClothes       = villagersClothes{villagerClothesStripedShirt} // 2655
	hopkinsClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorYellow}}
	hopkinsFlooring      = villagersFlooring{villagerFlooringCoolVinylFlooring}
	hopkinsFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureArcadeCombatGame, villagerFurnitureArcadeCombatGame, villagerFurnitureArcadeFightingGame, villagerFurnitureWallMountedTV50In, villagerFurnitureArcadeFightingGame, villagerFurnitureDrinkMachine, villagerFurnitureArcadeMahjongGame, villagerFurnitureArcadeMahjongGame, villagerFurnitureSnackMachine, villagerFurnitureSurveillanceCamera, villagerFurnitureVentilationFan, villagerFurnitureExitSign, villagerFurnitureArcadeSeat}}
	hopkinsGames         = villagersGames{[]VillagerGame{}} // TBD
	hopkinsGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	hopkinsInterest      = villagersInterest{villagerInterestNature}
	hopkinsName          = villagersName{villagerNameHopkins}
	hopkinsNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	hopkinsMusic         = villagersMusic{villagerMusicPondering}
	hopkinsPersonality   = villagersPersonality{villagerPersonalityLazy}
	hopkinsSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	hopkinsStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	hopkinsWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
