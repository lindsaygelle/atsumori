package atsumori

import "time"

// Ike is an Animal Crossing villager
var Ike = villager{
	ikeAstrology,
	ikeBirthDay,
	ikeBirthMonth,
	ikeBubbleColor,
	ikeCategory,
	ikeClothes,
	ikeClothesColors,
	ikeFlooring,
	ikeFurniture,
	ikeGames,
	ikeGender,
	ikeInterest,
	ikeName,
	ikeNameColor,
	ikeMusic,
	ikePersonality,
	ikeSpecies,
	ikeStyle,
	ikeWallpaper}

var (
	ikeAstrology     = villagersAstrology{villagerAstrologyTaurus}
	ikeBirthDay      = villagersBirthDay{16}
	ikeBirthMonth    = villagersBirthMonth{time.May} // 5
	ikeBubbleColor   = villagersBubbleColor{villagerBubbleColor4C3317}
	ikeCategory      = villagersCategory{villagerCategoryA}
	ikeClothes       = villagersClothes{villagerClothesCamoBomberStyleJacket} // 4402
	ikeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBlue}}
	ikeFlooring      = villagersFlooring{villagerFlooringSteelFlooring}
	ikeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronShelf, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureDIYWorkbench, villagerFurnitureCampingCot, villagerFurnitureStudioWallSpotlight, villagerFurnitureKeyHolder, villagerFurnitureBreaker, villagerFurnitureToolCart, villagerFurnitureWoodenToolbox, villagerFurnitureIronWorktable, villagerFurnitureToolbox, villagerFurniturePortableRadio, villagerFurnitureWallMountedToolBoard}}
	ikeGames         = villagersGames{[]VillagerGame{}} // TBD
	ikeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	ikeInterest      = villagersInterest{villagerInterestNature}
	ikeName          = villagersName{villagerNameIke}
	ikeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	ikeMusic         = villagersMusic{villagerMusicKKRock}
	ikePersonality   = villagersPersonality{villagerPersonalityCranky}
	ikeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	ikeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	ikeWallpaper     = villagersWallpaper{villagerWallpaperShutterWall}
)
