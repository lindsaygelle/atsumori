package atsumori

import "time"

// Rodeo is an Animal Crossing villager.
var Rodeo = villager{
	rodeoAstrology,
	rodeoBirthDay,
	rodeoBirthMonth,
	rodeoBubbleColor,
	rodeoCategory,
	rodeoClothes,
	rodeoClothesColors,
	rodeoFlooring,
	rodeoFurniture,
	rodeoGames,
	rodeoGender,
	rodeoInterest,
	rodeoName,
	rodeoNameColor,
	rodeoMusic,
	rodeoPersonality,
	rodeoSpecies,
	rodeoStyle,
	rodeoWallpaper}

var (
	rodeoAstrology     = villagersAstrology{villagerAstrologyScorpio}
	rodeoBirthDay      = villagersBirthDay{29}
	rodeoBirthMonth    = villagersBirthMonth{time.October} // 10
	rodeoBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	rodeoCategory      = villagersCategory{villagerCategoryB}
	rodeoClothes       = villagersClothes{} // 3140
	rodeoClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorRed}}
	rodeoFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	rodeoFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHighEndStereo, villagerFurnitureAntiqueBed, villagerFurnitureDenChair, villagerFurnitureTelescope, villagerFurnitureDenDesk, villagerFurnitureDesktopComputer, villagerFurnitureWoodenWasteBin, villagerFurnitureWoodenChest, villagerFurnitureWoodenBookshelf, villagerFurniturePlasmaBall}}
	rodeoGames         = villagersGames{[]VillagerGame{}} // TBD
	rodeoGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rodeoInterest      = villagersInterest{villagerInterestFitness}
	rodeoName          = villagersName{villagerNameRodeo}
	rodeoNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	rodeoMusic         = villagersMusic{villagerMusicKKTechnopop}
	rodeoPersonality   = villagersPersonality{villagerPersonalityLazy}
	rodeoSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBull}}
	rodeoStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	rodeoWallpaper     = villagersWallpaper{villagerWallpaperConcreteWall}
)
