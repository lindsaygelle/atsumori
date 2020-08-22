package atsumori

import "time"

// Doc is an Animal Crossing villager
var Doc = villager{
	docAstrology,
	docBirthDay,
	docBirthMonth,
	docBubbleColor,
	docCategory,
	docClothes,
	docClothesColors,
	docFlooring,
	docFurniture,
	docGames,
	docGender,
	docInterest,
	docName,
	docNameColor,
	docMusic,
	docPersonality,
	docSpecies,
	docStyle,
	docWallpaper}

var (
	docAstrology     = villagersAstrology{villagerAstrologyPisces}
	docBirthDay      = villagersBirthDay{16}
	docBirthMonth    = villagersBirthMonth{time.March} // 3
	docBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	docCategory      = villagersCategory{villagerCategoryB}
	docClothes       = villagersClothes{villagerClothesFlannelShirt} // 3225
	docClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorBeige}}
	docFlooring      = villagersFlooring{villagerFlooringDarkBlockFlooring}
	docFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSkeleton, villagerFurnitureDenChair, villagerFurnitureWoodenBookshelf, villagerFurnitureWoodenBookshelf, villagerFurnitureChalkboard, villagerFurnitureAntiqueConsoleTable, villagerFurniturePendulumClock, villagerFurniturePhonograph, villagerFurnitureWritingPoster, villagerFurnitureBathroomTowelRack, villagerFurnitureToiletCleaningSet, villagerFurniturePlainSink, villagerFurnitureDenDesk, villagerFurnitureRedCarpet, villagerFurnitureIvorySimpleBathMat, villagerFurnitureLabExperimentsSet, villagerFurnitureToilet}}
	docGames         = villagersGames{[]VillagerGame{}} // TBD
	docGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	docInterest      = villagersInterest{villagerInterestEducation}
	docName          = villagersName{villagerNameDoc}
	docNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	docMusic         = villagersMusic{villagerMusicPondering}
	docPersonality   = villagersPersonality{villagerPersonalityLazy}
	docSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	docStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	docWallpaper     = villagersWallpaper{villagerWallpaperClassicLibraryWall}
)
