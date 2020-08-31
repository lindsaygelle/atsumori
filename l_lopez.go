package atsumori

import "time"

// Lopez is an Animal Crossing villager.
var Lopez = villager{
	lopezAstrology,
	lopezBirthDay,
	lopezBirthMonth,
	lopezBubbleColor,
	lopezCategory,
	lopezClothes,
	lopezClothesColors,
	lopezFlooring,
	lopezFurniture,
	lopezGames,
	lopezGender,
	lopezInterest,
	lopezName,
	lopezNameColor,
	lopezMusic,
	lopezPersonality,
	lopezSpecies,
	lopezStyle,
	lopezWallpaper}

var (
	lopezAstrology     = villagersAstrology{villagerAstrologyLeo}
	lopezBirthDay      = villagersBirthDay{20}
	lopezBirthMonth    = villagersBirthMonth{time.August} // 8
	lopezBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	lopezCategory      = villagersCategory{villagerCategoryA}
	lopezClothes       = villagersClothes{villagerClothesChimayoVest} // 4507
	lopezClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGray, villagerClothesColorGray}}
	lopezFlooring      = villagersFlooring{villagerFlooringModernWoodFlooring}
	lopezFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodBurningStove, villagerFurnitureRefrigerator, villagerFurnitureIronwoodBed, villagerFurnitureRockingChair, villagerFurnitureIronwoodKitchenette, villagerFurnitureAntiqueConsoleTable, villagerFurniturePortableRecordPlayer, villagerFurnitureIronwoodClock, villagerFurnitureDeerDecoration, villagerFurnitureIronwoodCart, villagerFurnitureEspressoMaker, villagerFurnitureIronwoodCupboard, villagerFurnitureMicrowave, villagerFurnitureAnalogKitchenScale, villagerFurnitureIvorySmallRoundMat, villagerFurnitureTapestry}}
	lopezGames         = villagersGames{[]VillagerGame{}} // TBD
	lopezGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	lopezInterest      = villagersInterest{villagerInterestEducation}
	lopezName          = villagersName{villagerNameLopez}
	lopezNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	lopezMusic         = villagersMusic{villagerMusicKKJazz}
	lopezPersonality   = villagersPersonality{villagerPersonalitySmug}
	lopezSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDeer}}
	lopezStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleGorgeous}}
	lopezWallpaper     = villagersWallpaper{villagerWallpaperModernWoodWall}
)
