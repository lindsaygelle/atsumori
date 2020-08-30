package atsumori

import "time"

// Felicity is an Animal Crossing villager.
var Felicity = villager{
	felicityAstrology,
	felicityBirthDay,
	felicityBirthMonth,
	felicityBubbleColor,
	felicityCategory,
	felicityClothes,
	felicityClothesColors,
	felicityFlooring,
	felicityFurniture,
	felicityGames,
	felicityGender,
	felicityInterest,
	felicityName,
	felicityNameColor,
	felicityMusic,
	felicityPersonality,
	felicitySpecies,
	felicityStyle,
	felicityWallpaper}

var (
	felicityAstrology     = villagersAstrology{villagerAstrologyAries}
	felicityBirthDay      = villagersBirthDay{30}
	felicityBirthMonth    = villagersBirthMonth{time.March} // 3
	felicityBubbleColor   = villagersBubbleColor{villagerBubbleColorFFF98F}
	felicityCategory      = villagersCategory{villagerCategoryB}
	felicityClothes       = villagersClothes{} // 2677
	felicityClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorOrange}}
	felicityFlooring      = villagersFlooring{villagerFlooringDaisyMeadow}
	felicityFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCypressPlant, villagerFurnitureCypressPlant, villagerFurnitureGardenFaucet, villagerFurnitureLogStool, villagerFurnitureCuteMusicPlayer, villagerFurnitureIronGardenTable, villagerFurnitureIronGardenChair, villagerFurnitureBirdbath, villagerFurnitureIronGardenBench, villagerFurnitureBirdhouse, villagerFurnitureHammock}}
	felicityGames         = villagersGames{[]VillagerGame{}} // TBD
	felicityGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	felicityInterest      = villagersInterest{villagerInterestFashion}
	felicityName          = villagersName{villagerNameFelicity}
	felicityNameColor     = villagersNameColor{villagerNameColor879B96}
	felicityMusic         = villagersMusic{villagerMusicKKStroll}
	felicityPersonality   = villagersPersonality{villagerPersonalityPeppy}
	felicitySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	felicityStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	felicityWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
