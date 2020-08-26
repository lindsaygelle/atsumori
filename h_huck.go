package atsumori

import "time"

// Huck is an Animal Crossing villager
var Huck = villager{
	huckAstrology,
	huckBirthDay,
	huckBirthMonth,
	huckBubbleColor,
	huckCategory,
	huckClothes,
	huckClothesColors,
	huckFlooring,
	huckFurniture,
	huckGames,
	huckGender,
	huckInterest,
	huckName,
	huckNameColor,
	huckMusic,
	huckPersonality,
	huckSpecies,
	huckStyle,
	huckWallpaper}

var (
	huckAstrology     = villagersAstrology{villagerAstrologyCancer}
	huckBirthDay      = villagersBirthDay{9}
	huckBirthMonth    = villagersBirthMonth{time.July} // 7
	huckBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	huckCategory      = villagersCategory{villagerCategoryA}
	huckClothes       = villagersClothes{villagerClothesStripedTank} // 3383
	huckClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorYellow}}
	huckFlooring      = villagersFlooring{villagerFlooringDirtFlooring}
	huckFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureScarecrow, villagerFurnitureClayFurnace, villagerFurnitureShantyMat, villagerFurniturePondStone, villagerFurnitureLogRoundTable, villagerFurnitureClassicPitcher, villagerFurnitureHayBed, villagerFurnitureBambooSpeaker, villagerFurnitureLogWallMountedClock, villagerFurniturePot}}
	huckGames         = villagersGames{[]VillagerGame{}} // TBD
	huckGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	huckInterest      = villagersInterest{villagerInterestFitness}
	huckName          = villagersName{villagerNameHuck}
	huckNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	huckMusic         = villagersMusic{villagerMusicToTheEdge}
	huckPersonality   = villagersPersonality{villagerPersonalitySmug}
	huckSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	huckStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	huckWallpaper     = villagersWallpaper{villagerWallpaperStrawWall}
)
