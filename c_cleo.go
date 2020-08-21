package atsumori

import "time"

// Cleo is an Animal Crossing villager
var Cleo = villager{
	cleoAstrology,
	cleoBirthDay,
	cleoBirthMonth,
	cleoBubbleColor,
	cleoCategory,
	cleoClothes,
	cleoClothesColors,
	cleoFlooring,
	cleoFurniture,
	cleoGames,
	cleoGender,
	cleoInterest,
	cleoName,
	cleoNameColor,
	cleoMusic,
	cleoPersonality,
	cleoSpecies,
	cleoStyle,
	cleoWallpaper}

var (
	cleoAstrology     = villagersAstrology{villagerAstrologyAquarius}
	cleoBirthDay      = villagersBirthDay{9}
	cleoBirthMonth    = villagersBirthMonth{time.February} // 2
	cleoBubbleColor   = villagersBubbleColor{villagerBubbleColorA06FCE}
	cleoCategory      = villagersCategory{villagerCategoryA}
	cleoClothes       = villagersClothes{} // 4570
	cleoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorWhite}}
	cleoFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	cleoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCardboardBox, villagerFurnitureWoodenBookshelf, villagerFurnitureCardboardBox, villagerFurnitureCassettePlayer, villagerFurnitureModernOfficeChair, villagerFurnitureModernOfficeChair, villagerFurnitureWhiteboard, villagerFurnitureOfficeDesk, villagerFurnitureWallClock, villagerFurnitureWallMountedPhone, villagerFurnitureWallMountedPhone, villagerFurnitureLaptop, villagerFurnitureIronWorktable, villagerFurnitureOfficeDesk, villagerFurnitureDocumentStack, villagerFurnitureDesktopComputer}}
	cleoGames         = villagersGames{[]VillagerGame{}} // TBD
	cleoGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	cleoInterest      = villagersInterest{villagerInterestEducation}
	cleoName          = villagersName{villagerNameCleo}
	cleoNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	cleoMusic         = villagersMusic{villagerMusicKKFusion}
	cleoPersonality   = villagersPersonality{villagerPersonalitySnooty}
	cleoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	cleoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	cleoWallpaper     = villagersWallpaper{villagerWallpaperOfficeWall}
)
