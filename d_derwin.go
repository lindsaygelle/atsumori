package atsumori

import "time"

// Derwin is an Animal Crossing villager
var Derwin = villager{
	derwinAstrology,
	derwinBirthDay,
	derwinBirthMonth,
	derwinBubbleColor,
	derwinCategory,
	derwinClothes,
	derwinClothesColors,
	derwinFlooring,
	derwinFurniture,
	derwinGames,
	derwinGender,
	derwinInterest,
	derwinName,
	derwinNameColor,
	derwinMusic,
	derwinPersonality,
	derwinSpecies,
	derwinStyle,
	derwinWallpaper}

var (
	derwinAstrology     = villagersAstrology{villagerAstrologyGemini}
	derwinBirthDay      = villagersBirthDay{25}
	derwinBirthMonth    = villagersBirthMonth{time.May} // 5
	derwinBubbleColor   = villagersBubbleColor{villagerBubbleColor0961F6}
	derwinCategory      = villagersCategory{villagerCategoryB}
	derwinClothes       = villagersClothes{villagerClothesStripedTank} // 7764
	derwinClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorBeige}}
	derwinFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	derwinFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureCuteMusicPlayer, villagerFurnitureSpringyRideOn, villagerFurnitureSpringyRideOn, villagerFurnitureTireToy, villagerFurnitureGardenBench, villagerFurnitureTireToy, villagerFurnitureSandbox, villagerFurnitureGardenFaucet, villagerFurnitureTricycle}}
	derwinGames         = villagersGames{[]VillagerGame{}} // TBD
	derwinGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	derwinInterest      = villagersInterest{villagerInterestPlay}
	derwinName          = villagersName{villagerNameDerwin}
	derwinNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	derwinMusic         = villagersMusic{villagerMusicKKStroll}
	derwinPersonality   = villagersPersonality{villagerPersonalityLazy}
	derwinSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	derwinStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleElegant}}
	derwinWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
