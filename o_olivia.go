package atsumori

import "time"

// Olivia is an Animal Crossing villager.
var Olivia = villager{
	oliviaAstrology,
	oliviaBirthDay,
	oliviaBirthMonth,
	oliviaBubbleColor,
	oliviaCategory,
	oliviaClothes,
	oliviaClothesColors,
	oliviaFlooring,
	oliviaFurniture,
	oliviaGames,
	oliviaGender,
	oliviaInterest,
	oliviaName,
	oliviaNameColor,
	oliviaMusic,
	oliviaPersonality,
	oliviaSpecies,
	oliviaStyle,
	oliviaWallpaper}

var (
	oliviaAstrology     = villagersAstrology{villagerAstrologyAquarius}
	oliviaBirthDay      = villagersBirthDay{3}
	oliviaBirthMonth    = villagersBirthMonth{time.February} // 2
	oliviaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	oliviaCategory      = villagersCategory{villagerCategoryB}
	oliviaClothes       = villagersClothes{} // 3696
	oliviaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorBlack}}
	oliviaFlooring      = villagersFlooring{villagerFlooringSimpleBlueFlooring}
	oliviaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueMiniTable, villagerFurnitureAntiqueConsoleTable, villagerFurniturePianoBench, villagerFurnitureAntiqueClock, villagerFurnitureRattanLowTable, villagerFurnitureDoubleSofa, villagerFurnitureGrandPiano, villagerFurnitureFireplace, villagerFurniturePhonograph}}
	oliviaGames         = villagersGames{[]VillagerGame{}} // TBD
	oliviaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	oliviaInterest      = villagersInterest{villagerInterestMusic}
	oliviaName          = villagersName{villagerNameOlivia}
	oliviaNameColor     = villagersNameColor{villagerNameColor848484}
	oliviaMusic         = villagersMusic{villagerMusicKKSonata}
	oliviaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	oliviaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	oliviaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleElegant}}
	oliviaWallpaper     = villagersWallpaper{villagerWallpaperPurpleRoseWall}
)
