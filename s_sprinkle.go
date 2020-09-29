package atsumori

import "time"

// Sprinkle is an Animal Crossing villager.
var Sprinkle = villager{
	sprinkleAstrology,
	sprinkleBirthDay,
	sprinkleBirthMonth,
	sprinkleBubbleColor,
	sprinkleCategory,
	sprinkleClothes,
	sprinkleClothesColors,
	sprinkleFlooring,
	sprinkleFurniture,
	sprinkleGames,
	sprinkleGender,
	sprinkleInterest,
	sprinkleName,
	sprinkleNameColor,
	sprinkleMusic,
	sprinklePersonality,
	sprinkleSpecies,
	sprinkleStyle,
	sprinkleWallpaper}

var (
	sprinkleAstrology     = villagersAstrology{villagerAstrologyPisces}
	sprinkleBirthDay      = villagersBirthDay{20}
	sprinkleBirthMonth    = villagersBirthMonth{time.February} // 2
	sprinkleBubbleColor   = villagersBubbleColor{villagerBubbleColor00D1BD}
	sprinkleCategory      = villagersCategory{villagerCategoryA}
	sprinkleClothes       = villagersClothes{villagerClothesSnowySweater} // 3631
	sprinkleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorWhite}}
	sprinkleFlooring      = villagersFlooring{villagerFlooringIceFlooring}
	sprinkleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFrozenPillar, villagerFurnitureFrozenPillar, villagerFurnitureFrozenArch, villagerFurnitureShellBed, villagerFurnitureShellWreath, villagerFurnitureShellPartition, villagerFurnitureShellPartition, villagerFurnitureShellSpeaker, villagerFurnitureShellFountain, villagerFurnitureShellLamp}}
	sprinkleGames         = villagersGames{[]VillagerGame{}} // TBD
	sprinkleGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	sprinkleInterest      = villagersInterest{villagerInterestPlay}
	sprinkleName          = villagersName{villagerNameSprinkle}
	sprinkleNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	sprinkleMusic         = villagersMusic{villagerMusicKKMarathon}
	sprinklePersonality   = villagersPersonality{villagerPersonalityPeppy}
	sprinkleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	sprinkleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCute, villagerStyleElegant}}
	sprinkleWallpaper     = villagersWallpaper{villagerWallpaperIceWall}
)
