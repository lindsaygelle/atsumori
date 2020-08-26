package atsumori

import "time"

// Julian is an Animal Crossing villager
var Julian = villager{
	julianAstrology,
	julianBirthDay,
	julianBirthMonth,
	julianBubbleColor,
	julianCategory,
	julianClothes,
	julianClothesColors,
	julianFlooring,
	julianFurniture,
	julianGames,
	julianGender,
	julianInterest,
	julianName,
	julianNameColor,
	julianMusic,
	julianPersonality,
	julianSpecies,
	julianStyle,
	julianWallpaper}

var (
	julianAstrology     = villagersAstrology{villagerAstrologyPisces}
	julianBirthDay      = villagersBirthDay{15}
	julianBirthMonth    = villagersBirthMonth{time.March} // 3
	julianBubbleColor   = villagersBubbleColor{villagerBubbleColor3FD8E0}
	julianCategory      = villagersCategory{villagerCategoryA}
	julianClothes       = villagersClothes{villagerClothesSpaceParka} // 4418
	julianClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorBlue}}
	julianFlooring      = villagersFlooring{villagerFlooringCloudFlooring}
	julianFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAquariusUrn, villagerFurnitureCancerTable, villagerFurnitureStarryGarland, villagerFurnitureLibraScale, villagerFurnitureVirgoHarp, villagerFurnitureGeminiCloset, villagerFurnitureFortuneTellingSet, villagerFurnitureScorpioLamp, villagerFurnitureCrescentMoonChair, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureAriesRockingChair, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureLeoSculpture, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureCapricornOrnament, villagerFurnitureTaurusBathtub, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland, villagerFurniturePiscesLamp, villagerFurnitureSagittariusArrow, villagerFurnitureStarryGarland}}
	julianGames         = villagersGames{[]VillagerGame{}} // TBD
	julianGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	julianInterest      = villagersInterest{villagerInterestMusic}
	julianName          = villagersName{villagerNameJulian}
	julianNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	julianMusic         = villagersMusic{villagerMusicSpaceKK}
	julianPersonality   = villagersPersonality{villagerPersonalitySmug}
	julianSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	julianStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCool}}
	julianWallpaper     = villagersWallpaper{villagerWallpaperStarrySkyWall}
)
