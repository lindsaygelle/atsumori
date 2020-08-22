package atsumori

import "time"

// Elise is an Animal Crossing villager
var Elise = villager{
	eliseAstrology,
	eliseBirthDay,
	eliseBirthMonth,
	eliseBubbleColor,
	eliseCategory,
	eliseClothes,
	eliseClothesColors,
	eliseFlooring,
	eliseFurniture,
	eliseGames,
	eliseGender,
	eliseInterest,
	eliseName,
	eliseNameColor,
	eliseMusic,
	elisePersonality,
	eliseSpecies,
	eliseStyle,
	eliseWallpaper}

var (
	eliseAstrology     = villagersAstrology{villagerAstrologyAries}
	eliseBirthDay      = villagersBirthDay{21}
	eliseBirthMonth    = villagersBirthMonth{time.March} // 3
	eliseBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	eliseCategory      = villagersCategory{villagerCategoryB}
	eliseClothes       = villagersClothes{} // 4791
	eliseClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorRed}}
	eliseFlooring      = villagersFlooring{villagerFlooringSimplePurpleFlooring}
	eliseFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureImperialPartition, villagerFurnitureAntiqueBed, villagerFurnitureRattanArmchair, villagerFurnitureOpenFrameKitchen, villagerFurnitureRattanEndTable, villagerFurnitureDenChair, villagerFurnitureMonstera, villagerFurnitureFloralSwag, villagerFurnitureAntiqueMiniTable, villagerFurnitureShowerBooth, villagerFurnitureDenDesk, villagerFurniturePhonograph, villagerFurnitureTypewriter}}
	eliseGames         = villagersGames{[]VillagerGame{}} // TBD
	eliseGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	eliseInterest      = villagersInterest{villagerInterestFashion}
	eliseName          = villagersName{villagerNameElise}
	eliseNameColor     = villagersNameColor{villagerNameColor8B5F57}
	eliseMusic         = villagersMusic{villagerMusicKKBossa}
	elisePersonality   = villagersPersonality{villagerPersonalitySnooty}
	eliseSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMonkey}}
	eliseStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	eliseWallpaper     = villagersWallpaper{villagerWallpaperAbstractWall}
)
