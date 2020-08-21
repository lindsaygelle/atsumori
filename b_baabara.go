package atsumori

import "time"

// Baabara is an Animal Crossing villager
var Baabara = villager{
	baabaraAstrology,
	baabaraBirthDay,
	baabaraBirthMonth,
	baabaraBubbleColor,
	baabaraCategory,
	baabaraClothes,
	baabaraClothesColors,
	baabaraFlooring,
	baabaraFurniture,
	baabaraGames,
	baabaraGender,
	baabaraInterest,
	baabaraName,
	baabaraNameColor,
	baabaraMusic,
	baabaraPersonality,
	baabaraSpecies,
	baabaraStyle,
	baabaraWallpaper}

var (
	baabaraAstrology     = villagersAstrology{villagerAstrologyAries}
	baabaraBirthDay      = villagersBirthDay{28}
	baabaraBirthMonth    = villagersBirthMonth{time.March} // 3
	baabaraBubbleColor   = villagersBubbleColor{villagerBubbleColor8BCDEA}
	baabaraCategory      = villagersCategory{villagerCategoryB}
	baabaraClothes       = villagersClothes{} // 4616 Dress
	baabaraClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlue}}
	baabaraFlooring      = villagersFlooring{villagerFlooringBrownIronParquetFlooring}
	baabaraFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureShowerBooth, villagerFurnitureRedCarpet, villagerFurnitureRedCarpet, villagerFurnitureAntiqueMiniTable, villagerFurnitureWallMountedTV50In, villagerFurnitureFloralSwag, villagerFurniturePhonograph, villagerFurnitureAntiqueChair, villagerFurnitureWhirlpoolBath, villagerFurnitureAntiqueVanity, villagerFurnitureAntiqueConsoleTable, villagerFurnitureFragranceSticks, villagerFurniturePoolsideBed}}
	baabaraGames         = villagersGames{[]VillagerGame{}} // TBD
	baabaraGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	baabaraInterest      = villagersInterest{villagerInterestFashion}
	baabaraName          = villagersName{villagerNameBaabara}
	baabaraNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	baabaraMusic         = villagersMusic{villagerMusicKKCruisin}
	baabaraPersonality   = villagersPersonality{villagerPersonalitySnooty}
	baabaraSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSheep}}
	baabaraStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	baabaraWallpaper     = villagersWallpaper{villagerWallpaperStormyNightWall}
)
