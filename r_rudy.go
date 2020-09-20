package atsumori

import "time"

// Rudy is an Animal Crossing villager.
var Rudy = villager{
	rudyAstrology,
	rudyBirthDay,
	rudyBirthMonth,
	rudyBubbleColor,
	rudyCategory,
	rudyClothes,
	rudyClothesColors,
	rudyFlooring,
	rudyFurniture,
	rudyGames,
	rudyGender,
	rudyInterest,
	rudyName,
	rudyNameColor,
	rudyMusic,
	rudyPersonality,
	rudySpecies,
	rudyStyle,
	rudyWallpaper}

var (
	rudyAstrology     = villagersAstrology{villagerAstrologySagittarius}
	rudyBirthDay      = villagersBirthDay{20}
	rudyBirthMonth    = villagersBirthMonth{time.December} // 12
	rudyBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	rudyCategory      = villagersCategory{villagerCategoryA}
	rudyClothes       = villagersClothes{villagerClothesSimpleParka} // 7972
	rudyClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorYellow, villagerClothesColorBeige}}
	rudyFlooring      = villagersFlooring{villagerFlooringPineBoardFlooring}
	rudyFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWoodenBlockStereo, villagerFurnitureThrowbackWallClock, villagerFurnitureWoodenSimpleBed, villagerFurnitureThrowbackContainer, villagerFurnitureMiniDIYWorkbench, villagerFurnitureWoodenChest, villagerFurnitureTrainSet, villagerFurnitureMiniCactusSet, villagerFurnitureWoodenMiniTable, villagerFurnitureThrowbackMittChair}}
	rudyGames         = villagersGames{[]VillagerGame{}} // TBD
	rudyGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	rudyInterest      = villagersInterest{villagerInterestPlay}
	rudyName          = villagersName{villagerNameRudy}
	rudyNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	rudyMusic         = villagersMusic{villagerMusicKKBlues}
	rudyPersonality   = villagersPersonality{villagerPersonalityJock}
	rudySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	rudyStyle         = villagersStyle{[2]VillagerStyle{villagerStyleActive, villagerStyleSimple}}
	rudyWallpaper     = villagersWallpaper{villagerWallpaperGreenPlayroomWall}
)
