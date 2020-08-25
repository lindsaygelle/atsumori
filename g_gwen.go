package atsumori

import "time"

// Gwen is an Animal Crossing villager
var Gwen = villager{
	gwenAstrology,
	gwenBirthDay,
	gwenBirthMonth,
	gwenBubbleColor,
	gwenCategory,
	gwenClothes,
	gwenClothesColors,
	gwenFlooring,
	gwenFurniture,
	gwenGames,
	gwenGender,
	gwenInterest,
	gwenName,
	gwenNameColor,
	gwenMusic,
	gwenPersonality,
	gwenSpecies,
	gwenStyle,
	gwenWallpaper}

var (
	gwenAstrology     = villagersAstrology{villagerAstrologyAquarius}
	gwenBirthDay      = villagersBirthDay{23}
	gwenBirthMonth    = villagersBirthMonth{time.January} // 1
	gwenBubbleColor   = villagersBubbleColor{villagerBubbleColor7352E8}
	gwenCategory      = villagersCategory{villagerCategoryB}
	gwenClothes       = villagersClothes{} // 2706
	gwenClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorGreen, villagerClothesColorWhite}}
	gwenFlooring      = villagersFlooring{villagerFlooringIceFlooring}
	gwenFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurnitureWhirlpoolBath, villagerFurnitureWallMountedCandle, villagerFurnitureFrozenPartition, villagerFurnitureFrozenPartition, villagerFurnitureShowerSet, villagerFurnitureFrozenCounter, villagerFurniturePortableRecordPlayer, villagerFurnitureFrozenBed, villagerFurnitureMiniFridge, villagerFurnitureFrozenTreatSet, villagerFurnitureBathroomTowelRack}}
	gwenGames         = villagersGames{[]VillagerGame{}} // TBD
	gwenGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	gwenInterest      = villagersInterest{villagerInterestFashion}
	gwenName          = villagersName{villagerNameGwen}
	gwenNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	gwenMusic         = villagersMusic{villagerMusicDrivin}
	gwenPersonality   = villagersPersonality{villagerPersonalitySnooty}
	gwenSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	gwenStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	gwenWallpaper     = villagersWallpaper{villagerWallpaperIceWall}
)
