package atsumori

import "time"

// Cobb is an Animal Crossing villager.
var Cobb = villager{
	cobbAstrology,
	cobbBirthDay,
	cobbBirthMonth,
	cobbBubbleColor,
	cobbCategory,
	cobbClothes,
	cobbClothesColors,
	cobbFlooring,
	cobbFurniture,
	cobbGames,
	cobbGender,
	cobbInterest,
	cobbName,
	cobbNameColor,
	cobbMusic,
	cobbPersonality,
	cobbSpecies,
	cobbStyle,
	cobbWallpaper}

var (
	cobbAstrology     = villagersAstrology{villagerAstrologyLibra}
	cobbBirthDay      = villagersBirthDay{7}
	cobbBirthMonth    = villagersBirthMonth{time.October} // 10
	cobbBubbleColor   = villagersBubbleColor{villagerBubbleColorA8E989}
	cobbCategory      = villagersCategory{villagerCategoryB}
	cobbClothes       = villagersClothes{villagerClothesFlannelShirt} // 8946
	cobbClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorBlue}}
	cobbFlooring      = villagersFlooring{villagerFlooringConcreteFlooring}
	cobbFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSleepingBag, villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenBookshelf, villagerFurnitureDocumentStack, villagerFurnitureHomeworkSet, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureSloppyRug, villagerFurnitureStackOfBooks, villagerFurnitureWhiteboard, villagerFurnitureCardboardBox, villagerFurniturePortableRadio, villagerFurnitureScatteredPapers, villagerFurnitureCardboardSofa, villagerFurnitureScatteredPapers, villagerFurnitureStackOfBooks, villagerFurnitureStackOfBooks}}
	cobbGames         = villagersGames{[]VillagerGame{}} // TBD
	cobbGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	cobbInterest      = villagersInterest{villagerInterestEducation}
	cobbName          = villagersName{villagerNameCobb}
	cobbNameColor     = villagersNameColor{villagerNameColor878E86}
	cobbMusic         = villagersMusic{villagerMusicPondering}
	cobbPersonality   = villagersPersonality{villagerPersonalityJock}
	cobbSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	cobbStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleActive}}
	cobbWallpaper     = villagersWallpaper{villagerWallpaperGrayShantyWall}
)
