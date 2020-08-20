package atsumori

import "time"

// Celia is an Animal Crossing villager
var Celia = villager{
	celiaAstrology,
	celiaBirthDay,
	celiaBirthMonth,
	celiaBubbleColor,
	celiaCategory,
	celiaClothes,
	celiaClothesColors,
	celiaFlooring,
	celiaFurniture,
	celiaGames,
	celiaGender,
	celiaInterest,
	celiaName,
	celiaNameColor,
	celiaMusic,
	celiaPersonality,
	celiaSpecies,
	celiaStyle,
	celiaWallpaper}

var (
	celiaAstrology     = villagersAstrology{villagerAstrologyAries}
	celiaBirthDay      = villagersBirthDay{25}
	celiaBirthMonth    = villagersBirthMonth{time.March} // 3
	celiaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	celiaCategory      = villagersCategory{villagerCategoryA}
	celiaClothes       = villagersClothes{} // 3168
	celiaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorGreen}}
	celiaFlooring      = villagersFlooring{villagerFlooringLightHerringboneFlooring}
	celiaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanVanity, villagerFurnitureRattanStool, villagerFurnitureRattanWardrobe, villagerFurnitureFanPalm, villagerFurnitureLilyRecordPlayer, villagerFurnitureRefrigerator, villagerFurnitureRattanLowTable, villagerFurnitureTeaSet, villagerFurnitureRattanEndTable, villagerFurnitureCoconutWallPlanter, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureIronwoodKitchenette}}
	celiaGames         = villagersGames{[]VillagerGame{}} // TBD
	celiaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	celiaInterest      = villagersInterest{villagerInterestNature}
	celiaName          = villagersName{villagerNameCelia}
	celiaNameColor     = villagersNameColor{villagerNameColor848484}
	celiaMusic         = villagersMusic{} // K.K. Soul
	celiaPersonality   = villagersPersonality{villagerPersonalityNormal}
	celiaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	celiaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCute}}
	celiaWallpaper     = villagersWallpaper{villagerWallpaperWhiteBotanicalTileWall}
)
