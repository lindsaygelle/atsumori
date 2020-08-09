package atsumori

import "time"

// AgentS is an Animal Crossing villager.
var AgentS Villager = villager{
	agentSAstrology,
	agentSBirthDay,
	agentSBirthMonth,
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
	agentSBirthDay      = villagersBirthDay{2}
	agentSBirthMonth    = villagersBirthMonth{time.July}
	agentSBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	agentSCategory      = villagersCategory{villagerCategoryB}
	agentSClothes       = villagersClothes{villagerClothesNo2Shirt}
	agentSClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorBlack}}
	agentSFlooring      = villagersFlooring{villagerFlooringColorfulTileFlooring}
	agentSFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureBlueCorner, villagerFurnitureBrownWoodenDeckRug, villagerFurnitureCassettePlayer, villagerFurnitureDryingRack, villagerFurnitureImperialPartition, villagerFurniturePullUpBarStand, villagerFurnitureThrowbackWallClock, villagerFurnitureThrowbackWrestlingFigure, villagerFurnitureLongBathtub, villagerFurnitureTreadmill, villagerFurnitureWoodenChest, villagerFurnitureWoodenStool}}
	agentSGames         = villagersGames{[]VillagerGame{}} // TBD
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
