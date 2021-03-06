package atsumori

// Ace is an Animal Crossing villager.
var Ace Villager = villager{
	aceAstrology,
	aceBirthDay,
	aceBirthMonth,
	aceBubbleColor,
	aceCategory,
	aceClothes,
	aceClothesColors,
	aceFlooring,
	aceFurniture,
	aceGames,
	aceGender,
	aceInterest,
	aceName,
	aceNameColor,
	aceMusic,
	acePersonality,
	aceSpecies,
	aceStyle,
	aceWallpaper}

var (
	aceAstrology     = villagersAstrology{villagerAstrologyPisces}
	aceBirthDay      = villagersBirthDay{}
	aceBirthMonth    = villagersBirthMonth{}
	aceBubbleColor   = villagersBubbleColor{}
	aceCategory      = villagersCategory{}
	aceClothes       = villagersClothes{}
	aceClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	aceFlooring      = villagersFlooring{}
	aceFurniture     = villagersFurniture{[]VillagerFurniture{}}
	aceGames         = villagersGames{[]VillagerGame{}} // TBD
	aceGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	aceInterest      = villagersInterest{}
	aceName          = villagersName{villagerNameAce}
	aceNameColor     = villagersNameColor{}
	aceMusic         = villagersMusic{}
	acePersonality   = villagersPersonality{villagerPersonalityJock}
	aceSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	aceStyle         = villagersStyle{[2]VillagerStyle{}}
	aceWallpaper     = villagersWallpaper{}
)
