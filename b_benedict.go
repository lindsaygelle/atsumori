package atsumori

import "time"

// Benedict is an Animal Crossing villager.
var Benedict = villager{
	benedictAstrology,
	benedictBirthDay,
	benedictBirthMonth,
	benedictBubbleColor,
	benedictCategory,
	benedictClothes,
	benedictClothesColors,
	benedictFlooring,
	benedictFurniture,
	benedictGames,
	benedictGender,
	benedictInterest,
	benedictName,
	benedictNameColor,
	benedictMusic,
	benedictPersonality,
	benedictSpecies,
	benedictStyle,
	benedictWallpaper}

var (
	benedictAstrology     = villagersAstrology{villagerAstrologyLibra}
	benedictBirthDay      = villagersBirthDay{10}
	benedictBirthMonth    = villagersBirthMonth{time.October} // 10
	benedictBubbleColor   = villagersBubbleColor{villagerBubbleColorFF4040}
	benedictCategory      = villagersCategory{villagerCategoryB}
	benedictClothes       = villagersClothes{villagerClothesTwoBallTee} // 2541
	benedictClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorPurple}}
	benedictFlooring      = villagersFlooring{villagerFlooringBackyardLawn}
	benedictFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureInflatableSofa, villagerFurnitureTricycle, villagerFurnitureCuteMusicPlayer, villagerFurniturePlasticPool, villagerFurnitureHoseReel, villagerFurnitureGardenFaucet, villagerFurnitureBeachTowel, villagerFurnitureBarbecue, villagerFurnitureLifeRing, villagerFurnitureDoghouse}}
	benedictGames         = villagersGames{[]VillagerGame{}} // TBD
	benedictGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	benedictInterest      = villagersInterest{villagerInterestPlay}
	benedictName          = villagersName{villagerNameBenedict}
	benedictNameColor     = villagersNameColor{villagerNameColorFFFAD4}
	benedictMusic         = villagersMusic{villagerMusicILoveYou}
	benedictPersonality   = villagersPersonality{villagerPersonalityLazy}
	benedictSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	benedictStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	benedictWallpaper     = villagersWallpaper{villagerWallpaperBackyardFenceWall}
)
