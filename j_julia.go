package atsumori

import "time"

// Julia is an Animal Crossing villager
var Julia = villager{
	juliaAstrology,
	juliaBirthDay,
	juliaBirthMonth,
	juliaBubbleColor,
	juliaCategory,
	juliaClothes,
	juliaClothesColors,
	juliaFlooring,
	juliaFurniture,
	juliaGames,
	juliaGender,
	juliaInterest,
	juliaName,
	juliaNameColor,
	juliaMusic,
	juliaPersonality,
	juliaSpecies,
	juliaStyle,
	juliaWallpaper}

var (
	juliaAstrology     = villagersAstrology{villagerAstrologyLeo}
	juliaBirthDay      = villagersBirthDay{31}
	juliaBirthMonth    = villagersBirthMonth{time.July} // 7
	juliaBubbleColor   = villagersBubbleColor{villagerBubbleColor0CA54A}
	juliaCategory      = villagersCategory{villagerCategoryA}
	juliaClothes       = villagersClothes{} // 8817
	juliaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorRed}}
	juliaFlooring      = villagersFlooring{villagerFlooringSimpleBlueFlooring}
	juliaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanTowelBasket, villagerFurnitureShowerBooth, villagerFurnitureRattanLowTable, villagerFurnitureBeachTowel, villagerFurnitureRattanVanity, villagerFurnitureRattanStool, villagerFurnitureRattanEndTable, villagerFurniturePortableRecordPlayer, villagerFurnitureBlackWoodenDeckRug, villagerFurnitureLongBathtub, villagerFurnitureBathroomTowelRack, villagerFurnitureBlackWoodenDeckRug}}
	juliaGames         = villagersGames{[]VillagerGame{}} // TBD
	juliaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	juliaInterest      = villagersInterest{villagerInterestEducation}
	juliaName          = villagersName{villagerNameJulia}
	juliaNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	juliaMusic         = villagersMusic{villagerMusicKKSonata}
	juliaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	juliaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesOstrich}}
	juliaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	juliaWallpaper     = villagersWallpaper{villagerWallpaperGreenMoldedPanelWall}
)
