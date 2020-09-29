package atsumori

import "time"

// Sparro is an Animal Crossing villager.
var Sparro = villager{
	sparroAstrology,
	sparroBirthDay,
	sparroBirthMonth,
	sparroBubbleColor,
	sparroCategory,
	sparroClothes,
	sparroClothesColors,
	sparroFlooring,
	sparroFurniture,
	sparroGames,
	sparroGender,
	sparroInterest,
	sparroName,
	sparroNameColor,
	sparroMusic,
	sparroPersonality,
	sparroSpecies,
	sparroStyle,
	sparroWallpaper}

var (
	sparroAstrology     = villagersAstrology{villagerAstrologyScorpio}
	sparroBirthDay      = villagersBirthDay{20}
	sparroBirthMonth    = villagersBirthMonth{time.November} // 11
	sparroBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	sparroCategory      = villagersCategory{villagerCategoryA}
	sparroClothes       = villagersClothes{villagerClothesEarbudsCombo} // 8407
	sparroClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorGray}}
	sparroFlooring      = villagersFlooring{villagerFlooringSandlot}
	sparroFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGarbagePail, villagerFurnitureTricycle, villagerFurnitureTireToy, villagerFurnitureSandbox, villagerFurnitureElephantSlide, villagerFurnitureBirdhouse, villagerFurnitureGardenFaucet, villagerFurnitureTireToy}}
	sparroGames         = villagersGames{[]VillagerGame{}} // TBD
	sparroGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	sparroInterest      = villagersInterest{villagerInterestPlay}
	sparroName          = villagersName{villagerNameSparro}
	sparroNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	sparroMusic         = villagersMusic{villagerMusicILoveYou}
	sparroPersonality   = villagersPersonality{villagerPersonalityJock}
	sparroSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	sparroStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	sparroWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
