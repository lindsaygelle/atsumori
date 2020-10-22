package atsumori

import "time"

// Wolfgang is an Animal Crossing villager.
var Wolfgang = villager{
	wolfgangAstrology,
	wolfgangBirthDay,
	wolfgangBirthMonth,
	wolfgangBubbleColor,
	wolfgangCategory,
	wolfgangClothes,
	wolfgangClothesColors,
	wolfgangFlooring,
	wolfgangFurniture,
	wolfgangGames,
	wolfgangGender,
	wolfgangInterest,
	wolfgangName,
	wolfgangNameColor,
	wolfgangMusic,
	wolfgangPersonality,
	wolfgangSpecies,
	wolfgangStyle,
	wolfgangWallpaper}

var (
	wolfgangAstrology     = villagersAstrology{villagerAstrologySagittarius}
	wolfgangBirthDay      = villagersBirthDay{25}
	wolfgangBirthMonth    = villagersBirthMonth{time.November} // 11
	wolfgangBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	wolfgangCategory      = villagersCategory{villagerCategoryB}
	wolfgangClothes       = villagersClothes{} // 4207
	wolfgangClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGreen}}
	wolfgangFlooring      = villagersFlooring{villagerFlooringStripeFlooring}
	wolfgangFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodBurningStove, villagerFurnitureBrownWoodenDeckRug, villagerFurnitureBrownWoodenDeckRug, villagerFurnitureIronwoodClock, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureIronwoodKitchenette, villagerFurnitureIronwoodDresser, villagerFurniturePortableRecordPlayer, villagerFurnitureIronwoodBed, villagerFurnitureIronWallRack, villagerFurnitureIronwoodCupboard, villagerFurniturePennant, villagerFurnitureMicrowave, villagerFurnitureIronwoodLowTable, villagerFurniturePopUpToaster, villagerFurnitureIronwoodChair, villagerFurnitureTissueBox, villagerFurnitureProteinShakerBottle}}
	wolfgangGames         = villagersGames{[]VillagerGame{}} // TBD
	wolfgangGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	wolfgangInterest      = villagersInterest{villagerInterestEducation}
	wolfgangName          = villagersName{villagerNameWolfgang}
	wolfgangNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	wolfgangMusic         = villagersMusic{villagerMusicKKDB}
	wolfgangPersonality   = villagersPersonality{villagerPersonalityCranky}
	wolfgangSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	wolfgangStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	wolfgangWallpaper     = villagersWallpaper{villagerWallpaperDarkWoodenMosaicWall}
)
