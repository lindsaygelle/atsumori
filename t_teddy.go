package atsumori

import "time"

// Teddy is an Animal Crossing villager.
var Teddy = villager{
	teddyAstrology,
	teddyBirthDay,
	teddyBirthMonth,
	teddyBubbleColor,
	teddyCategory,
	teddyClothes,
	teddyClothesColors,
	teddyFlooring,
	teddyFurniture,
	teddyGames,
	teddyGender,
	teddyInterest,
	teddyName,
	teddyNameColor,
	teddyMusic,
	teddyPersonality,
	teddySpecies,
	teddyStyle,
	teddyWallpaper}

var (
	teddyAstrology     = villagersAstrology{villagerAstrologyLibra}
	teddyBirthDay      = villagersBirthDay{26}
	teddyBirthMonth    = villagersBirthMonth{time.September} // 9
	teddyBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	teddyCategory      = villagersCategory{villagerCategoryB}
	teddyClothes       = villagersClothes{villagerClothesEnergeticSweater} // 5262
	teddyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorOrange}}
	teddyFlooring      = villagersFlooring{villagerFlooringCorkFlooring}
	teddyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLogExtraLongSofa, villagerFurnitureDIYWorkbench, villagerFurnitureLogBed, villagerFurnitureLogDiningTable, villagerFurnitureFanPalm, villagerFurnitureLogWallMountedClock, villagerFurnitureMacrameTapestry, villagerFurnitureLogDecorativeShelves, villagerFurniturePortableRecordPlayer}}
	teddyGames         = villagersGames{[]VillagerGame{}} // TBD
	teddyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	teddyInterest      = villagersInterest{villagerInterestFitness}
	teddyName          = villagersName{villagerNameTeddy}
	teddyNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	teddyMusic         = villagersMusic{villagerMusicMrKK}
	teddyPersonality   = villagersPersonality{villagerPersonalityJock}
	teddySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBear}}
	teddyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	teddyWallpaper     = villagersWallpaper{villagerWallpaperCabinWall}
)
