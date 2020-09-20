package atsumori

import "time"

// Roscoe is an Animal Crossing villager.
var Roscoe = villager{
	roscoeAstrology,
	roscoeBirthDay,
	roscoeBirthMonth,
	roscoeBubbleColor,
	roscoeCategory,
	roscoeClothes,
	roscoeClothesColors,
	roscoeFlooring,
	roscoeFurniture,
	roscoeGames,
	roscoeGender,
	roscoeInterest,
	roscoeName,
	roscoeNameColor,
	roscoeMusic,
	roscoePersonality,
	roscoeSpecies,
	roscoeStyle,
	roscoeWallpaper}

var (
	roscoeAstrology     = villagersAstrology{villagerAstrologyGemini}
	roscoeBirthDay      = villagersBirthDay{16}
	roscoeBirthMonth    = villagersBirthMonth{time.June} // 6
	roscoeBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	roscoeCategory      = villagersCategory{villagerCategoryB}
	roscoeClothes       = villagersClothes{villagerClothesBikerJacket} // 3198
	roscoeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorGray}}
	roscoeFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	roscoeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureDoubleSofa, villagerFurnitureDIYWorkbench, villagerFurnitureIronWallLamp, villagerFurnitureElectricGuitar, villagerFurnitureRattanEndTable, villagerFurnitureThrowbackSkullRadio, villagerFurnitureRattanLowTable, villagerFurnitureRattanTableLamp, villagerFurniturePortableRecordPlayer, villagerFurnitureIronHangerStand, villagerFurnitureWallClock, villagerFurnitureSkullDoorplate, villagerFurnitureAmp, villagerFurnitureAirConditioner}}
	roscoeGames         = villagersGames{[]VillagerGame{}} // TBD
	roscoeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	roscoeInterest      = villagersInterest{villagerInterestMusic}
	roscoeName          = villagersName{villagerNameRoscoe}
	roscoeNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	roscoeMusic         = villagersMusic{villagerMusicKKDB}
	roscoePersonality   = villagersPersonality{villagerPersonalityCranky}
	roscoeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHorse}}
	roscoeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleGorgeous}}
	roscoeWallpaper     = villagersWallpaper{villagerWallpaperBlackCrownWall}
)
