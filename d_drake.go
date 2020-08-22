package atsumori

import "time"

// Drake is an Animal Crossing villager
var Drake = villager{
	drakeAstrology,
	drakeBirthDay,
	drakeBirthMonth,
	drakeBubbleColor,
	drakeCategory,
	drakeClothes,
	drakeClothesColors,
	drakeFlooring,
	drakeFurniture,
	drakeGames,
	drakeGender,
	drakeInterest,
	drakeName,
	drakeNameColor,
	drakeMusic,
	drakePersonality,
	drakeSpecies,
	drakeStyle,
	drakeWallpaper}

var (
	drakeAstrology     = villagersAstrology{villagerAstrologyCancer}
	drakeBirthDay      = villagersBirthDay{25}
	drakeBirthMonth    = villagersBirthMonth{time.June} // 6
	drakeBubbleColor   = villagersBubbleColor{villagerBubbleColor459ABA}
	drakeCategory      = villagersCategory{villagerCategoryB}
	drakeClothes       = villagersClothes{villagerClothesYodelSweater} // 3630
	drakeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorRed}}
	drakeFlooring      = villagersFlooring{villagerFlooringCommonFlooring}
	drakeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureTrainSet, villagerFurnitureWoodenSimpleBed, villagerFurnitureMiniDIYWorkbench, villagerFurnitureShadedFloorLamp, villagerFurnitureLogExtraLongSofa, villagerFurnitureWoodenChest, villagerFurniturePortableRecordPlayer, villagerFurnitureMacrameTapestry, villagerFurnitureBluePersianRug, villagerFurnitureWoodenEndTable, villagerFurnitureFireplace, villagerFurnitureDigitalAlarmClock}}
	drakeGames         = villagersGames{[]VillagerGame{}} // TBD
	drakeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	drakeInterest      = villagersInterest{villagerInterestPlay}
	drakeName          = villagersName{villagerNameDrake}
	drakeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	drakeMusic         = villagersMusic{villagerMusicKKBlues}
	drakePersonality   = villagersPersonality{villagerPersonalityLazy}
	drakeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	drakeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	drakeWallpaper     = villagersWallpaper{villagerWallpaperCabinWall}
)
