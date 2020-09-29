package atsumori

import "time"

// Tabby is an Animal Crossing villager.
var Tabby = villager{
	tabbyAstrology,
	tabbyBirthDay,
	tabbyBirthMonth,
	tabbyBubbleColor,
	tabbyCategory,
	tabbyClothes,
	tabbyClothesColors,
	tabbyFlooring,
	tabbyFurniture,
	tabbyGames,
	tabbyGender,
	tabbyInterest,
	tabbyName,
	tabbyNameColor,
	tabbyMusic,
	tabbyPersonality,
	tabbySpecies,
	tabbyStyle,
	tabbyWallpaper}

var (
	tabbyAstrology     = villagersAstrology{villagerAstrologyLeo}
	tabbyBirthDay      = villagersBirthDay{13}
	tabbyBirthMonth    = villagersBirthMonth{time.August} // 8
	tabbyBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	tabbyCategory      = villagersCategory{villagerCategoryB}
	tabbyClothes       = villagersClothes{} // 4559
	tabbyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	tabbyFlooring      = villagersFlooring{villagerFlooringTigerPrintFlooring}
	tabbyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFuton, villagerFurnitureStackOfBooks, villagerFurnitureWoodenBookshelf, villagerFurnitureRefrigerator, villagerFurnitureCushion, villagerFurnitureLCDTV50In, villagerFurnitureCardboardBox, villagerFurnitureCuteMusicPlayer, villagerFurnitureToilet, villagerFurnitureCardboardBox, villagerFurnitureIronwoodLowTable, villagerFurnitureMagazineRack, villagerFurnitureGasRange, villagerFurnitureBookStands, villagerFurnitureLaptop, villagerFurnitureSimpleKettle}}
	tabbyGames         = villagersGames{[]VillagerGame{}} // TBD
	tabbyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	tabbyInterest      = villagersInterest{villagerInterestMusic}
	tabbyName          = villagersName{villagerNameTabby}
	tabbyNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	tabbyMusic         = villagersMusic{villagerMusicDJKK}
	tabbyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	tabbySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	tabbyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	tabbyWallpaper     = villagersWallpaper{villagerWallpaperMangaLibraryWall}
)
