package atsumori

import "time"

// Copper is an Animal Crossing villager.
var Copper = villager{
	copperAstrology,
	copperBirthDay,
	copperBirthMonth,
	copperBubbleColor,
	copperCategory,
	copperClothes,
	copperClothesColors,
	copperFlooring,
	copperFurniture,
	copperGames,
	copperGender,
	copperInterest,
	copperName,
	copperNameColor,
	copperMusic,
	copperPersonality,
	copperSpecies,
	copperStyle,
	copperWallpaper}

var (
	copperAstrology     = villagersAstrology{villagerAstrologyCancer} // villagerAstrology
	copperBirthDay      = villagersBirthDay{28}
	copperBirthMonth    = villagersBirthMonth{time.June} // 2
	copperBubbleColor   = villagersBubbleColor{villagerBubbleColorE69111}
	copperCategory      = villagersCategory{}
	copperClothes       = villagersClothes{} //
	copperClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	copperFlooring      = villagersFlooring{}
	copperFurniture     = villagersFurniture{[]VillagerFurniture{}}
	copperGames         = villagersGames{[]VillagerGame{}} // TBD
	copperGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	copperInterest      = villagersInterest{}
	copperName          = villagersName{villagerNameCopper}
	copperNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	copperMusic         = villagersMusic{} //
	copperPersonality   = villagersPersonality{}
	copperSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDog}}
	copperStyle         = villagersStyle{[2]VillagerStyle{}}
	copperWallpaper     = villagersWallpaper{}
)
