package atsumori

import "time"

// Quillson is an Animal Crossing villager.
var Quillson = villager{
	quillsonAstrology,
	quillsonBirthDay,
	quillsonBirthMonth,
	quillsonBubbleColor,
	quillsonCategory,
	quillsonClothes,
	quillsonClothesColors,
	quillsonFlooring,
	quillsonFurniture,
	quillsonGames,
	quillsonGender,
	quillsonInterest,
	quillsonName,
	quillsonNameColor,
	quillsonMusic,
	quillsonPersonality,
	quillsonSpecies,
	quillsonStyle,
	quillsonWallpaper}

var (
	quillsonAstrology     = villagersAstrology{villagerAstrologyCapricorn}
	quillsonBirthDay      = villagersBirthDay{22}
	quillsonBirthMonth    = villagersBirthMonth{time.December} // 12
	quillsonBubbleColor   = villagersBubbleColor{villagerBubbleColor76B788}
	quillsonCategory      = villagersCategory{villagerCategoryA}
	quillsonClothes       = villagersClothes{villagerClothesCheckeredMuffler} // 4553
	quillsonClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorOrange}}
	quillsonFlooring      = villagersFlooring{villagerFlooringSimplePurpleFlooring}
	quillsonFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureMonstera, villagerFurnitureRattanBed, villagerFurnitureDartboard, villagerFurnitureCoconutWallPlanter, villagerFurnitureSurfboard, villagerFurnitureRattanArmchair, villagerFurnitureRattanLowTable, villagerFurnitureDJsTurntable, villagerFurnitureArowana, villagerFurnitureHighEndStereo}}
	quillsonGames         = villagersGames{[]VillagerGame{}} // TBD
	quillsonGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	quillsonInterest      = villagersInterest{villagerInterestMusic}
	quillsonName          = villagersName{villagerNameQuillson}
	quillsonNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	quillsonMusic         = villagersMusic{villagerMusicKKGroove}
	quillsonPersonality   = villagersPersonality{villagerPersonalitySmug}
	quillsonSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesDuck}}
	quillsonStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleCool}}
	quillsonWallpaper     = villagersWallpaper{villagerWallpaperDarkWoodenMosaicWall}
)
