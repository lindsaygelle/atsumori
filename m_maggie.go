package atsumori

import "time"

// Maggie is an Animal Crossing villager.
var Maggie = villager{
	maggieAstrology,
	maggieBirthDay,
	maggieBirthMonth,
	maggieBubbleColor,
	maggieCategory,
	maggieClothes,
	maggieClothesColors,
	maggieFlooring,
	maggieFurniture,
	maggieGames,
	maggieGender,
	maggieInterest,
	maggieName,
	maggieNameColor,
	maggieMusic,
	maggiePersonality,
	maggieSpecies,
	maggieStyle,
	maggieWallpaper}

var (
	maggieAstrology     = villagersAstrology{villagerAstrologyVirgo}
	maggieBirthDay      = villagersBirthDay{3}
	maggieBirthMonth    = villagersBirthMonth{time.September} // 9
	maggieBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	maggieCategory      = villagersCategory{villagerCategoryA}
	maggieClothes       = villagersClothes{} // 8886
	maggieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorYellow}}
	maggieFlooring      = villagersFlooring{villagerFlooringFlagstoneFlooring}
	maggieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGardenWagon, villagerFurnitureGardenWagon, villagerFurnitureGardenWagon, villagerFurnitureIronGardenChair, villagerFurnitureIronGardenTable, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureCuteMusicPlayer, villagerFurnitureIronGardenBench, villagerFurnitureGardenLantern}}
	maggieGames         = villagersGames{[]VillagerGame{}} // TBD
	maggieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	maggieInterest      = villagersInterest{villagerInterestNature}
	maggieName          = villagersName{villagerNameMaggie}
	maggieNameColor     = villagersNameColor{villagerNameColor9B553A}
	maggieMusic         = villagersMusic{villagerMusicDrivin}
	maggiePersonality   = villagersPersonality{villagerPersonalityNormal}
	maggieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	maggieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	maggieWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
