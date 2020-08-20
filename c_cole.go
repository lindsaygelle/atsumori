package atsumori

import "time"

// Cole is an Animal Crossing villager
var Cole = villager{
	coleAstrology,
	coleBirthDay,
	coleBirthMonth,
	coleBubbleColor,
	coleCategory,
	coleClothes,
	coleClothesColors,
	coleFlooring,
	coleFurniture,
	coleGames,
	coleGender,
	coleInterest,
	coleName,
	coleNameColor,
	coleMusic,
	colePersonality,
	coleSpecies,
	coleStyle,
	coleWallpaper}

var (
	coleAstrology     = villagersAstrology{villagerAstrologyLeo}
	coleBirthDay      = villagersBirthDay{10}
	coleBirthMonth    = villagersBirthMonth{time.August} // 8
	coleBubbleColor   = villagersBubbleColor{villagerBubbleColor515151}
	coleCategory      = villagersCategory{villagerCategoryA}
	coleClothes       = villagersClothes{} // 3288
	coleClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorOrange, villagerClothesColorColorful}}
	coleFlooring      = villagersFlooring{villagerFlooringSandyBeachFlooring}
	coleFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBeachTowel, villagerFurniturePalmTreeLamp, villagerFurnitureWaveBreaker, villagerFurnitureBeachChair, villagerFurnitureUkulele, villagerFurnitureSandCastle, villagerFurnitureLifeRing, villagerFurnitureSurfboard, villagerFurnitureBeachBall, villagerFurniturePortableRadio, villagerFurnitureLifeguardChair, villagerFurnitureColorfulVinylSheet}}
	coleGames         = villagersGames{[]VillagerGame{}} // TBD
	coleGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	coleInterest      = villagersInterest{villagerInterestNature}
	coleName          = villagersName{villagerNameCole}
	coleNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	coleMusic         = villagersMusic{} // K.K. Faire
	colePersonality   = villagersPersonality{villagerPersonalityLazy}
	coleSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesRabbit}}
	coleStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleCute}}
	coleWallpaper     = villagersWallpaper{villagerWallpaperTropicalVista}
)
