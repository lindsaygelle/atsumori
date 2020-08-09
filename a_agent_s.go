package atsumori

// AgentS is an Animal Crossing villager.
var AgentS Villager = villager{
	agentSAstrology,
	agentSBubbleColor,
	agentSCategory,
	agentSClothes,
	agentSClothesColors,
	agentSFlooring,
	agentSFurniture,
	agentSGames,
	agentSGender,
	agentSInterest,
	agentSName,
	agentSNameColor,
	agentSMusic,
	agentSPersonality,
	agentSSpecies,
	agentSStyle,
	agentSWallpaper}

var (
	agentSAstrology     = villagersAstrology{villagerAstrologyCancer}
	agentSBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	agentSCategory      = villagersCategory{villagerCategoryB}
	agentSClothes       = villagersClothes{villagerClothesNo2Shirt}
	agentSClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorBlack}}
	agentSFlooring      = villagersFlooring{villagerFlooringColorfulTileFlooring}
	agentSFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurniturePullUpBarStand, villagerFurnitureBlueCorner, villagerFurnitureBeachTowel, villagerFurnitureLongBathtub, villagerFurnitureTreadmill, villagerFurnitureImperialPartition, villagerFurnitureWoodenStool, villagerFurnitureCassettePlayer, villagerFurnitureThrowbackWallClock, villagerFurnitureWoodenChest, villagerFurnitureThrowbackWrestlingFigure, villagerFurnitureBrownWoodenDeckRug, villagerFurnitureDryingRack}}
	agentSGames         = villagersGames{[12]VillagerGame{}} // TBD
	agentSGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	agentSInterest      = villagersInterest{villagerInterestFitness}
	agentSName          = villagersName{villagerNameAgentS}
	agentSNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	agentSMusic         = villagersMusic{}
	agentSPersonality   = villagersPersonality{villagerPersonalityPeppy}
	agentSSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	agentSStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	agentSWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)