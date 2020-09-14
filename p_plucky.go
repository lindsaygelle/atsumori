package atsumori

import "time"

// Plucky is an Animal Crossing villager.
var Plucky = villager{
	pluckyAstrology,
	pluckyBirthDay,
	pluckyBirthMonth,
	pluckyBubbleColor,
	pluckyCategory,
	pluckyClothes,
	pluckyClothesColors,
	pluckyFlooring,
	pluckyFurniture,
	pluckyGames,
	pluckyGender,
	pluckyInterest,
	pluckyName,
	pluckyNameColor,
	pluckyMusic,
	pluckyPersonality,
	pluckySpecies,
	pluckyStyle,
	pluckyWallpaper}

var (
	pluckyAstrology     = villagersAstrology{villagerAstrologyLibra}
	pluckyBirthDay      = villagersBirthDay{12}
	pluckyBirthMonth    = villagersBirthMonth{time.October} // 10
	pluckyBubbleColor   = villagersBubbleColor{villagerBubbleColor7C6559}
	pluckyCategory      = villagersCategory{villagerCategoryA}
	pluckyClothes       = villagersClothes{} // 8604
	pluckyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorYellow}}
	pluckyFlooring      = villagersFlooring{villagerFlooringStarrySandsFlooring}
	pluckyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCassettePlayer, villagerFurnitureSurfboard, villagerFurnitureSurfboard, villagerFurnitureSurfboard, villagerFurnitureSurfboard, villagerFurnitureSurfboard, villagerFurnitureRattanEndTable, villagerFurnitureFrozenTreatSet, villagerFurnitureSurfboard, villagerFurnitureLawnChair, villagerFurnitureMenuChalkboard, villagerFurnitureStall, villagerFurnitureShantyMat, villagerFurnitureShavedIceMaker, villagerFurnitureFrozenTreatSet}}
	pluckyGames         = villagersGames{[]VillagerGame{}} // TBD
	pluckyGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pluckyInterest      = villagersInterest{villagerInterestPlay}
	pluckyName          = villagersName{villagerNamePlucky}
	pluckyNameColor     = villagersNameColor{villagerNameColorEAC113}
	pluckyMusic         = villagersMusic{villagerMusicKKIsland}
	pluckyPersonality   = villagersPersonality{villagerPersonalityBigSister}
	pluckySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	pluckyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	pluckyWallpaper     = villagersWallpaper{villagerWallpaperTropicalVista}
)
