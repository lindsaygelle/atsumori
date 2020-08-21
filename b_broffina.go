package atsumori

import "time"

// Broffina is an Animal Crossing villager
var Broffina = villager{
	broffinaAstrology,
	broffinaBirthDay,
	broffinaBirthMonth,
	broffinaBubbleColor,
	broffinaCategory,
	broffinaClothes,
	broffinaClothesColors,
	broffinaFlooring,
	broffinaFurniture,
	broffinaGames,
	broffinaGender,
	broffinaInterest,
	broffinaName,
	broffinaNameColor,
	broffinaMusic,
	broffinaPersonality,
	broffinaSpecies,
	broffinaStyle,
	broffinaWallpaper}

var (
	broffinaAstrology     = villagersAstrology{villagerAstrologyScorpio}
	broffinaBirthDay      = villagersBirthDay{24}
	broffinaBirthMonth    = villagersBirthMonth{time.October} // 10
	broffinaBubbleColor   = villagersBubbleColor{villagerBubbleColorFF791F}
	broffinaCategory      = villagersCategory{villagerCategoryA}
	broffinaClothes       = villagersClothes{} // 5645
	broffinaClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlack, villagerClothesColorRed}}
	broffinaFlooring      = villagersFlooring{villagerFlooringRosewoodFlooring}
	broffinaFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureWhirlpoolBath, villagerFurnitureAntiqueConsoleTable, villagerFurnitureUprightPiano, villagerFurnitureVelvetStool, villagerFurnitureWoodBurningStove, villagerFurnitureFanPalm, villagerFurnitureAntiqueMiniTable, villagerFurnitureShowerSet, villagerFurniturePhonograph, villagerFurnitureWallMountedTV50In, villagerFurnitureRoseBed, villagerFurnitureAntiqueClock, villagerFurnitureFloralSwag}}
	broffinaGames         = villagersGames{[]VillagerGame{}} // TBD
	broffinaGender        = villagersGender{[2]VillagerGender{villagerGenderFemale}}
	broffinaInterest      = villagersInterest{villagerInterestMusic}
	broffinaName          = villagersName{villagerNameBroffina}
	broffinaNameColor     = villagersNameColor{villagerNameColorFFF2BB}
	broffinaMusic         = villagersMusic{villagerMusicKKFlamenco}
	broffinaPersonality   = villagersPersonality{villagerPersonalitySnooty}
	broffinaSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesChicken}}
	broffinaStyle         = villagersStyle{[2]VillagerStyle{villagerStyleGorgeous, villagerStyleElegant}}
	broffinaWallpaper     = villagersWallpaper{villagerWallpaperPurpleDesertTileWall}
)
