package atsumori

import "time"

// Bones is an Animal Crossing villager
var Bones = villager{
	bonesAstrology,
	bonesBirthDay,
	bonesBirthMonth,
	bonesBubbleColor,
	bonesCategory,
	bonesClothes,
	bonesClothesColors,
	bonesFlooring,
	bonesFurniture,
	bonesGames,
	bonesGender,
	bonesInterest,
	bonesName,
	bonesNameColor,
	bonesMusic,
	bonesPersonality,
	bonesSpecies,
	bonesStyle,
	bonesWallpaper}

var (
	bonesAstrology     = villagersAstrology{villagerAstrologyLeo}
	bonesBirthDay      = villagersBirthDay{4}
	bonesBirthMonth    = villagersBirthMonth{time.August} // 8
	bonesBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	bonesCategory      = villagersCategory{villagerCategoryB}
	bonesClothes       = villagersClothes{} // 7851
	bonesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorBrown}}
	bonesFlooring      = villagersFlooring{villagerFlooringLightWoodPatternFlooring}
	bonesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBookshelf, villagerFurnitureSchoolChair, villagerFurnitureSchoolChair, villagerFurnitureCardboardBox, villagerFurnitureTapeDeck, villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenBookshelf, villagerFurnitureModernOfficeChair, villagerFurnitureLectureHallDesk, villagerFurnitureHomeworkSet, villagerFurnitureRubberMudMat, villagerFurnitureLectureHallDesk, villagerFurnitureStackOfBooks, villagerFurnitureCardboardBox, villagerFurnitureStackOfBooks}}
	bonesGames         = villagersGames{[]VillagerGame{}} // TBD
	bonesGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	bonesInterest      = villagersInterest{villagerInterestPlay}
	bonesName          = villagersName{villagerNameBones}
	bonesNameColor     = villagersNameColor{villagerNameColor848484}
	bonesMusic         = villagersMusic{} // K.K. Étude
	bonesPersonality   = villagersPersonality{villagerPersonalityLazy}
	bonesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	bonesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	bonesWallpaper     = villagersWallpaper{villagerWallpaperClassicLibraryWall}
)
