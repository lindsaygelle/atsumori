package atsumori

import "time"

// Rizzo is an Animal Crossing villager.
var Rizzo = villager{
	rizzoAstrology,
	rizzoBirthDay,
	rizzoBirthMonth,
	rizzoBubbleColor,
	rizzoCategory,
	rizzoClothes,
	rizzoClothesColors,
	rizzoFlooring,
	rizzoFurniture,
	rizzoGames,
	rizzoGender,
	rizzoInterest,
	rizzoName,
	rizzoNameColor,
	rizzoMusic,
	rizzoPersonality,
	rizzoSpecies,
	rizzoStyle,
	rizzoWallpaper}

var (
	rizzoAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	rizzoBirthDay      = villagersBirthDay{17}
	rizzoBirthMonth    = villagersBirthMonth{time.January} // 1
	rizzoBubbleColor   = villagersBubbleColor{villagerBubbleColor8BCDEA}
	rizzoCategory      = villagersCategory{villagerCategoryB}
	rizzoClothes       = villagersClothes{} // 3460
	rizzoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	rizzoFlooring      = villagersFlooring{villagerFlooringDirtFlooring}
	rizzoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBonsaiShelf, villagerFurnitureBambooSpeaker, villagerFurnitureBambooBench, villagerFurniturePlainWoodenShopSign, villagerFurnitureBlossomViewingLantern, villagerFurnitureBlossomViewingLantern, villagerFurnitureStall, villagerFurnitureBambooLunchBox, villagerFurnitureBambooLunchBox}}
	rizzoGames         = villagersGames{[]VillagerGame{}} // TBD
	rizzoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rizzoInterest      = villagersInterest{villagerInterestEducation}
	rizzoName          = villagersName{villagerNameRizzo}
	rizzoNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	rizzoMusic         = villagersMusic{villagerMusicKKFolk}
	rizzoPersonality   = villagersPersonality{villagerPersonalityCranky}
	rizzoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	rizzoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	rizzoWallpaper     = villagersWallpaper{villagerWallpaperMortarWall}
)
