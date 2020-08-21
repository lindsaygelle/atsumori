package atsumori

import "time"

// Boots is an Animal Crossing villager
var Boots = villager{
	bootsAstrology,
	bootsBirthDay,
	bootsBirthMonth,
	bootsBubbleColor,
	bootsCategory,
	bootsClothes,
	bootsClothesColors,
	bootsFlooring,
	bootsFurniture,
	bootsGames,
	bootsGender,
	bootsInterest,
	bootsName,
	bootsNameColor,
	bootsMusic,
	bootsPersonality,
	bootsSpecies,
	bootsStyle,
	bootsWallpaper}

var (
	bootsAstrology     = villagersAstrology{villagerAstrologyLeo}
	bootsBirthDay      = villagersBirthDay{7}
	bootsBirthMonth    = villagersBirthMonth{time.August} // 8
	bootsBubbleColor   = villagersBubbleColor{villagerBubbleColor0CA54A}
	bootsCategory      = villagersCategory{villagerCategoryA}
	bootsClothes       = villagersClothes{} // 3450
	bootsClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorPurple}}
	bootsFlooring      = villagersFlooring{villagerFlooringDirtFlooring}
	bootsFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHayBed, villagerFurnitureSimpleDIYWorkbench, villagerFurniturePondStone, villagerFurnitureSimpleWell, villagerFurnitureScarecrow, villagerFurnitureLogStool, villagerFurnitureDestinationsSignpost, villagerFurnitureBambooBench, villagerFurnitureHandcart, villagerFurniturePortableRadio, villagerFurnitureBambooLunchBox}}
	bootsGames         = villagersGames{[]VillagerGame{}} // TBD
	bootsGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	bootsInterest      = villagersInterest{villagerInterestPlay}
	bootsName          = villagersName{villagerNameBoots}
	bootsNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bootsMusic         = villagersMusic{villagerMusicWandering}
	bootsPersonality   = villagersPersonality{villagerPersonalityJock}
	bootsSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAlligator}}
	bootsStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleActive}}
	bootsWallpaper     = villagersWallpaper{villagerWallpaperRicePaddyWall}
)
