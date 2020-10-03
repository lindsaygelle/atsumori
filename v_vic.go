package atsumori

import "time"

// Vic is an Animal Crossing villager.
var Vic = villager{
	vicAstrology,
	vicBirthDay,
	vicBirthMonth,
	vicBubbleColor,
	vicCategory,
	vicClothes,
	vicClothesColors,
	vicFlooring,
	vicFurniture,
	vicGames,
	vicGender,
	vicInterest,
	vicName,
	vicNameColor,
	vicMusic,
	vicPersonality,
	vicSpecies,
	vicStyle,
	vicWallpaper}

var (
	vicAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	vicBirthDay      = villagersBirthDay{29}
	vicBirthMonth    = villagersBirthMonth{time.December} // 12
	vicBubbleColor   = villagersBubbleColor{villagerBubbleColorC0AB72}
	vicCategory      = villagersCategory{villagerCategoryA}
	vicClothes       = villagersClothes{villagerClothesVikingTop} // 11950
	vicClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorBlue}}
	vicFlooring      = villagersFlooring{villagerFlooringShipDeck}
	vicFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJukebox, villagerFurnitureBarrel, villagerFurnitureLogRoundTable, villagerFurnitureWoodenStool, villagerFurnitureWoodenStool, villagerFurnitureOpenFrameKitchen, villagerFurnitureFreezer, villagerFurnitureBarrel, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurniturePartyGarland, villagerFurnitureProteinShakerBottle}}
	vicGames         = villagersGames{[]VillagerGame{}} // TBD
	vicGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	vicInterest      = villagersInterest{villagerInterestFitness}
	vicName          = villagersName{villagerNameVic}
	vicNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	vicMusic         = villagersMusic{villagerMusicKKMarch}
	vicPersonality   = villagersPersonality{villagerPersonalityCranky}
	vicSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBull}}
	vicStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleActive}}
	vicWallpaper     = villagersWallpaper{villagerWallpaperSeaView}
)
