package atsumori

import "time"

// Cesar is an Animal Crossing villager.
var Cesar = villager{
	cesarAstrology,
	cesarBirthDay,
	cesarBirthMonth,
	cesarBubbleColor,
	cesarCategory,
	cesarClothes,
	cesarClothesColors,
	cesarFlooring,
	cesarFurniture,
	cesarGames,
	cesarGender,
	cesarInterest,
	cesarName,
	cesarNameColor,
	cesarMusic,
	cesarPersonality,
	cesarSpecies,
	cesarStyle,
	cesarWallpaper}

var (
	cesarAstrology     = villagersAstrology{villagerAstrologyVirgo}
	cesarBirthDay      = villagersBirthDay{6}
	cesarBirthMonth    = villagersBirthMonth{time.September} // 9
	cesarBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	cesarCategory      = villagersCategory{villagerCategoryB}
	cesarClothes       = villagersClothes{villagerClothesATee} // 8532
	cesarClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	cesarFlooring      = villagersFlooring{villagerFlooringSimpleBlueFlooring}
	cesarFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureMonstera, villagerFurniturePianoBench, villagerFurnitureDinerCounterChair, villagerFurnitureRattanEndTable, villagerFurnitureDinerCounterChair, villagerFurniturePhonograph, villagerFurnitureDinerCounterChair, villagerFurnitureUprightPiano, villagerFurnitureDinerCounterChair, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureOpenFrameKitchen, villagerFurnitureSimpleKettle, villagerFurnitureMiniFridge, villagerFurnitureDinnerware, villagerFurnitureDinerCounterTable, villagerFurnitureFragranceSticks, villagerFurnitureDinerCounterTable, villagerFurnitureMug, villagerFurnitureMagneticKnifeRack, villagerFurniturePotRack}}
	cesarGames         = villagersGames{[]VillagerGame{}} // TBD
	cesarGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	cesarInterest      = villagersInterest{villagerInterestFitness}
	cesarName          = villagersName{villagerNameCesar}
	cesarNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	cesarMusic         = villagersMusic{villagerMusicKKCruisin}
	cesarPersonality   = villagersPersonality{villagerPersonalityCranky}
	cesarSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGorilla}}
	cesarStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleElegant}}
	cesarWallpaper     = villagersWallpaper{villagerWallpaperCityscapeWall}
)
