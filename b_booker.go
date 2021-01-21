package atsumori

import "time"

// Booker is an Animal Crossing villager.
var Booker = villager{
	bookerAstrology,
	bookerBirthDay,
	bookerBirthMonth,
	bookerBubbleColor,
	bookerCategory,
	bookerClothes,
	bookerClothesColors,
	bookerFlooring,
	bookerFurniture,
	bookerGames,
	bookerGender,
	bookerInterest,
	bookerName,
	bookerNameColor,
	bookerMusic,
	bookerPersonality,
	bookerSpecies,
	bookerStyle,
	bookerWallpaper}

var (
	bookerAstrology     = villagersAstrology{villagerAstrologyTaurus} // villagerAstrology
	bookerBirthDay      = villagersBirthDay{23}
	bookerBirthMonth    = villagersBirthMonth{time.April} // 2
	bookerBubbleColor   = villagersBubbleColor{villagerBubbleColorC67022}
	bookerCategory      = villagersCategory{}
	bookerClothes       = villagersClothes{} //
	bookerClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	bookerFlooring      = villagersFlooring{}
	bookerFurniture     = villagersFurniture{[]VillagerFurniture{}}
	bookerGames         = villagersGames{[]VillagerGame{}} // TBD
	bookerGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	bookerInterest      = villagersInterest{}
	bookerName          = villagersName{villagerNameBooker}
	bookerNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bookerMusic         = villagersMusic{} //
	bookerPersonality   = villagersPersonality{}
	bookerSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	bookerStyle         = villagersStyle{[2]VillagerStyle{}}
	bookerWallpaper     = villagersWallpaper{}
)
