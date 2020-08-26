package atsumori

import "time"

// Jacob is an Animal Crossing villager
var Jacob = villager{
	jacobAstrology,
	jacobBirthDay,
	jacobBirthMonth,
	jacobBubbleColor,
	jacobCategory,
	jacobClothes,
	jacobClothesColors,
	jacobFlooring,
	jacobFurniture,
	jacobGames,
	jacobGender,
	jacobInterest,
	jacobName,
	jacobNameColor,
	jacobMusic,
	jacobPersonality,
	jacobSpecies,
	jacobStyle,
	jacobWallpaper}

var (
	jacobAstrology     = villagersAstrology{villagerAstrologyVirgo}
	jacobBirthDay      = villagersBirthDay{24}
	jacobBirthMonth    = villagersBirthMonth{time.August} // 8
	jacobBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	jacobCategory      = villagersCategory{villagerCategoryA}
	jacobClothes       = villagersClothes{villagerClothesCamoBomberStyleJacket} // 4402
	jacobClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorRed}}
	jacobFlooring      = villagersFlooring{villagerFlooringGarbageHeapFlooring}
	jacobFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureWildLogBench, villagerFurnitureHayBed, villagerFurnitureLogStool, villagerFurniturePortableRadio, villagerFurnitureClothesline, villagerFurnitureLogStakes, villagerFurnitureCardboardBox, villagerFurnitureTrashBags}}
	jacobGames         = villagersGames{[]VillagerGame{}} // TBD
	jacobGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	jacobInterest      = villagersInterest{villagerInterestNature}
	jacobName          = villagersName{villagerNameJacob}
	jacobNameColor     = villagersNameColor{villagerNameColor848484}
	jacobMusic         = villagersMusic{villagerMusicKKSong}
	jacobPersonality   = villagersPersonality{villagerPersonalityLazy}
	jacobSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	jacobStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	jacobWallpaper     = villagersWallpaper{villagerWallpaperGarbageHeapWall}
)
