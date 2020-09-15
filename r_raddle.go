package atsumori

import "time"

// Raddle is an Animal Crossing villager.
var Raddle = villager{
	raddleAstrology,
	raddleBirthDay,
	raddleBirthMonth,
	raddleBubbleColor,
	raddleCategory,
	raddleClothes,
	raddleClothesColors,
	raddleFlooring,
	raddleFurniture,
	raddleGames,
	raddleGender,
	raddleInterest,
	raddleName,
	raddleNameColor,
	raddleMusic,
	raddlePersonality,
	raddleSpecies,
	raddleStyle,
	raddleWallpaper}

var (
	raddleAstrology     = villagersAstrology{villagerAstrologyGemini}
	raddleBirthDay      = villagersBirthDay{6}
	raddleBirthMonth    = villagersBirthMonth{time.June} // 6
	raddleBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	raddleCategory      = villagersCategory{villagerCategoryA}
	raddleClothes       = villagersClothes{} // 7794
	raddleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorWhite, villagerClothesColorGray}}
	raddleFlooring      = villagersFlooring{villagerFlooringGreenRubberFlooring}
	raddleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSkeleton, villagerFurnitureUtilitySink, villagerFurnitureBathroomTowelRack, villagerFurnitureFoldingChair, villagerFurnitureToolCart, villagerFurnitureTapeDeck, villagerFurnitureFoldingChair, villagerFurnitureUprightLocker, villagerFurnitureAnatomicalModel, villagerFurnitureLectureHallDesk, villagerFurnitureHomeworkSet, villagerFurnitureLectureHallDesk, villagerFurnitureLabExperimentsSet, villagerFurnitureWallMountedTV50In, villagerFurnitureChalkboard}}
	raddleGames         = villagersGames{[]VillagerGame{}} // TBD
	raddleGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	raddleInterest      = villagersInterest{villagerInterestNature}
	raddleName          = villagersName{villagerNameRaddle}
	raddleNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	raddleMusic         = villagersMusic{villagerMusicMrKK}
	raddlePersonality   = villagersPersonality{villagerPersonalityLazy}
	raddleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	raddleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleSimple}}
	raddleWallpaper     = villagersWallpaper{villagerWallpaperWhiteSimpleClothWall}
)
