package atsumori

import "time"

// Puck is an Animal Crossing villager.
var Puck = villager{
	puckAstrology,
	puckBirthDay,
	puckBirthMonth,
	puckBubbleColor,
	puckCategory,
	puckClothes,
	puckClothesColors,
	puckFlooring,
	puckFurniture,
	puckGames,
	puckGender,
	puckInterest,
	puckName,
	puckNameColor,
	puckMusic,
	puckPersonality,
	puckSpecies,
	puckStyle,
	puckWallpaper}

var (
	puckAstrology     = villagersAstrology{villagerAstrologyPisces}
	puckBirthDay      = villagersBirthDay{21}
	puckBirthMonth    = villagersBirthMonth{time.February} // 2
	puckBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	puckCategory      = villagersCategory{villagerCategoryB}
	puckClothes       = villagersClothes{} // 5831
	puckClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorRed}}
	puckFlooring      = villagersFlooring{villagerFlooringIceFlooring}
	puckFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWeightBench, villagerFurnitureFrozenPartition, villagerFurnitureSleigh, villagerFurnitureBarbell, villagerFurniturePunchingBag, villagerFurnitureSpeedBag, villagerFurnitureTapeDeck, villagerFurnitureShowerBooth, villagerFurnitureTreadmill, villagerFurnitureExerciseBike, villagerFurnitureFrozenCounter, villagerFurnitureProteinShakerBottle}}
	puckGames         = villagersGames{[]VillagerGame{}} // TBD
	puckGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	puckInterest      = villagersInterest{villagerInterestFitness}
	puckName          = villagersName{villagerNamePuck}
	puckNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	puckMusic         = villagersMusic{villagerMusicKKMarch}
	puckPersonality   = villagersPersonality{villagerPersonalityLazy}
	puckSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	puckStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	puckWallpaper     = villagersWallpaper{villagerWallpaperIceWall}
)
