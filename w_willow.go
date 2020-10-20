package atsumori

import "time"

// Willow is an Animal Crossing villager.
var Willow = villager{
	willowAstrology,
	willowBirthDay,
	willowBirthMonth,
	willowBubbleColor,
	willowCategory,
	willowClothes,
	willowClothesColors,
	willowFlooring,
	willowFurniture,
	willowGames,
	willowGender,
	willowInterest,
	willowName,
	willowNameColor,
	willowMusic,
	willowPersonality,
	willowSpecies,
	willowStyle,
	willowWallpaper}

var (
	willowAstrology     = villagersAstrology{villagerAstrologySagittarius}
	willowBirthDay      = villagersBirthDay{26}
	willowBirthMonth    = villagersBirthMonth{time.November} // 11
	willowBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	willowCategory      = villagersCategory{villagerCategoryB}
	willowClothes       = villagersClothes{} // 4766
	willowClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorPink}}
	willowFlooring      = villagersFlooring{villagerFlooringArabesqueFlooring}
	willowFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenStool, villagerFurnitureOldSewingMachine, villagerFurnitureCuckooClock, villagerFurnitureImperialLowTable, villagerFurnitureAntiqueBed, villagerFurnitureMonstera, villagerFurnitureMusicStand, villagerFurnitureSturdySewingBox, villagerFurnitureSewingProject, villagerFurnitureCello, villagerFurnitureAntiqueConsoleTable, villagerFurnitureFireplace, villagerFurnitureFragranceSticks, villagerFurnitureTreesBountyLamp}}
	willowGames         = villagersGames{[]VillagerGame{}} // TBD
	willowGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	willowInterest      = villagersInterest{villagerInterestFashion}
	willowName          = villagersName{villagerNameWillow}
	willowNameColor     = villagersNameColor{villagerNameColor879B96}
	willowMusic         = villagersMusic{villagerMusicKKSoul}
	willowPersonality   = villagersPersonality{villagerPersonalitySnooty}
	willowSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	willowStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleGorgeous}}
	willowWallpaper     = villagersWallpaper{villagerWallpaperArchedWindowWall}
)
