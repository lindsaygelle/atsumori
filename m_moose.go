package atsumori

import "time"

// Moose is an Animal Crossing villager.
var Moose = villager{
	mooseAstrology,
	mooseBirthDay,
	mooseBirthMonth,
	mooseBubbleColor,
	mooseCategory,
	mooseClothes,
	mooseClothesColors,
	mooseFlooring,
	mooseFurniture,
	mooseGames,
	mooseGender,
	mooseInterest,
	mooseName,
	mooseNameColor,
	mooseMusic,
	moosePersonality,
	mooseSpecies,
	mooseStyle,
	mooseWallpaper}

var (
	mooseAstrology     = villagersAstrology{villagerAstrologyVirgo}
	mooseBirthDay      = villagersBirthDay{13}
	mooseBirthMonth    = villagersBirthMonth{time.September} // 9
	mooseBubbleColor   = villagersBubbleColor{villagerBubbleColor7FA9FF}
	mooseCategory      = villagersCategory{villagerCategoryB}
	mooseClothes       = villagersClothes{villagerClothesBigStarTee} // 6822
	mooseClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorRed}}
	mooseFlooring      = villagersFlooring{villagerFlooringBlueCamoFlooring}
	mooseFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJukebox, villagerFurnitureBilliardTable, villagerFurnitureIronShelf, villagerFurnitureIronWallLamp, villagerFurnitureIronwoodBed, villagerFurniturePinballMachine, villagerFurnitureDartboard, villagerFurnitureRedCarpet, villagerFurnitureDinerSofa}}
	mooseGames         = villagersGames{[]VillagerGame{}} // TBD
	mooseGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	mooseInterest      = villagersInterest{villagerInterestFitness}
	mooseName          = villagersName{villagerNameMoose}
	mooseNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	mooseMusic         = villagersMusic{villagerMusicKKRockabilly}
	moosePersonality   = villagersPersonality{villagerPersonalityJock}
	mooseSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	mooseStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	mooseWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
