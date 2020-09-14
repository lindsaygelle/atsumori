package atsumori

import "time"

// Phil is an Animal Crossing villager.
var Phil = villager{
	philAstrology,
	philBirthDay,
	philBirthMonth,
	philBubbleColor,
	philCategory,
	philClothes,
	philClothesColors,
	philFlooring,
	philFurniture,
	philGames,
	philGender,
	philInterest,
	philName,
	philNameColor,
	philMusic,
	philPersonality,
	philSpecies,
	philStyle,
	philWallpaper}

var (
	philAstrology     = villagersAstrology{villagerAstrologySagittarius}
	philBirthDay      = villagersBirthDay{27}
	philBirthMonth    = villagersBirthMonth{time.November} // 11
	philBubbleColor   = villagersBubbleColor{villagerBubbleColor7352E8}
	philCategory      = villagersCategory{villagerCategoryA}
	philClothes       = villagersClothes{villagerClothesFischerhemd} // 5658
	philClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorRed}}
	philFlooring      = villagersFlooring{villagerFlooringSimpleRedFlooring}
	philFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureLongBathtub, villagerFurnitureAntiqueConsoleTable, villagerFurniturePhonograph, villagerFurnitureRattanBed, villagerFurnitureTanklessToilet, villagerFurnitureRattanWardrobe, villagerFurnitureRattanLowTable, villagerFurnitureTraditionalTeaSet, villagerFurnitureImperialPartition, villagerFurnitureDeluxeWasher, villagerFurnitureShowerSet, villagerFurnitureBathroomTowelRack, villagerFurnitureSimpleNavyBathMat, villagerFurnitureRattanTowelBasket}}
	philGames         = villagersGames{[]VillagerGame{}} // TBD
	philGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	philInterest      = villagersInterest{villagerInterestMusic}
	philName          = villagersName{villagerNamePhil}
	philNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	philMusic         = villagersMusic{villagerMusicKKMoody}
	philPersonality   = villagersPersonality{villagerPersonalitySmug}
	philSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	philStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	philWallpaper     = villagersWallpaper{villagerWallpaperRedArtDecoWall}
)
