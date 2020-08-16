package atsumori

import "time"

// Blanche is an Animal Crossing villager
var Blanche = villager{
	blancheAstrology,
	blancheBirthDay,
	blancheBirthMonth,
	blancheBubbleColor,
	blancheCategory,
	blancheClothes,
	blancheClothesColors,
	blancheFlooring,
	blancheFurniture,
	blancheGames,
	blancheGender,
	blancheInterest,
	blancheName,
	blancheNameColor,
	blancheMusic,
	blanchePersonality,
	blancheSpecies,
	blancheStyle,
	blancheWallpaper}

var (
	blancheAstrology     = villagersAstrology{villagerAstrologySagittarius}
	blancheBirthDay      = villagersBirthDay{21}
	blancheBirthMonth    = villagersBirthMonth{time.December} // 12
	blancheBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	blancheCategory      = villagersCategory{villagerCategoryA}
	blancheClothes       = villagersClothes{} // 6025
	blancheClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorBrown}}
	blancheFlooring      = villagersFlooring{villagerFlooringTatamiFlooring}
	blancheFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFuton, villagerFurnitureLowScreen, villagerFurniturePaperLantern, villagerFurnitureZenCushion, villagerFurnitureElaborateKimonoStand, villagerFurnitureBambooShelf, villagerFurnitureImperialLowTable, villagerFurnitureTraditionalTeaSet, villagerFurnitureImperialDecorativeShelves}}
	blancheGames         = villagersGames{[]VillagerGame{}} // TBD
	blancheGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	blancheInterest      = villagersInterest{villagerInterestNature}
	blancheName          = villagersName{villagerNameBlanche}
	blancheNameColor     = villagersNameColor{villagerNameColor848484}
	blancheMusic         = villagersMusic{} // K.K. Jongara
	blanchePersonality   = villagersPersonality{villagerPersonalitySnooty}
	blancheSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	blancheStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	blancheWallpaper     = villagersWallpaper{villagerWallpaperStandardTearoomWall}
)
