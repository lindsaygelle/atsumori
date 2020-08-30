package atsumori

import "time"

// Kidd is an Animal Crossing villager.
var Kidd = villager{
	kiddAstrology,
	kiddBirthDay,
	kiddBirthMonth,
	kiddBubbleColor,
	kiddCategory,
	kiddClothes,
	kiddClothesColors,
	kiddFlooring,
	kiddFurniture,
	kiddGames,
	kiddGender,
	kiddInterest,
	kiddName,
	kiddNameColor,
	kiddMusic,
	kiddPersonality,
	kiddSpecies,
	kiddStyle,
	kiddWallpaper}

var (
	kiddAstrology     = villagersAstrology{villagerAstrologyCancer}
	kiddBirthDay      = villagersBirthDay{28}
	kiddBirthMonth    = villagersBirthMonth{time.June} // 6
	kiddBubbleColor   = villagersBubbleColor{villagerBubbleColorD9D9FF}
	kiddCategory      = villagersCategory{villagerCategoryB}
	kiddClothes       = villagersClothes{villagerClothesTailoredJacket} // 3474
	kiddClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorWhite}}
	kiddFlooring      = villagersFlooring{villagerFlooringPurpleDesertTileFlooring}
	kiddFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureRattanBed, villagerFurnitureLongBathtub, villagerFurnitureMonstera, villagerFurnitureTanklessToilet, villagerFurnitureRocketLamp, villagerFurnitureFireplace, villagerFurniturePhonograph, villagerFurnitureRattanEndTable, villagerFurnitureBathroomTowelRack, villagerFurnitureTissueBox, villagerFurnitureShowerSet}}
	kiddGames         = villagersGames{[]VillagerGame{}} // TBD
	kiddGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	kiddInterest      = villagersInterest{villagerInterestEducation}
	kiddName          = villagersName{villagerNameKidd}
	kiddNameColor     = villagersNameColor{villagerNameColor4F5D72}
	kiddMusic         = villagersMusic{villagerMusicKKTango}
	kiddPersonality   = villagersPersonality{villagerPersonalitySmug}
	kiddSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesGoat}}
	kiddStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	kiddWallpaper     = villagersWallpaper{villagerWallpaperGrayMoldedPanelWall}
)
