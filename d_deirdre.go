package atsumori

import "time"

// Deirdre is an Animal Crossing villager
var Deirdre = villager{
	deirdreAstrology,
	deirdreBirthDay,
	deirdreBirthMonth,
	deirdreBubbleColor,
	deirdreCategory,
	deirdreClothes,
	deirdreClothesColors,
	deirdreFlooring,
	deirdreFurniture,
	deirdreGames,
	deirdreGender,
	deirdreInterest,
	deirdreName,
	deirdreNameColor,
	deirdreMusic,
	deirdrePersonality,
	deirdreSpecies,
	deirdreStyle,
	deirdreWallpaper}

var (
	deirdreAstrology     = villagersAstrology{villagerAstrologyTaurus}
	deirdreBirthDay      = villagersBirthDay{4}
	deirdreBirthMonth    = villagersBirthMonth{time.May} // 5
	deirdreBubbleColor   = villagersBubbleColor{villagerBubbleColorDDCDCA}
	deirdreCategory      = villagersCategory{villagerCategoryA}
	deirdreClothes       = villagersClothes{villagerClothesFlowerSweater} // 7844
	deirdreClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorOrange}}
	deirdreFlooring      = villagersFlooring{villagerFlooringColoredLeavesFlooring}
	deirdreFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureLogStool, villagerFurnitureMushTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureTreesBountyArch, villagerFurnitureMushLowStool, villagerFurnitureMushLamp, villagerFurnitureHammock, villagerFurnitureMushLog, villagerFurnitureTreesBountyMobile, villagerFurnitureTreesBountyLamp}}
	deirdreGames         = villagersGames{[]VillagerGame{}} // TBD
	deirdreGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	deirdreInterest      = villagersInterest{villagerInterestPlay}
	deirdreName          = villagersName{villagerNameDeirdre}
	deirdreNameColor     = villagersNameColor{villagerNameColor4B4496}
	deirdreMusic         = villagersMusic{villagerMusicKKBossa}
	deirdrePersonality   = villagersPersonality{villagerPersonalityBigSister}
	deirdreSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	deirdreStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	deirdreWallpaper     = villagersWallpaper{villagerWallpaperAutumnWall}
)
