package atsumori

import "time"

// Barold is an Animal Crossing villager.
var Barold = villager{
	baroldAstrology,
	baroldBirthDay,
	baroldBirthMonth,
	baroldBubbleColor,
	baroldCategory,
	baroldClothes,
	baroldClothesColors,
	baroldFlooring,
	baroldFurniture,
	baroldGames,
	baroldGender,
	baroldInterest,
	baroldName,
	baroldNameColor,
	baroldMusic,
	baroldPersonality,
	baroldSpecies,
	baroldStyle,
	baroldWallpaper}

var (
	baroldAstrology     = villagersAstrology{villagerAstrologyPisces}
	baroldBirthDay      = villagersBirthDay{2}
	baroldBirthMonth    = villagersBirthMonth{time.March} // 3
	baroldBubbleColor   = villagersBubbleColor{villagerBubbleColorBFBFBF}
	baroldCategory      = villagersCategory{villagerCategoryA}
	baroldClothes       = villagersClothes{villagerClothesAnimalStripesTee} // 8198
	baroldClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorBlack}}
	baroldFlooring      = villagersFlooring{villagerFlooringMonochromaticTileFlooring}
	baroldFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureShowerBooth, villagerFurnitureBunkBed, villagerFurnitureUprightLocker, villagerFurnitureSleepingBag, villagerFurnitureBreaker, villagerFurnitureOfficeDesk, villagerFurnitureLaptop, villagerFurniturePortableRadio, villagerFurnitureModernOfficeChair, villagerFurnitureServer, villagerFurnitureWaterCooler, villagerFurnitureWallMountedPhone, villagerFurnitureWallMountedPhone, villagerFurnitureCardboardBox, villagerFurnitureSurveillanceCamera, villagerFurnitureStackOfBooks, villagerFurnitureAluminumRug}}
	baroldGames         = villagersGames{[]VillagerGame{}} // TBD
	baroldGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	baroldInterest      = villagersInterest{villagerInterestPlay}
	baroldName          = villagersName{villagerNameBarold}
	baroldNameColor     = villagersNameColor{villagerNameColor5E5E5E}
	baroldMusic         = villagersMusic{villagerMusicKKSong}
	baroldPersonality   = villagersPersonality{villagerPersonalityLazy}
	baroldSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCub}}
	baroldStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCool}}
	baroldWallpaper     = villagersWallpaper{villagerWallpaperSecurityMonitorsWall}
)
