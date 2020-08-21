package atsumori

import "time"

// Antonio is an Animal Crossing villager
var Antonio = villager{
	antonioAstrology,
	antonioBirthDay,
	antonioBirthMonth,
	antonioBubbleColor,
	antonioCategory,
	antonioClothes,
	antonioClothesColors,
	antonioFlooring,
	antonioFurniture,
	antonioGames,
	antonioGender,
	antonioInterest,
	antonioName,
	antonioNameColor,
	antonioMusic,
	antonioPersonality,
	antonioSpecies,
	antonioStyle,
	antonioWallpaper}

var (
	antonioAstrology     = villagersAstrology{villagerAstrologyLibra}
	antonioBirthDay      = villagersBirthDay{20}
	antonioBirthMonth    = villagersBirthMonth{time.October} // 10
	antonioBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	antonioCategory      = villagersCategory{villagerCategoryB}
	antonioClothes       = villagersClothes{villagerClothesBoneTee} // 8200
	antonioClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorBlue}}
	antonioFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	antonioFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenEndTable, villagerFurnitureStackOfBooks, villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenChair, villagerFurnitureMiniDIYWorkbench, villagerFurnitureWoodenTable, villagerFurnitureWoodenChest, villagerFurnitureUnfinishedPuzzle, villagerFurnitureThrowbackSkullRadio}}
	antonioGames         = villagersGames{[]VillagerGame{}} // TBD
	antonioGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	antonioInterest      = villagersInterest{villagerInterestFitness}
	antonioName          = villagersName{villagerNameAntonio}
	antonioNameColor     = villagersNameColor{villagerNameColor879B96}
	antonioMusic         = villagersMusic{villagerMusicKKRagtime}
	antonioPersonality   = villagersPersonality{villagerPersonalityJock}
	antonioSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesAnteater}}
	antonioStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	antonioWallpaper     = villagersWallpaper{villagerWallpaperStarryWall}
)
