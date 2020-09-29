package atsumori

import "time"

// Stu is an Animal Crossing villager.
var Stu = villager{
	stuAstrology,
	stuBirthDay,
	stuBirthMonth,
	stuBubbleColor,
	stuCategory,
	stuClothes,
	stuClothesColors,
	stuFlooring,
	stuFurniture,
	stuGames,
	stuGender,
	stuInterest,
	stuName,
	stuNameColor,
	stuMusic,
	stuPersonality,
	stuSpecies,
	stuStyle,
	stuWallpaper}

var (
	stuAstrology     = villagersAstrology{villagerAstrologyTaurus}
	stuBirthDay      = villagersBirthDay{20}
	stuBirthMonth    = villagersBirthMonth{time.April} // 4
	stuBubbleColor   = villagersBubbleColor{villagerBubbleColor6B75CE}
	stuCategory      = villagersCategory{villagerCategoryA}
	stuClothes       = villagersClothes{villagerClothesYodelSweater} // 7721
	stuClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorBeige}}
	stuFlooring      = villagersFlooring{villagerFlooringSidewalkFlooring}
	stuFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureElectricKickScooter, villagerFurnitureElectricKickScooter, villagerFurnitureCardboardBed, villagerFurniturePhoneBox, villagerFurnitureStreetOrgan, villagerFurnitureDestinationsSignpost, villagerFurnitureCardboardBox, villagerFurnitureCassettePlayer, villagerFurnitureMountainBike}}
	stuGames         = villagersGames{[]VillagerGame{}} // TBD
	stuGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	stuInterest      = villagersInterest{villagerInterestNature}
	stuName          = villagersName{villagerNameStu}
	stuNameColor     = villagersNameColor{villagerNameColor9AE8DF}
	stuMusic         = villagersMusic{villagerMusicKKRockabilly}
	stuPersonality   = villagersPersonality{villagerPersonalityLazy}
	stuSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBull}}
	stuStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	stuWallpaper     = villagersWallpaper{villagerWallpaperChainLinkFence}
)
