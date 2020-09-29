package atsumori

import "time"

// Soleil is an Animal Crossing villager.
var Soleil = villager{
	soleilAstrology,
	soleilBirthDay,
	soleilBirthMonth,
	soleilBubbleColor,
	soleilCategory,
	soleilClothes,
	soleilClothesColors,
	soleilFlooring,
	soleilFurniture,
	soleilGames,
	soleilGender,
	soleilInterest,
	soleilName,
	soleilNameColor,
	soleilMusic,
	soleilPersonality,
	soleilSpecies,
	soleilStyle,
	soleilWallpaper}

var (
	soleilAstrology     = villagersAstrology{villagerAstrologyLeo}
	soleilBirthDay      = villagersBirthDay{9}
	soleilBirthMonth    = villagersBirthMonth{time.August} // 8
	soleilBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	soleilCategory      = villagersCategory{villagerCategoryA}
	soleilClothes       = villagersClothes{} // 3685
	soleilClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorYellow}}
	soleilFlooring      = villagersFlooring{villagerFlooringImperialTile}
	soleilFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAquariusUrn, villagerFurnitureImperialPartition, villagerFurnitureCherryBlossomBranches, villagerFurnitureWhirlpoolBath, villagerFurnitureImperialChest, villagerFurniturePhonograph, villagerFurnitureImperialBed, villagerFurnitureImperialDecorativeShelves, villagerFurnitureImperialLowTable, villagerFurnitureShowerSet, villagerFurnitureIncenseBurner, villagerFurnitureHangingScroll}}
	soleilGames         = villagersGames{[]VillagerGame{}} // TBD
	soleilGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	soleilInterest      = villagersInterest{villagerInterestEducation}
	soleilName          = villagersName{villagerNameSoleil}
	soleilNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	soleilMusic         = villagersMusic{villagerMusicKKOasis}
	soleilPersonality   = villagersPersonality{villagerPersonalitySnooty}
	soleilSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHamster}}
	soleilStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	soleilWallpaper     = villagersWallpaper{villagerWallpaperImperialWall}
)
