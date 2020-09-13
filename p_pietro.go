package atsumori

import "time"

// Pietro is an Animal Crossing villager.
var Pietro = villager{
	pietroAstrology,
	pietroBirthDay,
	pietroBirthMonth,
	pietroBubbleColor,
	pietroCategory,
	pietroClothes,
	pietroClothesColors,
	pietroFlooring,
	pietroFurniture,
	pietroGames,
	pietroGender,
	pietroInterest,
	pietroName,
	pietroNameColor,
	pietroMusic,
	pietroPersonality,
	pietroSpecies,
	pietroStyle,
	pietroWallpaper}

var (
	pietroAstrology     = villagersAstrology{villagerAstrologyAries}
	pietroBirthDay      = villagersBirthDay{19}
	pietroBirthMonth    = villagersBirthMonth{time.April} // 4
	pietroBubbleColor   = villagersBubbleColor{villagerBubbleColorFF4040}
	pietroCategory      = villagersCategory{villagerCategoryA}
	pietroClothes       = villagersClothes{} // 3450
	pietroClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorRed}}
	pietroFlooring      = villagersFlooring{villagerFlooringCloudFlooring}
	pietroFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJukebox, villagerFurnitureColorfulWheel, villagerFurnitureColorfulWheel, villagerFurnitureSpringyRideOn, villagerFurnitureSpringyRideOn, villagerFurnitureCottonCandyStall, villagerFurnitureCrescentMoonChair, villagerFurnitureCrescentMoonChair}}
	pietroGames         = villagersGames{[]VillagerGame{}} // TBD
	pietroGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	pietroInterest      = villagersInterest{villagerInterestMusic}
	pietroName          = villagersName{villagerNamePietro}
	pietroNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	pietroMusic         = villagersMusic{villagerMusicKKParade}
	pietroPersonality   = villagersPersonality{villagerPersonalitySmug}
	pietroSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	pietroStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleSimple}}
	pietroWallpaper     = villagersWallpaper{villagerWallpaperSkyWall}
)
