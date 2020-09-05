package atsumori

import "time"

// Mint is an Animal Crossing villager.
var Mint = villager{
	mintAstrology,
	mintBirthDay,
	mintBirthMonth,
	mintBubbleColor,
	mintCategory,
	mintClothes,
	mintClothesColors,
	mintFlooring,
	mintFurniture,
	mintGames,
	mintGender,
	mintInterest,
	mintName,
	mintNameColor,
	mintMusic,
	mintPersonality,
	mintSpecies,
	mintStyle,
	mintWallpaper}

var (
	mintAstrology     = villagersAstrology{villagerAstrologyTaurus}
	mintBirthDay      = villagersBirthDay{2}
	mintBirthMonth    = villagersBirthMonth{time.May} // 5
	mintBubbleColor   = villagersBubbleColor{villagerBubbleColor87E0A9}
	mintCategory      = villagersCategory{villagerCategoryB}
	mintClothes       = villagersClothes{} // 2686
	mintClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorPurple}}
	mintFlooring      = villagersFlooring{villagerFlooringMintDotFlooring}
	mintFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanArmchair, villagerFurnitureLilyRecordPlayer, villagerFurnitureRattanWardrobe, villagerFurnitureRattanEndTable, villagerFurnitureShowerBooth, villagerFurnitureRattanVanity, villagerFurniturePlainSink, villagerFurnitureFloralSwag, villagerFurnitureBathroomTowelRack, villagerFurnitureRattanLowTable, villagerFurnitureTeaSet}}
	mintGames         = villagersGames{[]VillagerGame{}} // TBD
	mintGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	mintInterest      = villagersInterest{villagerInterestFashion}
	mintName          = villagersName{villagerNameMint}
	mintNameColor     = villagersNameColor{villagerNameColor219479}
	mintMusic         = villagersMusic{villagerMusicKKSoul}
	mintPersonality   = villagersPersonality{villagerPersonalitySnooty}
	mintSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesSquirrel}}
	mintStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCute}}
	mintWallpaper     = villagersWallpaper{villagerWallpaperBlueSubwayTileWall}
)
