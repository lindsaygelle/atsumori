package atsumori

import "time"

// Audie is an Animal Crossing villager.
var Audie = villager{
	audieAstrology,
	audieBirthDay,
	audieBirthMonth,
	audieBubbleColor,
	audieCategory,
	audieClothes,
	audieClothesColors,
	audieFlooring,
	audieFurniture,
	audieGames,
	audieGender,
	audieInterest,
	audieName,
	audieNameColor,
	audieMusic,
	audiePersonality,
	audieSpecies,
	audieStyle,
	audieWallpaper}

var (
	audieAstrology     = villagersAstrology{villagerAstrologyVirgo}
	audieBirthDay      = villagersBirthDay{31}
	audieBirthMonth    = villagersBirthMonth{time.August} // 8
	audieBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	audieCategory      = villagersCategory{villagerCategoryA}
	audieClothes       = villagersClothes{} // 2687
	audieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorWhite}}
	audieFlooring      = villagersFlooring{villagerFlooringGreenFloralFlooring}
	audieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanEndTable, villagerFurnitureTapeDeck, villagerFurnitureOpenFrameKitchen, villagerFurnitureDinerNeonSign, villagerFurnitureMenuChalkboard, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureCardboardBox, villagerFurnitureFruitBasket, villagerFurnitureWallFan, villagerFurnitureRattanLowTable, villagerFurnitureCoconutJuice, villagerFurnitureYellowKitchenMat, villagerFurniturePalmTreeLamp}}
	audieGames         = villagersGames{[]VillagerGame{}} // TBD
	audieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	audieInterest      = villagersInterest{villagerInterestFitness}
	audieName          = villagersName{villagerNameAudie}
	audieNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	audieMusic         = villagersMusic{villagerMusicKKIsland}
	audiePersonality   = villagersPersonality{villagerPersonalityPeppy}
	audieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	audieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleActive}}
	audieWallpaper     = villagersWallpaper{villagerWallpaperTropicalVista}
)
