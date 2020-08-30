package atsumori

import "time"

// Chevre is an Animal Crossing villager.
var Chevre = villager{
	chevreAstrology,
	chevreBirthDay,
	chevreBirthMonth,
	chevreBubbleColor,
	chevreCategory,
	chevreClothes,
	chevreClothesColors,
	chevreFlooring,
	chevreFurniture,
	chevreGames,
	chevreGender,
	chevreInterest,
	chevreName,
	chevreNameColor,
	chevreMusic,
	chevrePersonality,
	chevreSpecies,
	chevreStyle,
	chevreWallpaper}

var (
	chevreAstrology     = villagersAstrology{villagerAstrologyPisces}
	chevreBirthDay      = villagersBirthDay{6}
	chevreBirthMonth    = villagersBirthMonth{time.March} // 3
	chevreBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	chevreCategory      = villagersCategory{villagerCategoryB}
	chevreClothes       = villagersClothes{} // 4368
	chevreClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorPink}}
	chevreFlooring      = villagersFlooring{villagerFlooringBlackIronParquetFlooring}
	chevreFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanVanity, villagerFurnitureLilyRecordPlayer, villagerFurnitureDoubleSofa, villagerFurnitureGrandPiano, villagerFurniturePianoBench, villagerFurnitureKitchenIsland, villagerFurnitureRattanLowTable, villagerFurnitureTissueBox, villagerFurnitureNansPhoto}}
	chevreGames         = villagersGames{[]VillagerGame{}} // TBD
	chevreGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	chevreInterest      = villagersInterest{villagerInterestEducation}
	chevreName          = villagersName{villagerNameChevre}
	chevreNameColor     = villagersNameColor{villagerNameColor848484}
	chevreMusic         = villagersMusic{villagerMusicKKChorale}
	chevrePersonality   = villagersPersonality{villagerPersonalityNormal}
	chevreSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGoat}}
	chevreStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	chevreWallpaper     = villagersWallpaper{villagerWallpaperWhiteBotanicalTileWall}
)
