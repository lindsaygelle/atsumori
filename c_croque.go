package atsumori

import "time"

// Croque is an Animal Crossing villager
var Croque = villager{
	croqueAstrology,
	croqueBirthDay,
	croqueBirthMonth,
	croqueBubbleColor,
	croqueCategory,
	croqueClothes,
	croqueClothesColors,
	croqueFlooring,
	croqueFurniture,
	croqueGames,
	croqueGender,
	croqueInterest,
	croqueName,
	croqueNameColor,
	croqueMusic,
	croquePersonality,
	croqueSpecies,
	croqueStyle,
	croqueWallpaper}

var (
	croqueAstrology     = villagersAstrology{villagerAstrologyCancer}
	croqueBirthDay      = villagersBirthDay{18}
	croqueBirthMonth    = villagersBirthMonth{time.July} // 7
	croqueBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	croqueCategory      = villagersCategory{villagerCategoryA}
	croqueClothes       = villagersClothes{villagerClothesSilkShirt} // 9127
	croqueClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorOrange}}
	croqueFlooring      = villagersFlooring{villagerFlooringColoredLeavesFlooring}
	croqueFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureSimpleDIYWorkbench, villagerFurnitureBambooSpeaker, villagerFurnitureTreeStandee, villagerFurnitureTreeStandee, villagerFurnitureTreeStandee, villagerFurnitureHayBed, villagerFurnitureShantyMat, villagerFurnitureLeafCampfire, villagerFurnitureLeafStool, villagerFurnitureMapleLeafPondStone}}
	croqueGames         = villagersGames{[]VillagerGame{}} // TBD
	croqueGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	croqueInterest      = villagersInterest{villagerInterestNature}
	croqueName          = villagersName{villagerNameCroque}
	croqueNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	croqueMusic         = villagersMusic{} // K.K. Folk
	croquePersonality   = villagersPersonality{villagerPersonalityCranky}
	croqueSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	croqueStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	croqueWallpaper     = villagersWallpaper{villagerWallpaperAutumnWall}
)
