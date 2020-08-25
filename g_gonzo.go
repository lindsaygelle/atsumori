package atsumori

import "time"

// Gonzo is an Animal Crossing villager
var Gonzo = villager{
	gonzoAstrology,
	gonzoBirthDay,
	gonzoBirthMonth,
	gonzoBubbleColor,
	gonzoCategory,
	gonzoClothes,
	gonzoClothesColors,
	gonzoFlooring,
	gonzoFurniture,
	gonzoGames,
	gonzoGender,
	gonzoInterest,
	gonzoName,
	gonzoNameColor,
	gonzoMusic,
	gonzoPersonality,
	gonzoSpecies,
	gonzoStyle,
	gonzoWallpaper}

var (
	gonzoAstrology     = villagersAstrology{villagerAstrologyLibra}
	gonzoBirthDay      = villagersBirthDay{13}
	gonzoBirthMonth    = villagersBirthMonth{time.October} // 10
	gonzoBubbleColor   = villagersBubbleColor{villagerBubbleColorBFBFBF}
	gonzoCategory      = villagersCategory{villagerCategoryA}
	gonzoClothes       = villagersClothes{villagerClothesReindeerSweater} // 4566
	gonzoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGreen}}
	gonzoFlooring      = villagersFlooring{villagerFlooringWoodenKnotFlooring}
	gonzoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodBurningStove, villagerFurnitureLogRoundTable, villagerFurnitureFirewood, villagerFurnitureIronwoodKitchenette, villagerFurniturePopUpToaster, villagerFurnitureTapestry, villagerFurnitureLogBed, villagerFurnitureGreenKilimStyleCarpet, villagerFurnitureLogDecorativeShelves, villagerFurniturePortableRecordPlayer}}
	gonzoGames         = villagersGames{[]VillagerGame{}} // TBD
	gonzoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	gonzoInterest      = villagersInterest{villagerInterestNature}
	gonzoName          = villagersName{villagerNameGonzo}
	gonzoNameColor     = villagersNameColor{villagerNameColor5E5E5E}
	gonzoMusic         = villagersMusic{villagerMusicSurfinKK}
	gonzoPersonality   = villagersPersonality{villagerPersonalityCranky}
	gonzoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	gonzoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	gonzoWallpaper     = villagersWallpaper{villagerWallpaperCabinWall}
)
