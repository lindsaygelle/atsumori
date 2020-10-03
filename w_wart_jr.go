package atsumori

import "time"

// Wart Jr is an Animal Crossing villager.
var WartJr = villager{
	wartJrAstrology,
	wartJrBirthDay,
	wartJrBirthMonth,
	wartJrBubbleColor,
	wartJrCategory,
	wartJrClothes,
	wartJrClothesColors,
	wartJrFlooring,
	wartJrFurniture,
	wartJrGames,
	wartJrGender,
	wartJrInterest,
	wartJrName,
	wartJrNameColor,
	wartJrMusic,
	wartJrPersonality,
	wartJrSpecies,
	wartJrStyle,
	wartJrWallpaper}

var (
	wartJrAstrology     = villagersAstrology{villagerAstrologyLeo}
	wartJrBirthDay      = villagersBirthDay{21}
	wartJrBirthMonth    = villagersBirthMonth{time.August} // 8
	wartJrBubbleColor   = villagersBubbleColor{villagerBubbleColorD86808}
	wartJrCategory      = villagersCategory{villagerCategoryB}
	wartJrClothes       = villagersClothes{villagerClothesHantenJacket} // 9199
	wartJrClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorPurple}}
	wartJrFlooring      = villagersFlooring{villagerFlooringRushTatami}
	wartJrFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureClayFurnace, villagerFurnitureKotatsu, villagerFurnitureFuton, villagerFurniturePileOfZenCushions, villagerFurnitureFloorSeat, villagerFurnitureBambooCandleholder, villagerFurnitureBrownWoodenDeckRug, villagerFurnitureBrownWoodenDeckRug, villagerFurnitureVentilationFan, villagerFurnitureBambooBasket, villagerFurniturePendulumClock, villagerFurnitureCardboardBox, villagerFurniturePaperLantern, villagerFurnitureTapeDeck}}
	wartJrGames         = villagersGames{[]VillagerGame{}} // TBD
	wartJrGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	wartJrInterest      = villagersInterest{villagerInterestEducation}
	wartJrName          = villagersName{villagerNameWartJr}
	wartJrNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	wartJrMusic         = villagersMusic{villagerMusicSteepHill}
	wartJrPersonality   = villagersPersonality{villagerPersonalityCranky}
	wartJrSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesFrog}}
	wartJrStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	wartJrWallpaper     = villagersWallpaper{villagerWallpaperDirtClodWall}
)
