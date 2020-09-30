package atsumori

import "time"

// Tad is an Animal Crossing villager.
var Tad = villager{
	tadAstrology,
	tadBirthDay,
	tadBirthMonth,
	tadBubbleColor,
	tadCategory,
	tadClothes,
	tadClothesColors,
	tadFlooring,
	tadFurniture,
	tadGames,
	tadGender,
	tadInterest,
	tadName,
	tadNameColor,
	tadMusic,
	tadPersonality,
	tadSpecies,
	tadStyle,
	tadWallpaper}

var (
	tadAstrology     = villagersAstrology{villagerAstrologyLeo}
	tadBirthDay      = villagersBirthDay{3}
	tadBirthMonth    = villagersBirthMonth{time.August} // 8
	tadBubbleColor   = villagersBubbleColor{villagerBubbleColor76B788}
	tadCategory      = villagersCategory{villagerCategoryA}
	tadClothes       = villagersClothes{villagerClothesOneBallTee} // 3320
	tadClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorColorful}}
	tadFlooring      = villagersFlooring{villagerFlooringFieldFlooring}
	tadFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureOutdoorPicnicSet, villagerFurnitureTapeDeck, villagerFurnitureShantyMat, villagerFurnitureScarecrow, villagerFurniturePondStone, villagerFurnitureBambooBench, villagerFurnitureBambooLunchBox}}
	tadGames         = villagersGames{[]VillagerGame{}} // TBD
	tadGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	tadInterest      = villagersInterest{villagerInterestPlay}
	tadName          = villagersName{villagerNameTad}
	tadNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	tadMusic         = villagersMusic{villagerMusicWandering}
	tadPersonality   = villagersPersonality{villagerPersonalityJock}
	tadSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	tadStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	tadWallpaper     = villagersWallpaper{villagerWallpaperRicePaddyWall}
)
