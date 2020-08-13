package atsumori

import "time"

// Bella is an Animal Crossing villager
var Bella = villager{
	bellaAstrology,
	bellaBirthDay,
	bellaBirthMonth,
	bellaBubbleColor,
	bellaCategory,
	bellaClothes,
	bellaClothesColors,
	bellaFlooring,
	bellaFurniture,
	bellaGames,
	bellaGender,
	bellaInterest,
	bellaName,
	bellaNameColor,
	bellaMusic,
	bellaPersonality,
	bellaSpecies,
	bellaStyle,
	bellaWallpaper}

var (
	bellaAstrology     = villagersAstrology{villagerAstrologySagittarius}
	bellaBirthDay      = villagersBirthDay{28}
	bellaBirthMonth    = villagersBirthMonth{time.December} // 12
	bellaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	bellaCategory      = villagersCategory{villagerCategoryB}
	bellaClothes       = villagersClothes{} // 4559
	bellaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorPurple}}
	bellaFlooring      = villagersFlooring{villagerFlooringSkullPrintFlooring}
	bellaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureIronShelf, villagerFurnitureIronWallLamp, villagerFurnitureBathroomTowelRack, villagerFurnitureShowerSet, villagerFurnitureDinerMiniTable, villagerFurnitureDinerCounterTable, villagerFurnitureDinerCounterChair, villagerFurnitureDinerChair, villagerFurnitureThrowbackSkullRadio, villagerFurnitureMug, villagerFurnitureClawFootTub, villagerFurnitureLoftBedWithDesk, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureElectronicsKit, villagerFurnitureWallMountedToolBoard}}
	bellaGames         = villagersGames{[]VillagerGame{}} // TBD
	bellaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	bellaInterest      = villagersInterest{villagerInterestMusic}
	bellaName          = villagersName{villagerNameBella}
	bellaNameColor     = villagersNameColor{villagerNameColor848484}
	bellaMusic         = villagersMusic{} // K.K. Metal
	bellaPersonality   = villagersPersonality{villagerPersonalityPeppy}
	bellaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	bellaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	bellaWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
