package atsumori

import "time"

// Amelia is an Animal Crossing villager
var Amelia = villager{
	ameliaAstrology,
	ameliaBirthDay,
	ameliaBirthMonth,
	ameliaBubbleColor,
	ameliaCategory,
	ameliaClothes,
	ameliaClothesColors,
	ameliaFlooring,
	ameliaFurniture,
	ameliaGames,
	ameliaGender,
	ameliaInterest,
	ameliaName,
	ameliaNameColor,
	ameliaMusic,
	ameliaPersonality,
	ameliaSpecies,
	ameliaStyle,
	ameliaWallpaper}

var (
	ameliaAstrology     = villagersAstrology{villagerAstrologyScorpio}
	ameliaBirthDay      = villagersBirthDay{19}
	ameliaBirthMonth    = villagersBirthMonth{time.November} // 11
	ameliaBubbleColor   = villagersBubbleColor{villagerBubbleColorFF4040}
	ameliaCategory      = villagersCategory{villagerCategoryB}
	ameliaClothes       = villagersClothes{villagerClothesAranKnitSweater} // 7675
	ameliaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorWhite}}
	ameliaFlooring      = villagersFlooring{villagerFlooringCommonFlooring}
	ameliaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLogRoundTable, villagerFurnitureLogBench, villagerFurnitureLogBed, villagerFurniturePhonograph, villagerFurnitureAcousticGuitar, villagerFurnitureRedKilimStyleCarpet, villagerFurnitureLogExtraLongSofa, villagerFurnitureLogDecorativeShelves, villagerFurnitureOldFashionedWashtub, villagerFurnitureClothesline, villagerFurnitureMiniCactusSet}}
	ameliaGames         = villagersGames{[]VillagerGame{}} // TBD
	ameliaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	ameliaInterest      = villagersInterest{villagerInterestMusic}
	ameliaName          = villagersName{villagerNameAmelia}
	ameliaNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	ameliaMusic         = villagersMusic{villagerMusicKKCondor}
	ameliaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	ameliaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesEagle}}
	ameliaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleElegant}}
	ameliaWallpaper     = villagersWallpaper{villagerWallpaperSummitWall}
)
