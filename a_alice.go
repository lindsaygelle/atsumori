package atsumori

import "time"

// Alice is an Animal Crossing villager.
var Alice = villager{
	aliceAstrology,
	aliceBirthDay,
	aliceBirthMonth,
	aliceBubbleColor,
	aliceCategory,
	aliceClothes,
	aliceClothesColors,
	aliceFlooring,
	aliceFurniture,
	aliceGames,
	aliceGender,
	aliceInterest,
	aliceName,
	aliceNameColor,
	aliceMusic,
	alicePersonality,
	aliceSpecies,
	aliceStyle,
	aliceWallpaper}

var (
	aliceAstrology     = villagersAstrology{villagerAstrologyLeo}
	aliceBirthDay      = villagersBirthDay{19}
	aliceBirthMonth    = villagersBirthMonth{time.August}
	aliceBubbleColor   = villagersBubbleColor{villagerBubbleColorE8B010}
	aliceCategory      = villagersCategory{villagerCategoryB}
	aliceClothes       = villagersClothes{} // 4729
	aliceClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorRed, villagerClothesColorPink}}
	aliceFlooring      = villagersFlooring{villagerFlooringLightParquetFlooring}
	aliceFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureRattanBed, villagerFurnitureRattanLowTable, villagerFurnitureMonstera, villagerFurnitureRattanArmchair, villagerFurnitureRattanWardrobe, villagerFurnitureOpenFrameKitchen, villagerFurnitureRattanVanity, villagerFurnitureBotanicalRug, villagerFurnitureCoconutWallPlanter, villagerFurnitureWallMountedTV50In, villagerFurnitureRattanEndTable, villagerFurniturePortableRecordPlayer}}
	aliceGames         = villagersGames{[]VillagerGame{}} // TBD
	aliceGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	aliceInterest      = villagersInterest{villagerInterestEducation}
	aliceName          = villagersName{villagerNameAlice}
	aliceNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	aliceMusic         = villagersMusic{villagerMusicAlohaKK}
	alicePersonality   = villagersPersonality{villagerPersonalityNormal}
	aliceSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesKoala}}
	aliceStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleCute}}
	aliceWallpaper     = villagersWallpaper{villagerWallpaperWhiteBotanicalTileWall}
)
