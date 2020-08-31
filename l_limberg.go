package atsumori

import "time"

// Limberg is an Animal Crossing villager.
var Limberg = villager{
	limbergAstrology,
	limbergBirthDay,
	limbergBirthMonth,
	limbergBubbleColor,
	limbergCategory,
	limbergClothes,
	limbergClothesColors,
	limbergFlooring,
	limbergFurniture,
	limbergGames,
	limbergGender,
	limbergInterest,
	limbergName,
	limbergNameColor,
	limbergMusic,
	limbergPersonality,
	limbergSpecies,
	limbergStyle,
	limbergWallpaper}

var (
	limbergAstrology     = villagersAstrology{villagerAstrologyLibra}
	limbergBirthDay      = villagersBirthDay{17}
	limbergBirthMonth    = villagersBirthMonth{time.October} // 10
	limbergBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	limbergCategory      = villagersCategory{villagerCategoryB}
	limbergClothes       = villagersClothes{villagerClothesHantenJacket} // 3164
	limbergClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorBlue}}
	limbergFlooring      = villagersFlooring{villagerFlooringMossyGardenFlooring}
	limbergFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBambooSpeaker, villagerFurnitureMapleLeafPondStone, villagerFurnitureOutdoorBath, villagerFurnitureRedLeafPile, villagerFurnitureRedLeafPile, villagerFurnitureLeafCampfire, villagerFurnitureLeafStool}}
	limbergGames         = villagersGames{[]VillagerGame{}} // TBD
	limbergGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	limbergInterest      = villagersInterest{villagerInterestEducation}
	limbergName          = villagersName{villagerNameLimberg}
	limbergNameColor     = villagersNameColor{villagerNameColor874C25}
	limbergMusic         = villagersMusic{villagerMusicKKFolk}
	limbergPersonality   = villagersPersonality{villagerPersonalityCranky}
	limbergSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	limbergStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	limbergWallpaper     = villagersWallpaper{villagerWallpaperAutumnWall}
)
