package atsumori

var Aisle Villager = villager{
	aisleAstrology,
	aisleBirthDay,
	aisleBirthMonth,
	aisleBubbleColor,
	aisleCategory,
	aisleClothes,
	aisleClothesColors,
	aisleFlooring,
	aisleFurniture,
	aisleGames,
	aisleGender,
	aisleInterest,
	aisleName,
	aisleNameColor,
	aisleMusic,
	aislePersonality,
	aisleSpecies,
	aisleStyle,
	aisleWallpaper}

var (
	aisleAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	aisleBirthDay      = villagersBirthDay{}
	aisleBirthMonth    = villagersBirthMonth{}
	aisleBubbleColor   = villagersBubbleColor{}
	aisleCategory      = villagersCategory{}
	aisleClothes       = villagersClothes{}
	aisleClothesColors = villagersClothesColors{}
	aisleFlooring      = villagersFlooring{}
	aisleFurniture     = villagersFurniture{}
	aisleGames         = villagersGames{[]VillagerGame{villagerGameDoubutsuNoMori}}
	aisleGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	aisleInterest      = villagersInterest{}
	aisleName          = villagersName{villagerNameAisle}
	aisleNameColor     = villagersNameColor{}
	aisleMusic         = villagersMusic{}
	aislePersonality   = villagersPersonality{villagerPersonalityLazy}
	aisleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	aisleStyle         = villagersStyle{}
	aisleWallpaper     = villagersWallpaper{}
)
