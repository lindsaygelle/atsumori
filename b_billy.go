package atsumori

import "time"

// Billy is an Animal Crossing villager
var Billy = villager{
	billyAstrology,
	billyBirthDay,
	billyBirthMonth,
	billyBubbleColor,
	billyCategory,
	billyClothes,
	billyClothesColors,
	billyFlooring,
	billyFurniture,
	billyGames,
	billyGender,
	billyInterest,
	billyName,
	billyNameColor,
	billyMusic,
	billyPersonality,
	billySpecies,
	billyStyle,
	billyWallpaper}

var (
	billyAstrology     = villagersAstrology{villagerAstrologyAries}
	billyBirthDay      = villagersBirthDay{25}
	billyBirthMonth    = villagersBirthMonth{time.March} // 3
	billyBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	billyCategory      = villagersCategory{villagerCategoryA}
	billyClothes       = villagersClothes{} // 3164
	billyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorPurple}}
	billyFlooring      = villagersFlooring{villagerFlooringGarbageHeapFlooring}
	billyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCardboardSofa, villagerFurnitureCardboardChair, villagerFurnitureCardboardChair, villagerFurnitureScatteredPapers, villagerFurnitureStackOfBooks, villagerFurnitureCardboardBed, villagerFurnitureCardboardBox, villagerFurnitureCardboardBox, villagerFurnitureStackedMagazines, villagerFurnitureCardboardBox, villagerFurnitureCardboardTable, villagerFurnitureDocumentStack, villagerFurnitureTissueBox}}
	billyGames         = villagersGames{[]VillagerGame{}} // TBD
	billyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	billyInterest      = villagersInterest{villagerInterestPlay}
	billyName          = villagersName{villagerNameBilly}
	billyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	billyMusic         = villagersMusic{villagerMusicKKSong}
	billyPersonality   = villagersPersonality{villagerPersonalityJock}
	billySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGoat}}
	billyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	billyWallpaper     = villagersWallpaper{villagerWallpaperMangaLibraryWall}
)
