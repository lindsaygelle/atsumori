package atsumori

import "time"

// Deli is an Animal Crossing villager.
var Deli = villager{
	deliAstrology,
	deliBirthDay,
	deliBirthMonth,
	deliBubbleColor,
	deliCategory,
	deliClothes,
	deliClothesColors,
	deliFlooring,
	deliFurniture,
	deliGames,
	deliGender,
	deliInterest,
	deliName,
	deliNameColor,
	deliMusic,
	deliPersonality,
	deliSpecies,
	deliStyle,
	deliWallpaper}

var (
	deliAstrology     = villagersAstrology{villagerAstrologyGemini}
	deliBirthDay      = villagersBirthDay{24}
	deliBirthMonth    = villagersBirthMonth{time.May} // 5
	deliBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	deliCategory      = villagersCategory{villagerCategoryA}
	deliClothes       = villagersClothes{villagerClothesArgyleVest} // 3262
	deliClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBrown}}
	deliFlooring      = villagersFlooring{villagerFlooringSimpleBlueFlooring}
	deliFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBookshelf, villagerFurnitureWritingPoster, villagerFurnitureAirConditioner, villagerFurnitureWritingDesk, villagerFurnitureWritingChair, villagerFurnitureDoubleSofa, villagerFurnitureWoodenWasteBin, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenChest, villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenLowTable, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureGlobe, villagerFurnitureCuteMusicPlayer, villagerFurnitureLaptop, villagerFurnitureBlueDottedRug}}
	deliGames         = villagersGames{[]VillagerGame{}} // TBD
	deliGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	deliInterest      = villagersInterest{villagerInterestNature}
	deliName          = villagersName{villagerNameDeli}
	deliNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	deliMusic         = villagersMusic{villagerMusicKKEtude}
	deliPersonality   = villagersPersonality{villagerPersonalityLazy}
	deliSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMonkey}}
	deliStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	deliWallpaper     = villagersWallpaper{villagerWallpaperBlueSimpleClothWall}
)
