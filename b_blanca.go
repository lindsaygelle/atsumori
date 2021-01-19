package atsumori

import "time"

// Blanca is an Animal Crossing villager.
var Blanca = villager{
	blancaAstrology,
	blancaBirthDay,
	blancaBirthMonth,
	blancaBubbleColor,
	blancaCategory,
	blancaClothes,
	blancaClothesColors,
	blancaFlooring,
	blancaFurniture,
	blancaGames,
	blancaGender,
	blancaInterest,
	blancaName,
	blancaNameColor,
	blancaMusic,
	blancaPersonality,
	blancaSpecies,
	blancaStyle,
	blancaWallpaper}

var (
	blancaAstrology     = villagersAstrology{villagerAstrologyAquarius} // villagerAstrology
	blancaBirthDay      = villagersBirthDay{8}
	blancaBirthMonth    = villagersBirthMonth{time.February} // 2
	blancaBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	blancaCategory      = villagersCategory{}
	blancaClothes       = villagersClothes{} //
	blancaClothesColors = villagersClothesColors{[2]VillagerClothesColor{}}
	blancaFlooring      = villagersFlooring{}
	blancaFurniture     = villagersFurniture{[]VillagerFurniture{}}
	blancaGames         = villagersGames{[]VillagerGame{}} // TBD
	blancaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale, villagerGenderMale}}
	blancaInterest      = villagersInterest{}
	blancaName          = villagersName{villagerNameBlanca}
	blancaNameColor     = villagersNameColor{villagerNameColorFE8885}
	blancaMusic         = villagersMusic{} //
	blancaPersonality   = villagersPersonality{}
	blancaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	blancaStyle         = villagersStyle{[2]VillagerStyle{}}
	blancaWallpaper     = villagersWallpaper{}
)
