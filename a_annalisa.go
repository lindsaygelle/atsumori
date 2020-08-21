package atsumori

import (
	"time"
)

// Annalisa is an Animal Crossing villager
var Annalisa = villager{
	annalisaAstrology,
	annalisaBirthDay,
	annalisaBirthMonth,
	annalisaBubbleColor,
	annalisaCategory,
	annalisaClothes,
	annalisaClothesColors,
	annalisaFlooring,
	annalisaFurniture,
	annalisaGames,
	annalisaGender,
	annalisaInterest,
	annalisaName,
	annalisaNameColor,
	annalisaMusic,
	annalisaPersonality,
	annalisaSpecies,
	annalisaStyle,
	annalisaWallpaper}

var (
	annalisaAstrology     = villagersAstrology{villagerAstrologyAquarius}
	annalisaBirthDay      = villagersBirthDay{6}
	annalisaBirthMonth    = villagersBirthMonth{time.February}
	annalisaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	annalisaCategory      = villagersCategory{villagerCategoryA}
	annalisaClothes       = villagersClothes{} // 3572
	annalisaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorPink}}
	annalisaFlooring      = villagersFlooring{villagerFlooringTatami}
	annalisaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFuton, villagerFurniturePaperLantern, villagerFurnitureBambooSphere, villagerFurnitureScreen, villagerFurnitureBambooShelf, villagerFurniturePileOfZenCushions, villagerFurnitureFloorSeat, villagerFurnitureImperialLowTable, villagerFurnitureMossBall, villagerFurnitureElaborateKimonoStand}}
	annalisaGames         = villagersGames{[]VillagerGame{}} // TBD
	annalisaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	annalisaInterest      = villagersInterest{villagerInterestNature}
	annalisaName          = villagersName{villagerNameAnnalisa}
	annalisaNameColor     = villagersNameColor{villagerNameColor848484}
	annalisaMusic         = villagersMusic{villagerMusicKingKK}
	annalisaPersonality   = villagersPersonality{villagerPersonalityNormal}
	annalisaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAnteater}}
	annalisaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	annalisaWallpaper     = villagersWallpaper{villagerWallpaperGoldScreenWall}
)
