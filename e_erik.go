package atsumori

import "time"

// Erik is an Animal Crossing villager
var Erik = villager{
	erikAstrology,
	erikBirthDay,
	erikBirthMonth,
	erikBubbleColor,
	erikCategory,
	erikClothes,
	erikClothesColors,
	erikFlooring,
	erikFurniture,
	erikGames,
	erikGender,
	erikInterest,
	erikName,
	erikNameColor,
	erikMusic,
	erikPersonality,
	erikSpecies,
	erikStyle,
	erikWallpaper}

var (
	erikAstrology     = villagersAstrology{villagerAstrologyLeo}
	erikBirthDay      = villagersBirthDay{27}
	erikBirthMonth    = villagersBirthMonth{time.July} // 7
	erikBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	erikCategory      = villagersCategory{villagerCategoryA}
	erikClothes       = villagersClothes{villagerClothesYodelSweater} // 3630
	erikClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorRed}}
	erikFlooring      = villagersFlooring{villagerFlooringSkiSlopeFlooring}
	erikFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSleigh, villagerFurnitureFrozenTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureFrozenBed, villagerFurnitureFrozenCounter, villagerFurnitureFrozenPartition, villagerFurnitureFrozenChair}}
	erikGames         = villagersGames{[]VillagerGame{}} // TBD
	erikGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	erikInterest      = villagersInterest{villagerInterestNature}
	erikName          = villagersName{villagerNameErik}
	erikNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	erikMusic         = villagersMusic{villagerMusicDrivin}
	erikPersonality   = villagersPersonality{villagerPersonalityLazy}
	erikSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	erikStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	erikWallpaper     = villagersWallpaper{villagerWallpaperSkiSlopeWall}
)
