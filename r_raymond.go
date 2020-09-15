package atsumori

import "time"

// Raymond is an Animal Crossing villager.
var Raymond = villager{
	raymondAstrology,
	raymondBirthDay,
	raymondBirthMonth,
	raymondBubbleColor,
	raymondCategory,
	raymondClothes,
	raymondClothesColors,
	raymondFlooring,
	raymondFurniture,
	raymondGames,
	raymondGender,
	raymondInterest,
	raymondName,
	raymondNameColor,
	raymondMusic,
	raymondPersonality,
	raymondSpecies,
	raymondStyle,
	raymondWallpaper}

var (
	raymondAstrology     = villagersAstrology{villagerAstrologyLibra}
	raymondBirthDay      = villagersBirthDay{1}
	raymondBirthMonth    = villagersBirthMonth{time.October} // 10
	raymondBubbleColor   = villagersBubbleColor{villagerBubbleColorACC8CF}
	raymondCategory      = villagersCategory{villagerCategoryB}
	raymondClothes       = villagersClothes{villagerClothesWaistcoat} // 3339
	raymondClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	raymondFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	raymondFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureOfficeDesk, villagerFurnitureWhiteboard, villagerFurnitureDocumentStack, villagerFurnitureDesktopComputer, villagerFurnitureWaterCooler, villagerFurnitureModernOfficeChair, villagerFurnitureIronWorktable, villagerFurnitureSafe, villagerFurnitureDenChair, villagerFurnitureWallClock, villagerFurnitureDocumentStack, villagerFurnitureFaxMachine, villagerFurnitureFormalPaper, villagerFurnitureDenDesk, villagerFurnitureMug, villagerFurnitureNewtonsCradle}}
	raymondGames         = villagersGames{[]VillagerGame{}} // TBD
	raymondGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	raymondInterest      = villagersInterest{villagerInterestNature}
	raymondName          = villagersName{villagerNameRaymond}
	raymondNameColor     = villagersNameColor{villagerNameColor498992}
	raymondMusic         = villagersMusic{villagerMusicKKCruisin}
	raymondPersonality   = villagersPersonality{villagerPersonalitySmug}
	raymondSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	raymondStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	raymondWallpaper     = villagersWallpaper{villagerWallpaperOfficeWall}
)
