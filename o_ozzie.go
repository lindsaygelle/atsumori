package atsumori

import "time"

// Ozzie is an Animal Crossing villager.
var Ozzie = villager{
	ozzieAstrology,
	ozzieBirthDay,
	ozzieBirthMonth,
	ozzieBubbleColor,
	ozzieCategory,
	ozzieClothes,
	ozzieClothesColors,
	ozzieFlooring,
	ozzieFurniture,
	ozzieGames,
	ozzieGender,
	ozzieInterest,
	ozzieName,
	ozzieNameColor,
	ozzieMusic,
	ozziePersonality,
	ozzieSpecies,
	ozzieStyle,
	ozzieWallpaper}

var (
	ozzieAstrology     = villagersAstrology{villagerAstrologyTaurus}
	ozzieBirthDay      = villagersBirthDay{7}
	ozzieBirthMonth    = villagersBirthMonth{time.May} // 5
	ozzieBubbleColor   = villagersBubbleColor{villagerBubbleColorBFBFBF}
	ozzieCategory      = villagersCategory{villagerCategoryB}
	ozzieClothes       = villagersClothes{villagerClothesEnergeticSweater} // 5262
	ozzieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorOrange}}
	ozzieFlooring      = villagersFlooring{villagerFlooringCorkFlooring}
	ozzieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureChalkboard, villagerFurnitureSkeleton, villagerFurnitureIronwoodDIYWorkbench, villagerFurnitureUprightLocker, villagerFurnitureWritingPoster, villagerFurnitureWallMountedToolBoard, villagerFurnitureUprightVacuum, villagerFurnitureSchoolChair, villagerFurnitureWhiteSimpleSmallMat, villagerFurnitureSchoolChair, villagerFurnitureWhiteSimpleSmallMat, villagerFurnitureSchoolDesk, villagerFurnitureCuteMusicPlayer, villagerFurnitureSimplesmallAvocadoMat, villagerFurnitureToolCart, villagerFurnitureWoodenToolbox, villagerFurnitureSimplesmallAvocadoMat, villagerFurnitureSchoolDesk, villagerFurnitureElectronicsKit, villagerFurnitureSchoolDesk, villagerFurnitureSchoolDesk, villagerFurnitureHamsterCage, villagerFurnitureWoodenToolbox}}
	ozzieGames         = villagersGames{[]VillagerGame{}} // TBD
	ozzieGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	ozzieInterest      = villagersInterest{villagerInterestPlay}
	ozzieName          = villagersName{villagerNameOzzie}
	ozzieNameColor     = villagersNameColor{villagerNameColor5E5E5E}
	ozzieMusic         = villagersMusic{villagerMusicMrKK}
	ozziePersonality   = villagersPersonality{villagerPersonalityLazy}
	ozzieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	ozzieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	ozzieWallpaper     = villagersWallpaper{villagerWallpaperWhitePerforatedBoardWall}
)
