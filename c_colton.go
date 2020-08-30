package atsumori

import "time"

// Colton is an Animal Crossing villager.
var Colton = villager{
	coltonAstrology,
	coltonBirthDay,
	coltonBirthMonth,
	coltonBubbleColor,
	coltonCategory,
	coltonClothes,
	coltonClothesColors,
	coltonFlooring,
	coltonFurniture,
	coltonGames,
	coltonGender,
	coltonInterest,
	coltonName,
	coltonNameColor,
	coltonMusic,
	coltonPersonality,
	coltonSpecies,
	coltonStyle,
	coltonWallpaper}

var (
	coltonAstrology     = villagersAstrology{villagerAstrologyGemini}
	coltonBirthDay      = villagersBirthDay{22}
	coltonBirthMonth    = villagersBirthMonth{time.May} // 5
	coltonBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	coltonCategory      = villagersCategory{villagerCategoryA}
	coltonClothes       = villagersClothes{villagerClothesPrincesTunic} // 4433
	coltonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorRed}}
	coltonFlooring      = villagersFlooring{villagerFlooringPalaceTile}
	coltonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLilyRecordPlayer, villagerFurnitureDoubleSofa, villagerFurnitureAquariusUrn, villagerFurniturePianoBench, villagerFurnitureAntiqueBed, villagerFurnitureGrandPiano, villagerFurnitureFireplace, villagerFurnitureCuckooClock, villagerFurnitureDalaHorse}}
	coltonGames         = villagersGames{[]VillagerGame{}} // TBD
	coltonGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	coltonInterest      = villagersInterest{villagerInterestNature}
	coltonName          = villagersName{villagerNameColton}
	coltonNameColor     = villagersNameColor{villagerNameColor848484}
	coltonMusic         = villagersMusic{villagerMusicKKChorale}
	coltonPersonality   = villagersPersonality{villagerPersonalitySmug}
	coltonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	coltonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	coltonWallpaper     = villagersWallpaper{villagerWallpaperPalaceWall}
)
