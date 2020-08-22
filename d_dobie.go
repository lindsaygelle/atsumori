package atsumori

import "time"

// Dobie is an Animal Crossing villager
var Dobie = villager{
	dobieAstrology,
	dobieBirthDay,
	dobieBirthMonth,
	dobieBubbleColor,
	dobieCategory,
	dobieClothes,
	dobieClothesColors,
	dobieFlooring,
	dobieFurniture,
	dobieGames,
	dobieGender,
	dobieInterest,
	dobieName,
	dobieNameColor,
	dobieMusic,
	dobiePersonality,
	dobieSpecies,
	dobieStyle,
	dobieWallpaper}

var (
	dobieAstrology     = villagersAstrology{villagerAstrologyAquarius}
	dobieBirthDay      = villagersBirthDay{17}
	dobieBirthMonth    = villagersBirthMonth{time.February} // 2
	dobieBubbleColor   = villagersBubbleColor{villagerBubbleColor798040}
	dobieCategory      = villagersCategory{villagerCategoryA}
	dobieClothes       = villagersClothes{villagerClothesFuzzyVest} // 4212
	dobieClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorBeige}}
	dobieFlooring      = villagersFlooring{villagerFlooringDarkHerringboneFlooring}
	dobieFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDenChair, villagerFurnitureIronHangerStand, villagerFurnitureStackOfBooks, villagerFurnitureYellowPersianRug, villagerFurnitureAntiqueConsoleTable, villagerFurniturePendulumClock, villagerFurnitureRotaryPhone, villagerFurnitureScatteredPapers, villagerFurnitureWhiteboard, villagerFurnitureCardboardBox, villagerFurnitureDocumentStack, villagerFurnitureCardboardBox, villagerFurnitureAntiqueMiniTable, villagerFurnitureDenDesk, villagerFurniturePhonograph, villagerFurnitureStackOfBooks, villagerFurnitureEssaySet, villagerFurnitureTypewriter}}
	dobieGames         = villagersGames{[]VillagerGame{}} // TBD
	dobieGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	dobieInterest      = villagersInterest{villagerInterestNature}
	dobieName          = villagersName{villagerNameDobie}
	dobieNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	dobieMusic         = villagersMusic{villagerMusicAgentKK}
	dobiePersonality   = villagersPersonality{villagerPersonalityCranky}
	dobieSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesWolf}}
	dobieStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	dobieWallpaper     = villagersWallpaper{villagerWallpaperClassicLibraryWall}
)
