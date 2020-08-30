package atsumori

import "time"

// Francine is an Animal Crossing villager.
var Francine = villager{
	francineAstrology,
	francineBirthDay,
	francineBirthMonth,
	francineBubbleColor,
	francineCategory,
	francineClothes,
	francineClothesColors,
	francineFlooring,
	francineFurniture,
	francineGames,
	francineGender,
	francineInterest,
	francineName,
	francineNameColor,
	francineMusic,
	francinePersonality,
	francineSpecies,
	francineStyle,
	francineWallpaper}

var (
	francineAstrology     = villagersAstrology{villagerAstrologyAquarius}
	francineBirthDay      = villagersBirthDay{22}
	francineBirthMonth    = villagersBirthMonth{time.January} // 1
	francineBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	francineCategory      = villagersCategory{villagerCategoryB}
	francineClothes       = villagersClothes{} // 8187
	francineClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorBlack}}
	francineFlooring      = villagersFlooring{villagerFlooringBlueDotFlooring}
	francineFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanArmchair, villagerFurnitureRattanBed, villagerFurnitureShowerSet, villagerFurnitureBathroomTowelRack, villagerFurnitureCuteDIYTable, villagerFurnitureCuteTeaTable, villagerFurnitureCuteChair, villagerFurnitureCuteFloorLamp, villagerFurnitureWoodenChest, villagerFurnitureCuteVanity, villagerFurnitureChrissysPhoto, villagerFurniturePortableRecordPlayer, villagerFurnitureClawFootTub, villagerFurnitureCuteWallMountedClock}}
	francineGames         = villagersGames{[]VillagerGame{}} // TBD
	francineGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	francineInterest      = villagersInterest{villagerInterestFashion}
	francineName          = villagersName{villagerNameFrancine}
	francineNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	francineMusic         = villagersMusic{villagerMusicKKCruisin}
	francinePersonality   = villagersPersonality{villagerPersonalitySnooty}
	francineSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	francineStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	francineWallpaper     = villagersWallpaper{villagerWallpaperCuteBlueWall}
)
