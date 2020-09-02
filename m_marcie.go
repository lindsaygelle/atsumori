package atsumori

import "time"

// Marcie is an Animal Crossing villager.
var Marcie = villager{
	marcieAstrology,
	marcieBirthDay,
	marcieBirthMonth,
	marcieBubbleColor,
	marcieCategory,
	marcieClothes,
	marcieClothesColors,
	marcieFlooring,
	marcieFurniture,
	marcieGames,
	marcieGender,
	marcieInterest,
	marcieName,
	marcieNameColor,
	marcieMusic,
	marciePersonality,
	marcieSpecies,
	marcieStyle,
	marcieWallpaper}

var (
	marcieAstrology     = villagersAstrology{villagerAstrologyGemini}
	marcieBirthDay      = villagersBirthDay{31}
	marcieBirthMonth    = villagersBirthMonth{time.May} // 5
	marcieBubbleColor   = villagersBubbleColor{villagerBubbleColorF993CE}
	marcieCategory      = villagersCategory{villagerCategoryA}
	marcieClothes       = villagersClothes{villagerClothesHeartApron} // 4429
	marcieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorBeige}}
	marcieFlooring      = villagersFlooring{villagerFlooringPineBoardFlooring}
	marcieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenLowTable, villagerFurnitureWhiteSimpleSmallMat, villagerFurnitureSystemKitchen, villagerFurnitureIroningBoard, villagerFurnitureAutomaticWasher, villagerFurnitureDoubleDoorRefrigerator, villagerFurnitureRingtoss, villagerFurnitureWoodenChest, villagerFurnitureWoodenDoubleBed, villagerFurnitureCuteMusicPlayer, villagerFurnitureBabyChair, villagerFurnitureMomsPlayfulKitchenMat, villagerFurnitureCuteWallMountedClock, villagerFurnitureMobile}}
	marcieGames         = villagersGames{[]VillagerGame{}} // TBD
	marcieGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	marcieInterest      = villagersInterest{villagerInterestNature}
	marcieName          = villagersName{villagerNameMarcie}
	marcieNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	marcieMusic         = villagersMusic{villagerMusicKKAria}
	marciePersonality   = villagersPersonality{villagerPersonalityNormal}
	marcieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKangaroo}}
	marcieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	marcieWallpaper     = villagersWallpaper{villagerWallpaperPinkQuiltWall}
)
