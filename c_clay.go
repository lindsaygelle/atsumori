package atsumori

import "time"

// Clay is an Animal Crossing villager.
var Clay = villager{
	clayAstrology,
	clayBirthDay,
	clayBirthMonth,
	clayBubbleColor,
	clayCategory,
	clayClothes,
	clayClothesColors,
	clayFlooring,
	clayFurniture,
	clayGames,
	clayGender,
	clayInterest,
	clayName,
	clayNameColor,
	clayMusic,
	clayPersonality,
	claySpecies,
	clayStyle,
	clayWallpaper}

var (
	clayAstrology     = villagersAstrology{villagerAstrologyLibra}
	clayBirthDay      = villagersBirthDay{19}
	clayBirthMonth    = villagersBirthMonth{time.October} // 10
	clayBubbleColor   = villagersBubbleColor{villagerBubbleColor7A2500}
	clayCategory      = villagersCategory{villagerCategoryA}
	clayClothes       = villagersClothes{villagerClothesPonchoStyleSweater} // 7859
	clayClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBeige, villagerClothesColorBrown}}
	clayFlooring      = villagersFlooring{villagerFlooringDigSiteFlooring}
	clayFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureHayBed, villagerFurnitureTikiTorch, villagerFurnitureTikiTorch, villagerFurnitureBonfire, villagerFurnitureAmmonite, villagerFurnitureStoneStool, villagerFurnitureTrilobite, villagerFurnitureClassicPitcher, villagerFurniturePondStone, villagerFurnitureStoneStool, villagerFurniturePot}}
	clayGames         = villagersGames{[]VillagerGame{}} // TBD
	clayGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	clayInterest      = villagersInterest{villagerInterestNature}
	clayName          = villagersName{villagerNameClay}
	clayNameColor     = villagersNameColor{villagerNameColorFFFCE9}
	clayMusic         = villagersMusic{villagerMusicKKSafari}
	clayPersonality   = villagersPersonality{villagerPersonalityLazy}
	claySpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesHamster}}
	clayStyle         = villagersStyle{[2]VillagerStyle{villagerStyleElegant, villagerStyleSimple}}
	clayWallpaper     = villagersWallpaper{villagerWallpaperDigSiteWall}
)
