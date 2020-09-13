package atsumori

import "time"

// Pate is an Animal Crossing villager.
var Pate = villager{
	pateAstrology,
	pateBirthDay,
	pateBirthMonth,
	pateBubbleColor,
	pateCategory,
	pateClothes,
	pateClothesColors,
	pateFlooring,
	pateFurniture,
	pateGames,
	pateGender,
	pateInterest,
	pateName,
	pateNameColor,
	pateMusic,
	patePersonality,
	pateSpecies,
	pateStyle,
	pateWallpaper}

var (
	pateAstrology     = villagersAstrology{villagerAstrologyPisces}
	pateBirthDay      = villagersBirthDay{23}
	pateBirthMonth    = villagersBirthMonth{time.February} // 2
	pateBubbleColor   = villagersBubbleColor{villagerBubbleColor8BCDEA}
	pateCategory      = villagersCategory{villagerCategoryB}
	pateClothes       = villagersClothes{} // 4432
	pateClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorWhite}}
	pateFlooring      = villagersFlooring{villagerFlooringHexagonalFloralFlooring}
	pateFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHangingTerrarium, villagerFurnitureShowerSet, villagerFurnitureRattanBed, villagerFurnitureRattanWardrobe, villagerFurnitureRattanLowTable, villagerFurnitureSimpleGreenBathMat, villagerFurnitureRattanVanity, villagerFurnitureRattanEndTable, villagerFurnitureRattanStool, villagerFurnitureRattanArmchair, villagerFurnitureLongBathtub, villagerFurnitureRattanTowelBasket, villagerFurnitureIronShelf, villagerFurniturePortableRecordPlayer}}
	pateGames         = villagersGames{[]VillagerGame{}} // TBD
	pateGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	pateInterest      = villagersInterest{villagerInterestFashion}
	pateName          = villagersName{villagerNamePate}
	pateNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	pateMusic         = villagersMusic{villagerMusicKKSoul}
	patePersonality   = villagersPersonality{villagerPersonalityPeppy}
	pateSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	pateStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	pateWallpaper     = villagersWallpaper{villagerWallpaperBlueSubwayTileWall}
)
