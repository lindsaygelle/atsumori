package atsumori

import "time"

// CJ is an Animal Crossing villager.
var CJ = villager{
	cjAstrology,
	cjBirthDay,
	cjBirthMonth,
	cjBubbleColor,
	cjCategory,
	cjClothes,
	cjClothesColors,
	cjFlooring,
	cjFurniture,
	cjGames,
	cjGender,
	cjInterest,
	cjName,
	cjNameColor,
	cjMusic,
	cjPersonality,
	cjSpecies,
	cjStyle,
	cjWallpaper}

var (
	cjAstrology     = villagersAstrology{villagerAstrologyAquarius} // villagerAstrology
	cjBirthDay      = villagersBirthDay{7}
	cjBirthMonth    = villagersBirthMonth{time.March} // 2
	cjBubbleColor   = villagersBubbleColor{villagerBubbleColorA99D6E}
	cjCategory      = villagersCategory{}
	cjClothes       = villagersClothes{} //
	cjClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	cjFlooring      = villagersFlooring{}
	cjFurniture     = villagersFurniture{[]VillagerFurniture{}}
	cjGames         = villagersGames{[]VillagerGame{}} // TBD
	cjGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	cjInterest      = villagersInterest{}
	cjName          = villagersName{villagerNameCJ}
	cjNameColor     = villagersNameColor{villagerNameColorF9DB2F}
	cjMusic         = villagersMusic{} //
	cjPersonality   = villagersPersonality{}
	cjSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	cjStyle         = villagersStyle{[2]VillagerStyle{}}
	cjWallpaper     = villagersWallpaper{}
)
