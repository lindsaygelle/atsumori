package atsumori

import "time"

// Pashmina is an Animal Crossing villager.
var Pashmina = villager{
	pashminaAstrology,
	pashminaBirthDay,
	pashminaBirthMonth,
	pashminaBubbleColor,
	pashminaCategory,
	pashminaClothes,
	pashminaClothesColors,
	pashminaFlooring,
	pashminaFurniture,
	pashminaGames,
	pashminaGender,
	pashminaInterest,
	pashminaName,
	pashminaNameColor,
	pashminaMusic,
	pashminaPersonality,
	pashminaSpecies,
	pashminaStyle,
	pashminaWallpaper}

var (
	pashminaAstrology     = villagersAstrology{villagerAstrologySagittarius}
	pashminaBirthDay      = villagersBirthDay{26}
	pashminaBirthMonth    = villagersBirthMonth{time.December} // 12
	pashminaBubbleColor   = villagersBubbleColor{villagerBubbleColorE2856B}
	pashminaCategory      = villagersCategory{villagerCategoryA}
	pashminaClothes       = villagersClothes{villagerClothesRainbowSweater} // 4595
	pashminaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorColorful}}
	pashminaFlooring      = villagersFlooring{villagerFlooringCuteWhiteTileFlooring}
	pashminaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJukebox, villagerFurnitureCuteSofa, villagerFurnitureDinerNeonSign, villagerFurnitureDinerCounterTable, villagerFurnitureSystemKitchen, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureDinerMiniTable, villagerFurnitureDinerCounterTable, villagerFurnitureDinnerware, villagerFurnitureLaptop}}
	pashminaGames         = villagersGames{[]VillagerGame{}} // TBD
	pashminaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pashminaInterest      = villagersInterest{villagerInterestMusic}
	pashminaName          = villagersName{villagerNamePashmina}
	pashminaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	pashminaMusic         = villagersMusic{villagerMusicKKDisco}
	pashminaPersonality   = villagersPersonality{villagerPersonalityBigSister}
	pashminaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGoat}}
	pashminaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleElegant}}
	pashminaWallpaper     = villagersWallpaper{villagerWallpaperCuteRedWall}
)
