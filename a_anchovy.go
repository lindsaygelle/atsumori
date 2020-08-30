package atsumori

import "time"

// Anchovy is an Animal Crossing villager.
var Anchovy = villager{
	anchovyAstrology,
	anchovyBirthDay,
	anchovyBirthMonth,
	anchovyBubbleColor,
	anchovyCategory,
	anchovyClothes,
	anchovyClothesColors,
	anchovyFlooring,
	anchovyFurniture,
	anchovyGames,
	anchovyGender,
	anchovyInterest,
	anchovyName,
	anchovyNameColor,
	anchovyMusic,
	anchovyPersonality,
	anchovySpecies,
	anchovyStyle,
	anchovyWallpaper}

var (
	anchovyAstrology     = villagersAstrology{villagerAstrologyPisces}
	anchovyBirthDay      = villagersBirthDay{4}
	anchovyBirthMonth    = villagersBirthMonth{time.March}
	anchovyBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	anchovyCategory      = villagersCategory{villagerCategoryB}
	anchovyClothes       = villagersClothes{villagerClothesYodelSweater} //  3630
	anchovyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorColorful, villagerClothesColorOrange}}
	anchovyFlooring      = villagersFlooring{villagerFlooringSimpleRedFlooring}
	anchovyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBilliardTable, villagerFurnitureJukebox, villagerFurnitureDartboard, villagerFurnitureFoosballTable, villagerFurniturePianoBench, villagerFurnitureUprightPiano, villagerFurniturePinballMachine, villagerFurnitureFoldingFloorLamp, villagerFurnitureBoxCornerSofa, villagerFurnitureBoxSofa, villagerFurnitureBoxSofa}}
	anchovyGames         = villagersGames{[]VillagerGame{}} // TBD
	anchovyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	anchovyInterest      = villagersInterest{villagerInterestPlay}
	anchovyName          = villagersName{villagerNameAnchovy}
	anchovyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	anchovyMusic         = villagersMusic{villagerMusicKKRagtime}
	anchovyPersonality   = villagersPersonality{villagerPersonalityLazy}
	anchovySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	anchovyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	anchovyWallpaper     = villagersWallpaper{villagerWallpaperModernWoodWall}
)
