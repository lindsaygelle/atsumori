package atsumori

import "time"

// Snake is an Animal Crossing villager.
var Snake = villager{
	snakeAstrology,
	snakeBirthDay,
	snakeBirthMonth,
	snakeBubbleColor,
	snakeCategory,
	snakeClothes,
	snakeClothesColors,
	snakeFlooring,
	snakeFurniture,
	snakeGames,
	snakeGender,
	snakeInterest,
	snakeName,
	snakeNameColor,
	snakeMusic,
	snakePersonality,
	snakeSpecies,
	snakeStyle,
	snakeWallpaper}

var (
	snakeAstrology     = villagersAstrology{villagerAstrologyScorpio}
	snakeBirthDay      = villagersBirthDay{3}
	snakeBirthMonth    = villagersBirthMonth{time.November} // 11
	snakeBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	snakeCategory      = villagersCategory{villagerCategoryB}
	snakeClothes       = villagersClothes{} // 3460
	snakeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorBlue}}
	snakeFlooring      = villagersFlooring{villagerFlooringDirtFlooring}
	snakeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBonsaiShelf, villagerFurnitureBambooPartition, villagerFurnitureBambooPartition, villagerFurnitureTikiTorch, villagerFurnitureTikiTorch, villagerFurniturePondStone, villagerFurnitureBambooBench, villagerFurniturePot}}
	snakeGames         = villagersGames{[]VillagerGame{}} // TBD
	snakeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	snakeInterest      = villagersInterest{villagerInterestFitness}
	snakeName          = villagersName{villagerNameSnake}
	snakeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	snakeMusic         = villagersMusic{villagerMusicKingKK}
	snakePersonality   = villagersPersonality{villagerPersonalityJock}
	snakeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	snakeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	snakeWallpaper     = villagersWallpaper{villagerWallpaperMortarWall}
)
