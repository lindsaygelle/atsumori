package atsumori

// CarmenMouse is an Animal Crossing villager.
var CarmenMouse Villager = villager{
	carmenMouseAstrology,
	carmenMouseBirthDay,
	carmenMouseBirthMonth,
	carmenMouseBubbleColor,
	carmenMouseCategory,
	carmenMouseClothes,
	carmenMouseClothesColors,
	carmenMouseFlooring,
	carmenMouseFurniture,
	carmenMouseGames,
	carmenMouseGender,
	carmenMouseInterest,
	carmenMouseName,
	carmenMouseNameColor,
	carmenMouseMusic,
	carmenMousePersonality,
	carmenMouseSpecies,
	carmenMouseStyle,
	carmenMouseWallpaper}

var (
	carmenMouseAstrology     = villagersAstrology{villagerAstrologyAries}
	carmenMouseBirthDay      = villagersBirthDay{}
	carmenMouseBirthMonth    = villagersBirthMonth{}
	carmenMouseBubbleColor   = villagersBubbleColor{}
	carmenMouseCategory      = villagersCategory{}
	carmenMouseClothes       = villagersClothes{}
	carmenMouseClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	carmenMouseFlooring      = villagersFlooring{}
	carmenMouseFurniture     = villagersFurniture{[]VillagerFurniture{}}
	carmenMouseGames         = villagersGames{[]VillagerGame{}} // TBD
	carmenMouseGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	carmenMouseInterest      = villagersInterest{}
	carmenMouseName          = villagersName{villagerNameCarmen}
	carmenMouseNameColor     = villagersNameColor{}
	carmenMouseMusic         = villagersMusic{}
	carmenMousePersonality   = villagersPersonality{villagerPersonalitySnooty}
	carmenMouseSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	carmenMouseStyle         = villagersStyle{[2]VillagerStyle{}}
	carmenMouseWallpaper     = villagersWallpaper{}
)
