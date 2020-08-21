package atsumori

import "time"

// Coco is an Animal Crossing villager
var Coco = villager{
	cocoAstrology,
	cocoBirthDay,
	cocoBirthMonth,
	cocoBubbleColor,
	cocoCategory,
	cocoClothes,
	cocoClothesColors,
	cocoFlooring,
	cocoFurniture,
	cocoGames,
	cocoGender,
	cocoInterest,
	cocoName,
	cocoNameColor,
	cocoMusic,
	cocoPersonality,
	cocoSpecies,
	cocoStyle,
	cocoWallpaper}

var (
	cocoAstrology     = villagersAstrology{villagerAstrologyPisces}
	cocoBirthDay      = villagersBirthDay{1}
	cocoBirthMonth    = villagersBirthMonth{time.March} // 3
	cocoBubbleColor   = villagersBubbleColor{villagerBubbleColorE4B887}
	cocoCategory      = villagersCategory{villagerCategoryB}
	cocoClothes       = villagersClothes{} // 4167
	cocoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorGreen}}
	cocoFlooring      = villagersFlooring{villagerFlooringFloralRushMatFlooring}
	cocoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureOldFashionedWashtub, villagerFurnitureClassicPitcher, villagerFurnitureTikiTorch, villagerFurnitureTikiTorch, villagerFurnitureStoneStool, villagerFurnitureStoneStool, villagerFurnitureHayBed, villagerFurnitureClothesline, villagerFurnitureStoneTable, villagerFurnitureClassicPitcher, villagerFurnitureUnglazedDishSet, villagerFurniturePot}}
	cocoGames         = villagersGames{[]VillagerGame{}} // TBD
	cocoGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	cocoInterest      = villagersInterest{villagerInterestEducation}
	cocoName          = villagersName{villagerNameCoco}
	cocoNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	cocoMusic         = villagersMusic{villagerMusicKKJongara}
	cocoPersonality   = villagersPersonality{villagerPersonalityNormal}
	cocoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	cocoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	cocoWallpaper     = villagersWallpaper{villagerWallpaperRammedEarthWall}
)
