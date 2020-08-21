package atsumori

import "time"

// Benjamin is an Animal Crossing villager
var Benjamin = villager{
	benjaminAstrology,
	benjaminBirthDay,
	benjaminBirthMonth,
	benjaminBubbleColor,
	benjaminCategory,
	benjaminClothes,
	benjaminClothesColors,
	benjaminFlooring,
	benjaminFurniture,
	benjaminGames,
	benjaminGender,
	benjaminInterest,
	benjaminName,
	benjaminNameColor,
	benjaminMusic,
	benjaminPersonality,
	benjaminSpecies,
	benjaminStyle,
	benjaminWallpaper}

var (
	benjaminAstrology     = villagersAstrology{villagerAstrologyLeo}
	benjaminBirthDay      = villagersBirthDay{3}
	benjaminBirthMonth    = villagersBirthMonth{time.August} // 8
	benjaminBubbleColor   = villagersBubbleColor{villagerBubbleColorFFD00D}
	benjaminCategory      = villagersCategory{villagerCategoryA}
	benjaminClothes       = villagersClothes{villagerClothesStripedShirt} // 4145
	benjaminClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorWhite}}
	benjaminFlooring      = villagersFlooring{villagerFlooringCommonFlooring}
	benjaminFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureUprightLocker, villagerFurnitureWoodenBookshelf, villagerFurnitureSchoolChair, villagerFurnitureBasicTeachersDesk, villagerFurnitureCardboardBox, villagerFurnitureSchoolChair, villagerFurnitureChalkboard, villagerFurnitureCuteMusicPlayer, villagerFurnitureWoodenBlockBookshelf, villagerFurnitureSchoolDesk, villagerFurnitureBookStands, villagerFurnitureHomeworkSet, villagerFurnitureBroomAndDustpan, villagerFurnitureSchoolDesk, villagerFurnitureWritingPoster, villagerFurniturePaintingSet}}
	benjaminGames         = villagersGames{[]VillagerGame{}} // TBD
	benjaminGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	benjaminInterest      = villagersInterest{villagerInterestNature}
	benjaminName          = villagersName{villagerNameBenjamin}
	benjaminNameColor     = villagersNameColor{villagerNameColor9B553A}
	benjaminMusic         = villagersMusic{villagerMusicMrKK}
	benjaminPersonality   = villagersPersonality{villagerPersonalityLazy}
	benjaminSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	benjaminStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	benjaminWallpaper     = villagersWallpaper{villagerWallpaperCommonWall}
)
