package atsumori

import "time"

// Chrissy is an Animal Crossing villager
var Chrissy = villager{
	chrissyAstrology,
	chrissyBirthDay,
	chrissyBirthMonth,
	chrissyBubbleColor,
	chrissyCategory,
	chrissyClothes,
	chrissyClothesColors,
	chrissyFlooring,
	chrissyFurniture,
	chrissyGames,
	chrissyGender,
	chrissyInterest,
	chrissyName,
	chrissyNameColor,
	chrissyMusic,
	chrissyPersonality,
	chrissySpecies,
	chrissyStyle,
	chrissyWallpaper}

var (
	chrissyAstrology     = villagersAstrology{villagerAstrologyVirgo}
	chrissyBirthDay      = villagersBirthDay{28}
	chrissyBirthMonth    = villagersBirthMonth{time.August} // 8
	chrissyBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	chrissyCategory      = villagersCategory{villagerCategoryB}
	chrissyClothes       = villagersClothes{} // 8295
	chrissyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorWhite}}
	chrissyFlooring      = villagersFlooring{villagerFlooringCuteRedTileFlooring}
	chrissyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteChair, villagerFurnitureCuteBed, villagerFurnitureCuteFloorLamp, villagerFurnitureCuteWallMountedClock, villagerFurnitureCuteTeaTable, villagerFurnitureCuteWardrobe, villagerFurnitureCuteMusicPlayer, villagerFurnitureFrancinesPhoto, villagerFurnitureCuteVanity, villagerFurnitureCuteSofa, villagerFurnitureCuteDIYTable}}
	chrissyGames         = villagersGames{[]VillagerGame{}} // TBD
	chrissyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	chrissyInterest      = villagersInterest{villagerInterestFashion}
	chrissyName          = villagersName{villagerNameChrissy}
	chrissyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	chrissyMusic         = villagersMusic{} // Bubblegum K.K.
	chrissyPersonality   = villagersPersonality{villagerPersonalityPeppy}
	chrissySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	chrissyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	chrissyWallpaper     = villagersWallpaper{villagerWallpaperRedDottedWall}
)
