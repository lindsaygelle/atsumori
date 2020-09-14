package atsumori

import "time"

// Punchy is an Animal Crossing villager.
var Punchy = villager{
	punchyAstrology,
	punchyBirthDay,
	punchyBirthMonth,
	punchyBubbleColor,
	punchyCategory,
	punchyClothes,
	punchyClothesColors,
	punchyFlooring,
	punchyFurniture,
	punchyGames,
	punchyGender,
	punchyInterest,
	punchyName,
	punchyNameColor,
	punchyMusic,
	punchyPersonality,
	punchySpecies,
	punchyStyle,
	punchyWallpaper}

var (
	punchyAstrology     = villagersAstrology{villagerAstrologyAries}
	punchyBirthDay      = villagersBirthDay{11}
	punchyBirthMonth    = villagersBirthMonth{time.April} // 4
	punchyBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	punchyCategory      = villagersCategory{villagerCategoryB}
	punchyClothes       = villagersClothes{villagerClothesMadrasPlaidShirt} // 8967
	punchyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	punchyFlooring      = villagersFlooring{villagerFlooringJointedMatFlooring}
	punchyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenLowTable, villagerFurnitureWoodenSimpleBed, villagerFurnitureCuteDIYTable, villagerFurnitureWoodenMiniTable, villagerFurnitureGasRange, villagerFurnitureCuteMusicPlayer, villagerFurnitureRefrigerator, villagerFurnitureWoodenChest, villagerFurnitureVentilationFan, villagerFurnitureCatGrass}}
	punchyGames         = villagersGames{[]VillagerGame{}} // TBD
	punchyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	punchyInterest      = villagersInterest{villagerInterestPlay}
	punchyName          = villagersName{villagerNamePunchy}
	punchyNameColor     = villagersNameColor{villagerNameColor848484}
	punchyMusic         = villagersMusic{villagerMusicForestLife}
	punchyPersonality   = villagersPersonality{villagerPersonalityLazy}
	punchySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	punchyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	punchyWallpaper     = villagersWallpaper{villagerWallpaperBlueSimpleClothWall}
)
