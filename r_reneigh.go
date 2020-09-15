package atsumori

import "time"

// Reneigh is an Animal Crossing villager.
var Reneigh = villager{
	reneighAstrology,
	reneighBirthDay,
	reneighBirthMonth,
	reneighBubbleColor,
	reneighCategory,
	reneighClothes,
	reneighClothesColors,
	reneighFlooring,
	reneighFurniture,
	reneighGames,
	reneighGender,
	reneighInterest,
	reneighName,
	reneighNameColor,
	reneighMusic,
	reneighPersonality,
	reneighSpecies,
	reneighStyle,
	reneighWallpaper}

var (
	reneighAstrology     = villagersAstrology{villagerAstrologyGemini}
	reneighBirthDay      = villagersBirthDay{4}
	reneighBirthMonth    = villagersBirthMonth{time.June} // 6
	reneighBubbleColor   = villagersBubbleColor{villagerBubbleColor76494E}
	reneighCategory      = villagersCategory{villagerCategoryA}
	reneighClothes       = villagersClothes{} // 4616
	reneighClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorPurple}}
	reneighFlooring      = villagersFlooring{villagerFlooringBrownFloralFlooring}
	reneighFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanVanity, villagerFurnitureRattanWardrobe, villagerFurnitureRattanStool, villagerFurnitureMonstera, villagerFurnitureIroningBoard, villagerFurnitureRattanWasteBin, villagerFurnitureRattanLowTable, villagerFurnitureDeluxeWasher, villagerFurnitureMagazine, villagerFurnitureCoconutWallPlanter, villagerFurnitureRattanTowelBasket, villagerFurnitureRattanEndTable, villagerFurnitureRattanTableLamp}}
	reneighGames         = villagersGames{[]VillagerGame{}} // TBD
	reneighGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	reneighInterest      = villagersInterest{villagerInterestPlay}
	reneighName          = villagersName{villagerNameReneigh}
	reneighNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	reneighMusic         = villagersMusic{villagerMusicKKSynth}
	reneighPersonality   = villagersPersonality{villagerPersonalityBigSister}
	reneighSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	reneighStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	reneighWallpaper     = villagersWallpaper{villagerWallpaperBotanicalTileWall}
)
