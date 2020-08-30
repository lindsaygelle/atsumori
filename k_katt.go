package atsumori

import "time"

// Katt is an Animal Crossing villager.
var Katt = villager{
	kattAstrology,
	kattBirthDay,
	kattBirthMonth,
	kattBubbleColor,
	kattCategory,
	kattClothes,
	kattClothesColors,
	kattFlooring,
	kattFurniture,
	kattGames,
	kattGender,
	kattInterest,
	kattName,
	kattNameColor,
	kattMusic,
	kattPersonality,
	kattSpecies,
	kattStyle,
	kattWallpaper}

var (
	kattAstrology     = villagersAstrology{villagerAstrologyTaurus}
	kattBirthDay      = villagersBirthDay{27}
	kattBirthMonth    = villagersBirthMonth{time.April} // 4
	kattBubbleColor   = villagersBubbleColor{villagerBubbleColorA87850}
	kattCategory      = villagersCategory{villagerCategoryA}
	kattClothes       = villagersClothes{villagerClothesOldSchoolJacket} // 5292
	kattClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlack}}
	kattFlooring      = villagersFlooring{villagerFlooringSteelFlooring}
	kattFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAmp, villagerFurnitureTrashBags, villagerFurnitureCardboardBox, villagerFurnitureCardboardBox, villagerFurnitureVentilationFan, villagerFurnitureElectricGuitar, villagerFurnitureDrumSet, villagerFurnitureDIYWorkbench, villagerFurnitureCardboardBox, villagerFurnitureStackedMagazines, villagerFurnitureIronWallLamp, villagerFurnitureEffectsRack, villagerFurniturePortableRecordPlayer, villagerFurnitureDinerSofa}}
	kattGames         = villagersGames{[]VillagerGame{}} // TBD
	kattGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	kattInterest      = villagersInterest{villagerInterestMusic}
	kattName          = villagersName{villagerNameKatt}
	kattNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	kattMusic         = villagersMusic{villagerMusicKKDisco}
	kattPersonality   = villagersPersonality{villagerPersonalityBigSister}
	kattSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	kattStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleCool}}
	kattWallpaper     = villagersWallpaper{villagerWallpaperWhitePerforatedBoardWall}
)
