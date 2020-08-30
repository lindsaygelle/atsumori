package atsumori

import "time"

// Hippeux is an Animal Crossing villager.
var Hippeux = villager{
	hippeuxAstrology,
	hippeuxBirthDay,
	hippeuxBirthMonth,
	hippeuxBubbleColor,
	hippeuxCategory,
	hippeuxClothes,
	hippeuxClothesColors,
	hippeuxFlooring,
	hippeuxFurniture,
	hippeuxGames,
	hippeuxGender,
	hippeuxInterest,
	hippeuxName,
	hippeuxNameColor,
	hippeuxMusic,
	hippeuxPersonality,
	hippeuxSpecies,
	hippeuxStyle,
	hippeuxWallpaper}

var (
	hippeuxAstrology     = villagersAstrology{villagerAstrologyLibra}
	hippeuxBirthDay      = villagersBirthDay{15}
	hippeuxBirthMonth    = villagersBirthMonth{time.October} // 10
	hippeuxBubbleColor   = villagersBubbleColor{villagerBubbleColorD8CC39}
	hippeuxCategory      = villagersCategory{villagerCategoryA}
	hippeuxClothes       = villagersClothes{villagerClothesEmblemBlazer} // 8110
	hippeuxClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBrown, villagerClothesColorGreen}}
	hippeuxFlooring      = villagersFlooring{villagerFlooringDarkHerringboneFlooring}
	hippeuxFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureGrandPiano, villagerFurniturePianoBench, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueMiniTable, villagerFurnitureMetronome, villagerFurnitureRattanArmchair, villagerFurnitureRattanBed, villagerFurniturePendulumClock, villagerFurnitureRattanLowTable, villagerFurniturePhonograph, villagerFurnitureRattanEndTable, villagerFurnitureRattanTableLamp}}
	hippeuxGames         = villagersGames{[]VillagerGame{}} // TBD
	hippeuxGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	hippeuxInterest      = villagersInterest{villagerInterestEducation}
	hippeuxName          = villagersName{villagerNameHippeux}
	hippeuxNameColor     = villagersNameColor{villagerNameColor8B5F57}
	hippeuxMusic         = villagersMusic{villagerMusicKKSonata}
	hippeuxPersonality   = villagersPersonality{villagerPersonalitySmug}
	hippeuxSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHippo}}
	hippeuxStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleGorgeous}}
	hippeuxWallpaper     = villagersWallpaper{villagerWallpaperArchedWindowWall}
)
