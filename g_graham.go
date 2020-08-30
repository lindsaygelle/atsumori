package atsumori

import "time"

// Graham is an Animal Crossing villager.
var Graham = villager{
	grahamAstrology,
	grahamBirthDay,
	grahamBirthMonth,
	grahamBubbleColor,
	grahamCategory,
	grahamClothes,
	grahamClothesColors,
	grahamFlooring,
	grahamFurniture,
	grahamGames,
	grahamGender,
	grahamInterest,
	grahamName,
	grahamNameColor,
	grahamMusic,
	grahamPersonality,
	grahamSpecies,
	grahamStyle,
	grahamWallpaper}

var (
	grahamAstrology     = villagersAstrology{villagerAstrologyGemini}
	grahamBirthDay      = villagersBirthDay{20}
	grahamBirthMonth    = villagersBirthMonth{time.June} // 6
	grahamBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	grahamCategory      = villagersCategory{villagerCategoryA}
	grahamClothes       = villagersClothes{villagerClothesMadrasPlaidShirt} // 4615
	grahamClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorOrange}}
	grahamFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	grahamFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHighEndStereo, villagerFurnitureBreaker, villagerFurnitureServer, villagerFurnitureAirConditioner, villagerFurnitureServer, villagerFurnitureWoodenBookshelf, villagerFurnitureStackOfBooks, villagerFurnitureCardboardSofa, villagerFurnitureCardboardBed, villagerFurnitureToolCart, villagerFurnitureChrissysPoster, villagerFurnitureCardboardTable, villagerFurnitureBubblegumKK, villagerFurnitureFrancinesPoster, villagerFurnitureLaptop, villagerFurnitureDesktopComputer}}
	grahamGames         = villagersGames{[]VillagerGame{}} // TBD
	grahamGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	grahamInterest      = villagersInterest{villagerInterestEducation}
	grahamName          = villagersName{villagerNameGraham}
	grahamNameColor     = villagersNameColor{villagerNameColor9B553A}
	grahamMusic         = villagersMusic{villagerMusicBubblegumKK}
	grahamPersonality   = villagersPersonality{villagerPersonalitySmug}
	grahamSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHamster}}
	grahamStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	grahamWallpaper     = villagersWallpaper{villagerWallpaperServerRoomWall}
)
