package atsumori

import "time"

// Ed is an Animal Crossing villager.
var Ed = villager{
	edAstrology,
	edBirthDay,
	edBirthMonth,
	edBubbleColor,
	edCategory,
	edClothes,
	edClothesColors,
	edFlooring,
	edFurniture,
	edGames,
	edGender,
	edInterest,
	edName,
	edNameColor,
	edMusic,
	edPersonality,
	edSpecies,
	edStyle,
	edWallpaper}

var (
	edAstrology     = villagersAstrology{villagerAstrologyVirgo}
	edBirthDay      = villagersBirthDay{16}
	edBirthMonth    = villagersBirthMonth{time.September} // 9
	edBubbleColor   = villagersBubbleColor{villagerBubbleColor55CFFF}
	edCategory      = villagersCategory{villagerCategoryB}
	edClothes       = villagersClothes{villagerClothesCollarlessShirt} // 3690
	edClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	edFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	edFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureShowerBooth, villagerFurnitureAntiqueBed, villagerFurniturePendulumClock, villagerFurnitureAntiqueConsoleTable, villagerFurnitureAnthuriumPlant, villagerFurnitureRetroRadiator, villagerFurnitureFireplace, villagerFurnitureEdsPhoto, villagerFurnitureRetroStereo}}
	edGames         = villagersGames{[]VillagerGame{}} // TBD
	edGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	edInterest      = villagersInterest{villagerInterestNature}
	edName          = villagersName{villagerNameEd}
	edNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	edMusic         = villagersMusic{villagerMusicKKBallad}
	edPersonality   = villagersPersonality{villagerPersonalitySmug}
	edSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	edStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	edWallpaper     = villagersWallpaper{villagerWallpaperBlueDelicateBloomsWall}
)
