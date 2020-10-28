package atsumori

import "time"

// Yuka is an Animal Crossing villager.
var Yuka = villager{
	yukaAstrology,
	yukaBirthDay,
	yukaBirthMonth,
	yukaBubbleColor,
	yukaCategory,
	yukaClothes,
	yukaClothesColors,
	yukaFlooring,
	yukaFurniture,
	yukaGames,
	yukaGender,
	yukaInterest,
	yukaName,
	yukaNameColor,
	yukaMusic,
	yukaPersonality,
	yukaSpecies,
	yukaStyle,
	yukaWallpaper}

var (
	yukaAstrology     = villagersAstrology{villagerAstrologyCancer}
	yukaBirthDay      = villagersBirthDay{20}
	yukaBirthMonth    = villagersBirthMonth{time.July} // 7
	yukaBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	yukaCategory      = villagersCategory{villagerCategoryB}
	yukaClothes       = villagersClothes{villagerClothesAranKnitSweater} // 7676
	yukaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorYellow}}
	yukaFlooring      = villagersFlooring{villagerFlooringSimplePurpleFlooring}
	yukaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBed, villagerFurnitureAntiqueVanity, villagerFurnitureDoubleSofa, villagerFurnitureImperialLowTable, villagerFurnitureAntiqueConsoleTable, villagerFurniturePhonograph, villagerFurniturePendulumClock, villagerFurnitureAntiqueWardrobe, villagerFurnitureAntiqueMiniTable, villagerFurnitureIncenseBurner}}
	yukaGames         = villagersGames{[]VillagerGame{}} // TBD
	yukaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	yukaInterest      = villagersInterest{villagerInterestFashion}
	yukaName          = villagersName{villagerNameYuka}
	yukaNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	yukaMusic         = villagersMusic{villagerMusicSoulfulKK}
	yukaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	yukaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	yukaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleElegant}}
	yukaWallpaper     = villagersWallpaper{villagerWallpaperBeigeArtDecoWall}
)
