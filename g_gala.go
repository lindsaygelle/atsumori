package atsumori

import "time"

// Gala is an Animal Crossing villager
var Gala = villager{
	galaAstrology,
	galaBirthDay,
	galaBirthMonth,
	galaBubbleColor,
	galaCategory,
	galaClothes,
	galaClothesColors,
	galaFlooring,
	galaFurniture,
	galaGames,
	galaGender,
	galaInterest,
	galaName,
	galaNameColor,
	galaMusic,
	galaPersonality,
	galaSpecies,
	galaStyle,
	galaWallpaper}

var (
	galaAstrology     = villagersAstrology{villagerAstrologyPisces}
	galaBirthDay      = villagersBirthDay{5}
	galaBirthMonth    = villagersBirthMonth{time.March} // 3
	galaBubbleColor   = villagersBubbleColor{villagerBubbleColorFEEAE7}
	galaCategory      = villagersCategory{villagerCategoryB}
	galaClothes       = villagersClothes{} // 6901
	galaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorWhite}}
	galaFlooring      = villagersFlooring{villagerFlooringLightWoodPatternFlooring}
	galaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureRedCarpet, villagerFurnitureJailBars, villagerFurnitureCardboardBox, villagerFurnitureDocumentStack, villagerFurnitureDenChair, villagerFurnitureFireplace, villagerFurniturePhonograph, villagerFurnitureSurveillanceCamera, villagerFurnitureAntiqueMiniTable, villagerFurnitureDenDesk, villagerFurnitureLuckygoldCat, villagerFurnitureLaptop, villagerFurnitureSafe, villagerFurnitureSafe, villagerFurnitureWallClock, villagerFurnitureAluminumBriefcase, villagerFurnitureGoldBars}}
	galaGames         = villagersGames{[]VillagerGame{}} // TBD
	galaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	galaInterest      = villagersInterest{villagerInterestEducation}
	galaName          = villagersName{villagerNameGala}
	galaNameColor     = villagersNameColor{villagerNameColor97858E}
	galaMusic         = villagersMusic{villagerMusicAnimalCity}
	galaPersonality   = villagersPersonality{villagerPersonalityNormal}
	galaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPig}}
	galaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	galaWallpaper     = villagersWallpaper{villagerWallpaperStatelyWall}
)
