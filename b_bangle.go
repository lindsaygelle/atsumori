package atsumori

import "time"

// Bangle is an Animal Crossing villager
var Bangle = villager{
	bangleAstrology,
	bangleBirthDay,
	bangleBirthMonth,
	bangleBubbleColor,
	bangleCategory,
	bangleClothes,
	bangleClothesColors,
	bangleFlooring,
	bangleFurniture,
	bangleGames,
	bangleGender,
	bangleInterest,
	bangleName,
	bangleNameColor,
	bangleMusic,
	banglePersonality,
	bangleSpecies,
	bangleStyle,
	bangleWallpaper}

var (
	bangleAstrology     = villagersAstrology{villagerAstrologyVirgo}
	bangleBirthDay      = villagersBirthDay{27}
	bangleBirthMonth    = villagersBirthMonth{time.August} // 8
	bangleBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	bangleCategory      = villagersCategory{villagerCategoryB}
	bangleClothes       = villagersClothes{} // 2687
	bangleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorGreen}}
	bangleFlooring      = villagersFlooring{villagerFlooringBlueFloralFlooring}
	bangleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenFull - LengthMirror, villagerFurnitureWoodenLowTable, villagerFurnitureCuteMusicPlayer, villagerFurnitureDouble - DoorRefrigerator, villagerFurnitureAirConditioner, villagerFurnitureDoubleSofa, villagerFurnitureWoodenDoubleBed, villagerFurnitureWoodenEndTable, villagerFurnitureMagneticKnifeRack, villagerFurnitureWoodenChest, villagerFurnitureAromaPot, villagerFurnitureIronwoodKitchenette, villagerFurnitureDish - DryingRack, villagerFurnitureBabyBear}}
	bangleGames         = villagersGames{[]VillagerGame{}} // TBD
	bangleGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	bangleInterest      = villagersInterest{villagerInterestFashion}
	bangleName          = villagersName{villagerNameBangle}
	bangleNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	bangleMusic         = villagersMusic{} // K.K. Soul
	banglePersonality   = villagersPersonality{villagerPersonalityPeppy}
	bangleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesTiger}}
	bangleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCute}}
	bangleWallpaper     = villagersWallpaper{villagerWallpaperYellowIntricateWall}
)
