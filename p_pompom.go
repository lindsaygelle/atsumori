package atsumori

import "time"

// Pompom is an Animal Crossing villager.
var Pompom = villager{
	pompomAstrology,
	pompomBirthDay,
	pompomBirthMonth,
	pompomBubbleColor,
	pompomCategory,
	pompomClothes,
	pompomClothesColors,
	pompomFlooring,
	pompomFurniture,
	pompomGames,
	pompomGender,
	pompomInterest,
	pompomName,
	pompomNameColor,
	pompomMusic,
	pompomPersonality,
	pompomSpecies,
	pompomStyle,
	pompomWallpaper}

var (
	pompomAstrology     = villagersAstrology{villagerAstrologyAquarius}
	pompomBirthDay      = villagersBirthDay{11}
	pompomBirthMonth    = villagersBirthMonth{time.February} // 2
	pompomBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	pompomCategory      = villagersCategory{villagerCategoryB}
	pompomClothes       = villagersClothes{} // 4163
	pompomClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorLightBlue}}
	pompomFlooring      = villagersFlooring{villagerFlooringGreenRetroFlooring}
	pompomFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRefrigerator, villagerFurnitureWoodenLowTable, villagerFurnitureDoubleSofa, villagerFurnitureWoodenDoubleBed, villagerFurnitureWoodenFullLengthMirror, villagerFurnitureWoodenChest, villagerFurnitureCuteMusicPlayer, villagerFurnitureCorkboard, villagerFurnitureGasRange, villagerFurnitureSimpleKettle, villagerFurnitureYellowSmallRoundMat, villagerFurnitureAnthuriumPlant}}
	pompomGames         = villagersGames{[]VillagerGame{}} // TBD
	pompomGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pompomInterest      = villagersInterest{villagerInterestMusic}
	pompomName          = villagersName{villagerNamePompom}
	pompomNameColor     = villagersNameColor{villagerNameColor848484}
	pompomMusic         = villagersMusic{villagerMusicKKSoul}
	pompomPersonality   = villagersPersonality{villagerPersonalityPeppy}
	pompomSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	pompomStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleSimple}}
	pompomWallpaper     = villagersWallpaper{villagerWallpaperBeadedCurtainWall}
)
