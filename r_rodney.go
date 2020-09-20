package atsumori

import "time"

// Rodney is an Animal Crossing villager.
var Rodney = villager{
	rodneyAstrology,
	rodneyBirthDay,
	rodneyBirthMonth,
	rodneyBubbleColor,
	rodneyCategory,
	rodneyClothes,
	rodneyClothesColors,
	rodneyFlooring,
	rodneyFurniture,
	rodneyGames,
	rodneyGender,
	rodneyInterest,
	rodneyName,
	rodneyNameColor,
	rodneyMusic,
	rodneyPersonality,
	rodneySpecies,
	rodneyStyle,
	rodneyWallpaper}

var (
	rodneyAstrology     = villagersAstrology{villagerAstrologyScorpio}
	rodneyBirthDay      = villagersBirthDay{10}
	rodneyBirthMonth    = villagersBirthMonth{time.November} // 11
	rodneyBubbleColor   = villagersBubbleColor{villagerBubbleColor459ABA}
	rodneyCategory      = villagersCategory{villagerCategoryA}
	rodneyClothes       = villagersClothes{villagerClothesStripedTank} // 7768
	rodneyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPink, villagerClothesColorColorful}}
	rodneyFlooring      = villagersFlooring{villagerFlooringGreenVinylFlooring}
	rodneyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBilliardTable, villagerFurnitureDartboard, villagerFurnitureJukebox, villagerFurnitureBoxSofa, villagerFurnitureBoxCornerSofa, villagerFurnitureBoxSofa, villagerFurnitureDinerMiniTable, villagerFurnitureCacaoTree, villagerFurnitureDinerNeonClock, villagerFurnitureCandyMachine, villagerFurniturePinballMachine}}
	rodneyGames         = villagersGames{[]VillagerGame{}} // TBD
	rodneyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rodneyInterest      = villagersInterest{villagerInterestMusic}
	rodneyName          = villagersName{villagerNameRodney}
	rodneyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	rodneyMusic         = villagersMusic{villagerMusicKKRockabilly}
	rodneyPersonality   = villagersPersonality{villagerPersonalitySmug}
	rodneySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHamster}}
	rodneyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleCool}}
	rodneyWallpaper     = villagersWallpaper{villagerWallpaperBlueDinerWall}
)
