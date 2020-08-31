package atsumori

import "time"

// Leopold is an Animal Crossing villager.
var Leopold = villager{
	leopoldAstrology,
	leopoldBirthDay,
	leopoldBirthMonth,
	leopoldBubbleColor,
	leopoldCategory,
	leopoldClothes,
	leopoldClothesColors,
	leopoldFlooring,
	leopoldFurniture,
	leopoldGames,
	leopoldGender,
	leopoldInterest,
	leopoldName,
	leopoldNameColor,
	leopoldMusic,
	leopoldPersonality,
	leopoldSpecies,
	leopoldStyle,
	leopoldWallpaper}

var (
	leopoldAstrology     = villagersAstrology{villagerAstrologyLeo}
	leopoldBirthDay      = villagersBirthDay{14}
	leopoldBirthMonth    = villagersBirthMonth{time.August} // 8
	leopoldBubbleColor   = villagersBubbleColor{villagerBubbleColorE4B887}
	leopoldCategory      = villagersCategory{villagerCategoryA}
	leopoldClothes       = villagersClothes{villagerClothesTennisSweater} // 3599
	leopoldClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorGreen}}
	leopoldFlooring      = villagersFlooring{villagerFlooringNaturalBlockFlooring}
	leopoldFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSkeleton, villagerFurnitureUprightLocker, villagerFurnitureWritingPoster, villagerFurnitureLectureHallDesk, villagerFurnitureHomeworkSet, villagerFurnitureLectureHallDesk, villagerFurnitureFormalPaper, villagerFurnitureBookStands, villagerFurnitureBasicTeachersDesk, villagerFurnitureWallMountedTV50In, villagerFurnitureDocumentStack, villagerFurnitureLectureHallBench, villagerFurnitureUprightLocker}}
	leopoldGames         = villagersGames{[]VillagerGame{}} // TBD
	leopoldGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	leopoldInterest      = villagersInterest{villagerInterestEducation}
	leopoldName          = villagersName{villagerNameLeopold}
	leopoldNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	leopoldMusic         = villagersMusic{villagerMusicMrKK}
	leopoldPersonality   = villagersPersonality{villagerPersonalitySmug}
	leopoldSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesLion}}
	leopoldStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	leopoldWallpaper     = villagersWallpaper{villagerWallpaperGreenPaintedWoodWall}
)
