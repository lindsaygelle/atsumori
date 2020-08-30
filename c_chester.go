package atsumori

import "time"

// Chester is an Animal Crossing villager.
var Chester = villager{
	chesterAstrology,
	chesterBirthDay,
	chesterBirthMonth,
	chesterBubbleColor,
	chesterCategory,
	chesterClothes,
	chesterClothesColors,
	chesterFlooring,
	chesterFurniture,
	chesterGames,
	chesterGender,
	chesterInterest,
	chesterName,
	chesterNameColor,
	chesterMusic,
	chesterPersonality,
	chesterSpecies,
	chesterStyle,
	chesterWallpaper}

var (
	chesterAstrology     = villagersAstrology{villagerAstrologyLeo}
	chesterBirthDay      = villagersBirthDay{6}
	chesterBirthMonth    = villagersBirthMonth{time.August} // 8
	chesterBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	chesterCategory      = villagersCategory{villagerCategoryB}
	chesterClothes       = villagersClothes{} // 3462
	chesterClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorGreen}}
	chesterFlooring      = villagersFlooring{villagerFlooringBambooFlooring}
	chesterFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBambooNoodleSlide, villagerFurnitureBambooDoll, villagerFurnitureBambooSpeaker, villagerFurnitureBambooBench, villagerFurnitureBambooLunchBox, villagerFurnitureBambooShootLamp, villagerFurnitureBambooShootLamp, villagerFurnitureBambooBench, villagerFurnitureRedCarpet, villagerFurnitureBabyPanda}}
	chesterGames         = villagersGames{[]VillagerGame{}} // TBD
	chesterGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	chesterInterest      = villagersInterest{villagerInterestPlay}
	chesterName          = villagersName{villagerNameChester}
	chesterNameColor     = villagersNameColor{villagerNameColor848484}
	chesterMusic         = villagersMusic{villagerMusicImperialKK}
	chesterPersonality   = villagersPersonality{villagerPersonalityLazy}
	chesterSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	chesterStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	chesterWallpaper     = villagersWallpaper{villagerWallpaperBambooGroveWall}
)
