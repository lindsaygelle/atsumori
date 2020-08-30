package atsumori

import "time"

// Gaston is an Animal Crossing villager.
var Gaston = villager{
	gastonAstrology,
	gastonBirthDay,
	gastonBirthMonth,
	gastonBubbleColor,
	gastonCategory,
	gastonClothes,
	gastonClothesColors,
	gastonFlooring,
	gastonFurniture,
	gastonGames,
	gastonGender,
	gastonInterest,
	gastonName,
	gastonNameColor,
	gastonMusic,
	gastonPersonality,
	gastonSpecies,
	gastonStyle,
	gastonWallpaper}

var (
	gastonAstrology     = villagersAstrology{villagerAstrologyScorpio}
	gastonBirthDay      = villagersBirthDay{28}
	gastonBirthMonth    = villagersBirthMonth{time.October} // 10
	gastonBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	gastonCategory      = villagersCategory{villagerCategoryB}
	gastonClothes       = villagersClothes{villagerClothesColorBlockDressShirt} // 8489
	gastonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorOrange}}
	gastonFlooring      = villagersFlooring{villagerFlooringRamshackleFlooring}
	gastonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureKettleBathtub, villagerFurnitureClayFurnace, villagerFurnitureClothesline, villagerFurnitureOldFashionedWashtub, villagerFurnitureCardboardBed, villagerFurnitureLogStool, villagerFurnitureBambooStool, villagerFurnitureCardboardTable, villagerFurnitureFirewood, villagerFurniturePortableRadio, villagerFurnitureShantyMat}}
	gastonGames         = villagersGames{[]VillagerGame{}} // TBD
	gastonGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	gastonInterest      = villagersInterest{villagerInterestEducation}
	gastonName          = villagersName{villagerNameGaston}
	gastonNameColor     = villagersNameColor{villagerNameColor9B553A}
	gastonMusic         = villagersMusic{villagerMusicKKSong}
	gastonPersonality   = villagersPersonality{villagerPersonalityCranky}
	gastonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	gastonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleGorgeous}}
	gastonWallpaper     = villagersWallpaper{villagerWallpaperRamshackleWall}
)
