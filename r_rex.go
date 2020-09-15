package atsumori

import "time"

// Rex is an Animal Crossing villager.
var Rex = villager{
	rexAstrology,
	rexBirthDay,
	rexBirthMonth,
	rexBubbleColor,
	rexCategory,
	rexClothes,
	rexClothesColors,
	rexFlooring,
	rexFurniture,
	rexGames,
	rexGender,
	rexInterest,
	rexName,
	rexNameColor,
	rexMusic,
	rexPersonality,
	rexSpecies,
	rexStyle,
	rexWallpaper}

var (
	rexAstrology     = villagersAstrology{villagerAstrologyLeo}
	rexBirthDay      = villagersBirthDay{24}
	rexBirthMonth    = villagersBirthMonth{time.July} // 7
	rexBubbleColor   = villagersBubbleColor{villagerBubbleColorFFAA3B}
	rexCategory      = villagersCategory{villagerCategoryA}
	rexClothes       = villagersClothes{villagerClothesStripedShirt} // 2655
	rexClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	rexFlooring      = villagersFlooring{villagerFlooringDaisyMeadow}
	rexFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureBonfire, villagerFurnitureLantern, villagerFurnitureSmoker, villagerFurnitureLogBench, villagerFurniturePicnicBasket, villagerFurnitureSleepingBag, villagerFurnitureCampfireCookware, villagerFurnitureLogStool, villagerFurniturePortableRadio}}
	rexGames         = villagersGames{[]VillagerGame{}} // TBD
	rexGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rexInterest      = villagersInterest{villagerInterestNature}
	rexName          = villagersName{villagerNameRex}
	rexNameColor     = villagersNameColor{villagerNameColor874C25}
	rexMusic         = villagersMusic{villagerMusicWandering}
	rexPersonality   = villagersPersonality{villagerPersonalityLazy}
	rexSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesLion}}
	rexStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	rexWallpaper     = villagersWallpaper{villagerWallpaperSummitWall}
)
