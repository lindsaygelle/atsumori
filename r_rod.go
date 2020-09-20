package atsumori

import "time"

// Rod is an Animal Crossing villager.
var Rod = villager{
	rodAstrology,
	rodBirthDay,
	rodBirthMonth,
	rodBubbleColor,
	rodCategory,
	rodClothes,
	rodClothesColors,
	rodFlooring,
	rodFurniture,
	rodGames,
	rodGender,
	rodInterest,
	rodName,
	rodNameColor,
	rodMusic,
	rodPersonality,
	rodSpecies,
	rodStyle,
	rodWallpaper}

var (
	rodAstrology     = villagersAstrology{villagerAstrologyLeo}
	rodBirthDay      = villagersBirthDay{14}
	rodBirthMonth    = villagersBirthMonth{time.August} // 8
	rodBubbleColor   = villagersBubbleColor{villagerBubbleColorEC7EFC}
	rodCategory      = villagersCategory{villagerCategoryB}
	rodClothes       = villagersClothes{villagerClothesStripedTank} // 7763
	rodClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorLightBlue}}
	rodFlooring      = villagersFlooring{villagerFlooringShipDeck}
	rodFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureJukebox, villagerFurnitureBeachTowel, villagerFurnitureBilliardTable, villagerFurniturePool, villagerFurniturePoolsideBed, villagerFurniturePalmTreeLamp}}
	rodGames         = villagersGames{[]VillagerGame{}} // TBD
	rodGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rodInterest      = villagersInterest{villagerInterestFitness}
	rodName          = villagersName{villagerNameRod}
	rodNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	rodMusic         = villagersMusic{villagerMusicKKMarch}
	rodPersonality   = villagersPersonality{villagerPersonalityJock}
	rodSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesMouse}}
	rodStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleCool}}
	rodWallpaper     = villagersWallpaper{villagerWallpaperSeaView}
)
