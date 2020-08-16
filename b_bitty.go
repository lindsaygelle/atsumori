package atsumori

import "time"

// Bitty is an Animal Crossing villager
var Bitty = villager{
	bittyAstrology,
	bittyBirthDay,
	bittyBirthMonth,
	bittyBubbleColor,
	bittyCategory,
	bittyClothes,
	bittyClothesColors,
	bittyFlooring,
	bittyFurniture,
	bittyGames,
	bittyGender,
	bittyInterest,
	bittyName,
	bittyNameColor,
	bittyMusic,
	bittyPersonality,
	bittySpecies,
	bittyStyle,
	bittyWallpaper}

var (
	bittyAstrology     = villagersAstrology{villagerAstrologyLibra}
	bittyBirthDay      = villagersBirthDay{6}
	bittyBirthMonth    = villagersBirthMonth{time.October} // 10
	bittyBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	bittyCategory      = villagersCategory{villagerCategoryA}
	bittyClothes       = villagersClothes{} // 4403
	bittyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorOrange}}
	bittyFlooring      = villagersFlooring{villagerFlooringLobbyFlooring}
	bittyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureOldSewingMachine, villagerFurnitureChangingRoom, villagerFurnitureIroningBoard, villagerFurnitureIronWorktable, villagerFurnitureSewingMachine, villagerFurniturePortableRecordPlayer, villagerFurnitureIronwoodCart, villagerFurnitureAccessoriesStand}}
	bittyGames         = villagersGames{[]VillagerGame{}} // TBD
	bittyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	bittyInterest      = villagersInterest{villagerInterestEducation}
	bittyName          = villagersName{villagerNameBitty}
	bittyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bittyMusic         = villagersMusic{} // K.K. Disco
	bittyPersonality   = villagersPersonality{villagerPersonalitySnooty}
	bittySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHippo}}
	bittyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	bittyWallpaper     = villagersWallpaper{villagerWallpaperGrayMoldedPanelWall}
)
