package atsumori

import "time"

// Miranda is an Animal Crossing villager.
var Miranda = villager{
	mirandaAstrology,
	mirandaBirthDay,
	mirandaBirthMonth,
	mirandaBubbleColor,
	mirandaCategory,
	mirandaClothes,
	mirandaClothesColors,
	mirandaFlooring,
	mirandaFurniture,
	mirandaGames,
	mirandaGender,
	mirandaInterest,
	mirandaName,
	mirandaNameColor,
	mirandaMusic,
	mirandaPersonality,
	mirandaSpecies,
	mirandaStyle,
	mirandaWallpaper}

var (
	mirandaAstrology     = villagersAstrology{villagerAstrologyTaurus}
	mirandaBirthDay      = villagersBirthDay{23}
	mirandaBirthMonth    = villagersBirthMonth{time.April} // 4
	mirandaBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	mirandaCategory      = villagersCategory{villagerCategoryB}
	mirandaClothes       = villagersClothes{} // 3387
	mirandaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorPurple}}
	mirandaFlooring      = villagersFlooring{villagerFlooringBrownIronParquetFlooring}
	mirandaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWhirlpoolBath, villagerFurnitureAntiqueBed, villagerFurnitureAntiqueVanity, villagerFurnitureDoubleSofa, villagerFurnitureHarp, villagerFurnitureAntiqueMiniTable, villagerFurniturePhonograph, villagerFurnitureImperialLowTable, villagerFurnitureMonstera, villagerFurniturePendulumClock}}
	mirandaGames         = villagersGames{[]VillagerGame{}} // TBD
	mirandaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	mirandaInterest      = villagersInterest{villagerInterestFashion}
	mirandaName          = villagersName{villagerNameMiranda}
	mirandaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	mirandaMusic         = villagersMusic{villagerMusicKKSonata}
	mirandaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	mirandaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	mirandaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	mirandaWallpaper     = villagersWallpaper{villagerWallpaperArchedWindowWall}
)
