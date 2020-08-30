package atsumori

import "time"

// Hopper is an Animal Crossing villager.
var Hopper = villager{
	hopperAstrology,
	hopperBirthDay,
	hopperBirthMonth,
	hopperBubbleColor,
	hopperCategory,
	hopperClothes,
	hopperClothesColors,
	hopperFlooring,
	hopperFurniture,
	hopperGames,
	hopperGender,
	hopperInterest,
	hopperName,
	hopperNameColor,
	hopperMusic,
	hopperPersonality,
	hopperSpecies,
	hopperStyle,
	hopperWallpaper}

var (
	hopperAstrology     = villagersAstrology{villagerAstrologyAries}
	hopperBirthDay      = villagersBirthDay{6}
	hopperBirthMonth    = villagersBirthMonth{time.April} // 4
	hopperBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	hopperCategory      = villagersCategory{villagerCategoryB}
	hopperClothes       = villagersClothes{} // 3462
	hopperClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorRed}}
	hopperFlooring      = villagersFlooring{villagerFlooringIcebergFlooring}
	hopperFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureTreeStandee, villagerFurnitureTreeStandee, villagerFurnitureFrozenPillar, villagerFurnitureStoneStool, villagerFurnitureOutdoorBath, villagerFurnitureThreeTieredSnowperson, villagerFurniturePortableRadio, villagerFurnitureFrozenPillar}}
	hopperGames         = villagersGames{[]VillagerGame{}} // TBD
	hopperGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	hopperInterest      = villagersInterest{villagerInterestMusic}
	hopperName          = villagersName{villagerNameHopper}
	hopperNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	hopperMusic         = villagersMusic{villagerMusicComradeKK}
	hopperPersonality   = villagersPersonality{villagerPersonalityCranky}
	hopperSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	hopperStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	hopperWallpaper     = villagersWallpaper{villagerWallpaperIcebergWall}
)
