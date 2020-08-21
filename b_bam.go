package atsumori

import "time"

// Bam is an Animal Crossing villager
var Bam = villager{
	bamAstrology,
	bamBirthDay,
	bamBirthMonth,
	bamBubbleColor,
	bamCategory,
	bamClothes,
	bamClothesColors,
	bamFlooring,
	bamFurniture,
	bamGames,
	bamGender,
	bamInterest,
	bamName,
	bamNameColor,
	bamMusic,
	bamPersonality,
	bamSpecies,
	bamStyle,
	bamWallpaper}

var (
	bamAstrology     = villagersAstrology{villagerAstrologyScorpio}
	bamBirthDay      = villagersBirthDay{7}
	bamBirthMonth    = villagersBirthMonth{time.November} // 11
	bamBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	bamCategory      = villagersCategory{villagerCategoryA}
	bamClothes       = villagersClothes{villagerClothesAthleticJacket} // 7198
	bamClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBrown}}
	bamFlooring      = villagersFlooring{villagerFlooringBlueRubberFlooring}
	bamFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureUprightLocker, villagerFurnitureFoldingChair, villagerFurnitureFoldingChair, villagerFurnitureCuteMusicPlayer, villagerFurnitureUtilitySink, villagerFurnitureBlueCorner, villagerFurnitureTennisTable, villagerFurnitureOutdoorTable, villagerFurnitureStackOfBooks, villagerFurnitureWhiteboard, villagerFurnitureWallFan}}
	bamGames         = villagersGames{[]VillagerGame{}} // TBD
	bamGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	bamInterest      = villagersInterest{villagerInterestPlay}
	bamName          = villagersName{villagerNameBam}
	bamNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bamMusic         = villagersMusic{villagerMusicKKBlues}
	bamPersonality   = villagersPersonality{villagerPersonalityJock}
	bamSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	bamStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	bamWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
