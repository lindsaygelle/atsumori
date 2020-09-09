package atsumori

import "time"

// Olive is an Animal Crossing villager.
var Olive = villager{
	oliveAstrology,
	oliveBirthDay,
	oliveBirthMonth,
	oliveBubbleColor,
	oliveCategory,
	oliveClothes,
	oliveClothesColors,
	oliveFlooring,
	oliveFurniture,
	oliveGames,
	oliveGender,
	oliveInterest,
	oliveName,
	oliveNameColor,
	oliveMusic,
	olivePersonality,
	oliveSpecies,
	oliveStyle,
	oliveWallpaper}

var (
	oliveAstrology     = villagersAstrology{villagerAstrologyCancer}
	oliveBirthDay      = villagersBirthDay{12}
	oliveBirthMonth    = villagersBirthMonth{time.July} // 7
	oliveBubbleColor   = villagersBubbleColor{villagerBubbleColorBFBFBF}
	oliveCategory      = villagersCategory{villagerCategoryA}
	oliveClothes       = villagersClothes{} // 3128
	oliveClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorYellow}}
	oliveFlooring      = villagersFlooring{villagerFlooringMintDotFlooring}
	oliveFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenSimpleBed, villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenWasteBin, villagerFurnitureCorkboard, villagerFurnitureWoodenLowTable, villagerFurnitureWritingPoster, villagerFurnitureWoodenTableMirror, villagerFurnitureFluffyRug, villagerFurnitureWoodenEndTable, villagerFurnitureWoodenChest, villagerFurnitureWritingChair, villagerFurnitureWritingDesk, villagerFurnitureWoodenMiniTable, villagerFurnitureOldFashionedAlarmClock, villagerFurnitureBookStands, villagerFurniturePortableRadio, villagerFurnitureGlobe}}
	oliveGames         = villagersGames{[]VillagerGame{}} // TBD
	oliveGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	oliveInterest      = villagersInterest{villagerInterestNature}
	oliveName          = villagersName{villagerNameOlive}
	oliveNameColor     = villagersNameColor{villagerNameColor5E5E5E}
	oliveMusic         = villagersMusic{villagerMusicNeapolitan}
	olivePersonality   = villagersPersonality{villagerPersonalityNormal}
	oliveSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	oliveStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	oliveWallpaper     = villagersWallpaper{villagerWallpaperWhiteSimpleClothWall}
)
