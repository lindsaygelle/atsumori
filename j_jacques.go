package atsumori

import "time"

// Jacques is an Animal Crossing villager
var Jacques = villager{
	jacquesAstrology,
	jacquesBirthDay,
	jacquesBirthMonth,
	jacquesBubbleColor,
	jacquesCategory,
	jacquesClothes,
	jacquesClothesColors,
	jacquesFlooring,
	jacquesFurniture,
	jacquesGames,
	jacquesGender,
	jacquesInterest,
	jacquesName,
	jacquesNameColor,
	jacquesMusic,
	jacquesPersonality,
	jacquesSpecies,
	jacquesStyle,
	jacquesWallpaper}

var (
	jacquesAstrology     = villagersAstrology{villagerAstrologyCancer}
	jacquesBirthDay      = villagersBirthDay{22}
	jacquesBirthMonth    = villagersBirthMonth{time.June} // 6
	jacquesBubbleColor   = villagersBubbleColor{villagerBubbleColor194C89}
	jacquesCategory      = villagersCategory{villagerCategoryA}
	jacquesClothes       = villagersClothes{villagerClothesPuffyVest} // 4277
	jacquesClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBlack}}
	jacquesFlooring      = villagersFlooring{villagerFlooringSteelFlooring}
	jacquesFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDinerCounterChair, villagerFurnitureSurveillanceCamera, villagerFurnitureDinerCounterTable, villagerFurnitureDinerMiniTable, villagerFurnitureDJsTurntable, villagerFurnitureLaptop, villagerFurnitureDinerNeonClock, villagerFurnitureStarryGarland, villagerFurnitureDinerCounterChair, villagerFurnitureDinerNeonSign, villagerFurnitureSynthesizer, villagerFurnitureDinerSofa, villagerFurnitureExitSign, villagerFurnitureStarryGarland, villagerFurnitureStarryGarland}}
	jacquesGames         = villagersGames{[]VillagerGame{}} // TBD
	jacquesGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	jacquesInterest      = villagersInterest{villagerInterestMusic}
	jacquesName          = villagersName{villagerNameJacques}
	jacquesNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	jacquesMusic         = villagersMusic{villagerMusicKKTechnopop}
	jacquesPersonality   = villagersPersonality{villagerPersonalitySmug}
	jacquesSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBird}}
	jacquesStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	jacquesWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
