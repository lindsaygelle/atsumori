package atsumori

import "time"

// Shep is an Animal Crossing villager.
var Shep = villager{
	shepAstrology,
	shepBirthDay,
	shepBirthMonth,
	shepBubbleColor,
	shepCategory,
	shepClothes,
	shepClothesColors,
	shepFlooring,
	shepFurniture,
	shepGames,
	shepGender,
	shepInterest,
	shepName,
	shepNameColor,
	shepMusic,
	shepPersonality,
	shepSpecies,
	shepStyle,
	shepWallpaper}

var (
	shepAstrology     = villagersAstrology{villagerAstrologySagittarius}
	shepBirthDay      = villagersBirthDay{24}
	shepBirthMonth    = villagersBirthMonth{time.November} // 11
	shepBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	shepCategory      = villagersCategory{villagerCategoryA}
	shepClothes       = villagersClothes{villagerClothesDenimJacket} // 3227
	shepClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorBlue}}
	shepFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	shepFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodBurningStove, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureShowerBooth, villagerFurnitureMiniFridge, villagerFurnitureLogBed, villagerFurnitureGasRange, villagerFurniturePotRack, villagerFurnitureDeerDecoration, villagerFurnitureLogDecorativeShelves, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureMagneticKnifeRack, villagerFurnitureRedKilimStyleCarpet, villagerFurnitureRetroStereo}}
	shepGames         = villagersGames{[]VillagerGame{}} // TBD
	shepGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	shepInterest      = villagersInterest{villagerInterestEducation}
	shepName          = villagersName{villagerNameShep}
	shepNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	shepMusic         = villagersMusic{villagerMusicKKSwing}
	shepPersonality   = villagersPersonality{villagerPersonalitySmug}
	shepSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	shepStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	shepWallpaper     = villagersWallpaper{villagerWallpaperCabinWall}
)
