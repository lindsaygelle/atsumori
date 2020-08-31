package atsumori

import "time"

// Lyman is an Animal Crossing villager.
var Lyman = villager{
	lymanAstrology,
	lymanBirthDay,
	lymanBirthMonth,
	lymanBubbleColor,
	lymanCategory,
	lymanClothes,
	lymanClothesColors,
	lymanFlooring,
	lymanFurniture,
	lymanGames,
	lymanGender,
	lymanInterest,
	lymanName,
	lymanNameColor,
	lymanMusic,
	lymanPersonality,
	lymanSpecies,
	lymanStyle,
	lymanWallpaper}

var (
	lymanAstrology     = villagersAstrology{villagerAstrologyLibra}
	lymanBirthDay      = villagersBirthDay{12}
	lymanBirthMonth    = villagersBirthMonth{time.October} // 10
	lymanBubbleColor   = villagersBubbleColor{villagerBubbleColorA8E989}
	lymanCategory      = villagersCategory{villagerCategoryA}
	lymanClothes       = villagersClothes{villagerClothesVerticalStripesShirt} // 3677
	lymanClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorYellow}}
	lymanFlooring      = villagersFlooring{villagerFlooringLightParquetFlooring}
	lymanFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureToyBox, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureWoodenSimpleBed, villagerFurnitureMonstera, villagerFurnitureWoodenChair, villagerFurnitureWoodenTable, villagerFurnitureBotanicalRug, villagerFurnitureWoodenChest, villagerFurnitureWoodenToolbox, villagerFurnitureThrowbackSkullRadio}}
	lymanGames         = villagersGames{[]VillagerGame{}} // TBD
	lymanGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	lymanInterest      = villagersInterest{villagerInterestPlay}
	lymanName          = villagersName{villagerNameLyman}
	lymanNameColor     = villagersNameColor{villagerNameColor878E86}
	lymanMusic         = villagersMusic{villagerMusicKKSafari}
	lymanPersonality   = villagersPersonality{villagerPersonalityJock}
	lymanSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	lymanStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	lymanWallpaper     = villagersWallpaper{villagerWallpaperJungleWall}
)
