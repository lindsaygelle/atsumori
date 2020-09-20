package atsumori

import "time"

// Rocco is an Animal Crossing villager.
var Rocco = villager{
	roccoAstrology,
	roccoBirthDay,
	roccoBirthMonth,
	roccoBubbleColor,
	roccoCategory,
	roccoClothes,
	roccoClothesColors,
	roccoFlooring,
	roccoFurniture,
	roccoGames,
	roccoGender,
	roccoInterest,
	roccoName,
	roccoNameColor,
	roccoMusic,
	roccoPersonality,
	roccoSpecies,
	roccoStyle,
	roccoWallpaper}

var (
	roccoAstrology     = villagersAstrology{villagerAstrologyLeo}
	roccoBirthDay      = villagersBirthDay{18}
	roccoBirthMonth    = villagersBirthMonth{time.August} // 8
	roccoBubbleColor   = villagersBubbleColor{villagerBubbleColor87E0A9}
	roccoCategory      = villagersCategory{villagerCategoryB}
	roccoClothes       = villagersClothes{villagerClothesDangerTank} // 8192
	roccoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorBlack}}
	roccoFlooring      = villagersFlooring{villagerFlooringConstructionSiteFlooring}
	roccoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHandcart, villagerFurnitureConstructionSign, villagerFurnitureUtilityPole, villagerFurnitureUtilityPole, villagerFurnitureOutdoorsyShovel, villagerFurnitureIronFrame, villagerFurnitureIronFrame, villagerFurnitureOilBarrel, villagerFurnitureIronwoodDIYWorkbench, villagerFurniturePortableToilet, villagerFurnitureMetalCan, villagerFurniturePortableRadio, villagerFurnitureCone, villagerFurnitureCone}}
	roccoGames         = villagersGames{[]VillagerGame{}} // TBD
	roccoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	roccoInterest      = villagersInterest{villagerInterestEducation}
	roccoName          = villagersName{villagerNameRocco}
	roccoNameColor     = villagersNameColor{villagerNameColor219479}
	roccoMusic         = villagersMusic{villagerMusicKKLament}
	roccoPersonality   = villagersPersonality{villagerPersonalityCranky}
	roccoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHippo}}
	roccoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	roccoWallpaper     = villagersWallpaper{villagerWallpaperConstructionSiteWall}
)
