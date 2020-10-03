package atsumori

import "time"

// Wade is an Animal Crossing villager.
var Wade = villager{
	wadeAstrology,
	wadeBirthDay,
	wadeBirthMonth,
	wadeBubbleColor,
	wadeCategory,
	wadeClothes,
	wadeClothesColors,
	wadeFlooring,
	wadeFurniture,
	wadeGames,
	wadeGender,
	wadeInterest,
	wadeName,
	wadeNameColor,
	wadeMusic,
	wadePersonality,
	wadeSpecies,
	wadeStyle,
	wadeWallpaper}

var (
	wadeAstrology     = villagersAstrology{villagerAstrologyScorpio}
	wadeBirthDay      = villagersBirthDay{30}
	wadeBirthMonth    = villagersBirthMonth{time.October} // 10
	wadeBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	wadeCategory      = villagersCategory{villagerCategoryA}
	wadeClothes       = villagersClothes{villagerClothesFrogTee} // 8195
	wadeClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorLightBlue, villagerClothesColorBlue}}
	wadeFlooring      = villagersFlooring{villagerFlooringIceFlooring}
	wadeFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureFrozenChair, villagerFurnitureShellLamp, villagerFurnitureFrozenTable, villagerFurnitureShellFountain, villagerFurnitureShellPartition, villagerFurnitureSeaButterfly, villagerFurnitureShellSpeaker, villagerFurnitureOliveFlounder, villagerFurnitureSeaBass, villagerFurnitureTuna, villagerFurnitureDab, villagerFurnitureSquid, villagerFurnitureHorseMackerel, villagerFurnitureRedSnapper}}
	wadeGames         = villagersGames{[]VillagerGame{}} // TBD
	wadeGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	wadeInterest      = villagersInterest{villagerInterestNature}
	wadeName          = villagersName{villagerNameWade}
	wadeNameColor     = villagersNameColor{villagerNameColor848484}
	wadeMusic         = villagersMusic{villagerMusicStaleCupcakes}
	wadePersonality   = villagersPersonality{villagerPersonalityLazy}
	wadeSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesPenguin}}
	wadeStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleGorgeous}}
	wadeWallpaper     = villagersWallpaper{villagerWallpaperIceWall}
)
