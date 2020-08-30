package atsumori

import "time"

// Ken is an Animal Crossing villager.
var Ken = villager{
	kenAstrology,
	kenBirthDay,
	kenBirthMonth,
	kenBubbleColor,
	kenCategory,
	kenClothes,
	kenClothesColors,
	kenFlooring,
	kenFurniture,
	kenGames,
	kenGender,
	kenInterest,
	kenName,
	kenNameColor,
	kenMusic,
	kenPersonality,
	kenSpecies,
	kenStyle,
	kenWallpaper}

var (
	kenAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	kenBirthDay      = villagersBirthDay{23}
	kenBirthMonth    = villagersBirthMonth{time.December} // 12
	kenBubbleColor   = villagersBubbleColor{villagerBubbleColor6B75CE}
	kenCategory      = villagersCategory{villagerCategoryA}
	kenClothes       = villagersClothes{} // 3460
	kenClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlue}}
	kenFlooring      = villagersFlooring{villagerFlooringTatami}
	kenFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFloorSeat, villagerFurnitureHearth, villagerFurnitureTatamiBed, villagerFurniturePaperLantern, villagerFurnitureBambooSpeaker, villagerFurnitureBambooStool, villagerFurniturePineBonsaiTree, villagerFurnitureKatana, villagerFurnitureWoodenPlankSign, villagerFurnitureHangingScroll}}
	kenGames         = villagersGames{[]VillagerGame{}} // TBD
	kenGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	kenInterest      = villagersInterest{villagerInterestEducation}
	kenName          = villagersName{villagerNameKen}
	kenNameColor     = villagersNameColor{villagerNameColor9AE8DF}
	kenMusic         = villagersMusic{} // K.K. Rally
	kenPersonality   = villagersPersonality{villagerPersonalitySmug}
	kenSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	kenStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	kenWallpaper     = villagersWallpaper{villagerWallpaperStandardTearoomWall}
)
