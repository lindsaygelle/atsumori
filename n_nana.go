package atsumori

import "time"

// Nana is an Animal Crossing villager.
var Nana = villager{
	nanaAstrology,
	nanaBirthDay,
	nanaBirthMonth,
	nanaBubbleColor,
	nanaCategory,
	nanaClothes,
	nanaClothesColors,
	nanaFlooring,
	nanaFurniture,
	nanaGames,
	nanaGender,
	nanaInterest,
	nanaName,
	nanaNameColor,
	nanaMusic,
	nanaPersonality,
	nanaSpecies,
	nanaStyle,
	nanaWallpaper}

var (
	nanaAstrology     = villagersAstrology{villagerAstrologyVirgo}
	nanaBirthDay      = villagersBirthDay{23}
	nanaBirthMonth    = villagersBirthMonth{time.August} // 8
	nanaBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	nanaCategory      = villagersCategory{villagerCategoryB}
	nanaClothes       = villagersClothes{villagerClothesPomPomSweater} // 4766
	nanaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorWhite}}
	nanaFlooring      = villagersFlooring{villagerFlooringCorkFlooring}
	nanaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenSimpleBed, villagerFurnitureSystemKitchen, villagerFurnitureAnthuriumPlant, villagerFurnitureWoodenTable, villagerFurnitureWoodenChest, villagerFurnitureRedDottedRug, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenMiniTable, villagerFurnitureHumidifier, villagerFurnitureWoodenChair, villagerFurnitureAirConditioner}}
	nanaGames         = villagersGames{[]VillagerGame{}} // TBD
	nanaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	nanaInterest      = villagersInterest{villagerInterestEducation}
	nanaName          = villagersName{villagerNameNana}
	nanaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	nanaMusic         = villagersMusic{villagerMusicKKBallad}
	nanaPersonality   = villagersPersonality{villagerPersonalityNormal}
	nanaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMonkey}}
	nanaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	nanaWallpaper     = villagersWallpaper{villagerWallpaperPastelDottedWall}
)
