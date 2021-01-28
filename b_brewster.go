package atsumori

import "time"

// Brewster is an Animal Crossing villager.
var Brewster = villager{
	brewsterAstrology,
	brewsterBirthDay,
	brewsterBirthMonth,
	brewsterBubbleColor,
	brewsterCategory,
	brewsterClothes,
	brewsterClothesColors,
	brewsterFlooring,
	brewsterFurniture,
	brewsterGames,
	brewsterGender,
	brewsterInterest,
	brewsterName,
	brewsterNameColor,
	brewsterMusic,
	brewsterPersonality,
	brewsterSpecies,
	brewsterStyle,
	brewsterWallpaper}

var (
	brewsterAstrology     = villagersAstrology{villagerAstrologyLibra} // villagerAstrology
	brewsterBirthDay      = villagersBirthDay{15}
	brewsterBirthMonth    = villagersBirthMonth{time.October} // 2
	brewsterBubbleColor   = villagersBubbleColor{villagerBubbleColor689A8B}
	brewsterCategory      = villagersCategory{}
	brewsterClothes       = villagersClothes{} //
	brewsterClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	brewsterFlooring      = villagersFlooring{}
	brewsterFurniture     = villagersFurniture{[]VillagerFurniture{}}
	brewsterGames         = villagersGames{[]VillagerGame{}} // TBD
	brewsterGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	brewsterInterest      = villagersInterest{}
	brewsterName          = villagersName{villagerNameBrewster}
	brewsterNameColor     = villagersNameColor{villagerNameColor333333}
	brewsterMusic         = villagersMusic{} //
	brewsterPersonality   = villagersPersonality{}
	brewsterSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPigeon}}
	brewsterStyle         = villagersStyle{[2]VillagerStyle{}}
	brewsterWallpaper     = villagersWallpaper{}
)
