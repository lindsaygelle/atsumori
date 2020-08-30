package atsumori

import "time"

// Eugene is an Animal Crossing villager.
var Eugene = villager{
	eugeneAstrology,
	eugeneBirthDay,
	eugeneBirthMonth,
	eugeneBubbleColor,
	eugeneCategory,
	eugeneClothes,
	eugeneClothesColors,
	eugeneFlooring,
	eugeneFurniture,
	eugeneGames,
	eugeneGender,
	eugeneInterest,
	eugeneName,
	eugeneNameColor,
	eugeneMusic,
	eugenePersonality,
	eugeneSpecies,
	eugeneStyle,
	eugeneWallpaper}

var (
	eugeneAstrology     = villagersAstrology{villagerAstrologyScorpio}
	eugeneBirthDay      = villagersBirthDay{26}
	eugeneBirthMonth    = villagersBirthMonth{time.October} // 10
	eugeneBubbleColor   = villagersBubbleColor{villagerBubbleColor4C3317}
	eugeneCategory      = villagersCategory{villagerCategoryA}
	eugeneClothes       = villagersClothes{villagerClothesBikerJacket} // 3198
	eugeneClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	eugeneFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	eugeneFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWaterCooler, villagerFurnitureCardboardBed, villagerFurnitureWhiteboard, villagerFurnitureFoldingChair, villagerFurnitureFoldingChair, villagerFurnitureWallClock, villagerFurnitureLectureHallDesk, villagerFurnitureLaptop, villagerFurnitureCardboardBox, villagerFurnitureDocumentStack, villagerFurnitureLectureHallDesk, villagerFurnitureDocumentStack, villagerFurnitureMug, villagerFurnitureCardboardBox, villagerFurnitureProTapeRecorder, villagerFurnitureFanPalm}}
	eugeneGames         = villagersGames{[]VillagerGame{}} // TBD
	eugeneGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	eugeneInterest      = villagersInterest{villagerInterestMusic}
	eugeneName          = villagersName{villagerNameEugene}
	eugeneNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	eugeneMusic         = villagersMusic{villagerMusicAgentKK}
	eugenePersonality   = villagersPersonality{villagerPersonalitySmug}
	eugeneSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	eugeneStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	eugeneWallpaper     = villagersWallpaper{villagerWallpaperOfficeWall}
)
