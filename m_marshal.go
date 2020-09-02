package atsumori

import "time"

// Marshal is an Animal Crossing villager.
var Marshal = villager{
	marshalAstrology,
	marshalBirthDay,
	marshalBirthMonth,
	marshalBubbleColor,
	marshalCategory,
	marshalClothes,
	marshalClothesColors,
	marshalFlooring,
	marshalFurniture,
	marshalGames,
	marshalGender,
	marshalInterest,
	marshalName,
	marshalNameColor,
	marshalMusic,
	marshalPersonality,
	marshalSpecies,
	marshalStyle,
	marshalWallpaper}

var (
	marshalAstrology     = villagersAstrology{villagerAstrologyLibra}
	marshalBirthDay      = villagersBirthDay{29}
	marshalBirthMonth    = villagersBirthMonth{time.September} // 9
	marshalBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	marshalCategory      = villagersCategory{villagerCategoryA}
	marshalClothes       = villagersClothes{villagerClothesPuffyVest} // 4277
	marshalClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorBlue}}
	marshalFlooring      = villagersFlooring{villagerFlooringModernWoodFlooring}
	marshalFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureDoubleSofa, villagerFurnitureUprightPiano, villagerFurniturePianoBench, villagerFurnitureMiniFridge, villagerFurnitureCreamAndSugar, villagerFurnitureOpenFrameKitchen, villagerFurnitureSimpleKettle, villagerFurnitureBlueKitchenMat, villagerFurnitureIronwoodCupboard, villagerFurnitureIronWorktable, villagerFurnitureCuckooClock, villagerFurnitureIronEntranceMat, villagerFurniturePortableRecordPlayer, villagerFurnitureCoffeeGrinder, villagerFurnitureEspressoMaker, villagerFurnitureStovetopEspressoMaker, villagerFurnitureMenuChalkboard, villagerFurnitureIronwoodLowTable, villagerFurnitureCoffeeCup}}
	marshalGames         = villagersGames{[]VillagerGame{}} // TBD
	marshalGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	marshalInterest      = villagersInterest{villagerInterestMusic}
	marshalName          = villagersName{villagerNameMarshal}
	marshalNameColor     = villagersNameColor{villagerNameColor848484}
	marshalMusic         = villagersMusic{villagerMusicKKBossa}
	marshalPersonality   = villagersPersonality{villagerPersonalitySmug}
	marshalSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	marshalStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	marshalWallpaper     = villagersWallpaper{villagerWallpaperBlueDelicateBloomsWall}
)
