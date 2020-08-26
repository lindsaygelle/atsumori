package atsumori

import "time"

// Kabuki is an Animal Crossing villager
var Kabuki = villager{
	kabukiAstrology,
	kabukiBirthDay,
	kabukiBirthMonth,
	kabukiBubbleColor,
	kabukiCategory,
	kabukiClothes,
	kabukiClothesColors,
	kabukiFlooring,
	kabukiFurniture,
	kabukiGames,
	kabukiGender,
	kabukiInterest,
	kabukiName,
	kabukiNameColor,
	kabukiMusic,
	kabukiPersonality,
	kabukiSpecies,
	kabukiStyle,
	kabukiWallpaper}

var (
	kabukiAstrology     = villagersAstrology{villagerAstrologySagittarius}
	kabukiBirthDay      = villagersBirthDay{29}
	kabukiBirthMonth    = villagersBirthMonth{time.November} // 11
	kabukiBubbleColor   = villagersBubbleColor{villagerBubbleColorFFFFFF}
	kabukiCategory      = villagersCategory{villagerCategoryB}
	kabukiClothes       = villagersClothes{} // 6024
	kabukiClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorPurple, villagerClothesColorRed}}
	kabukiFlooring      = villagersFlooring{villagerFlooringTatami}
	kabukiFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureBambooPartition, villagerFurnitureBambooPartition, villagerFurnitureSquatToilet, villagerFurnitureFuton, villagerFurnitureKettleBathtub, villagerFurnitureBonsaiShelf, villagerFurniturePaperLantern, villagerFurnitureWoodenPlankSign}}
	kabukiGames         = villagersGames{[]VillagerGame{}} // TBD
	kabukiGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	kabukiInterest      = villagersInterest{villagerInterestMusic}
	kabukiName          = villagersName{villagerNameKabuki}
	kabukiNameColor     = villagersNameColor{villagerNameColor848484}
	kabukiMusic         = villagersMusic{villagerMusicKKJongara}
	kabukiPersonality   = villagersPersonality{villagerPersonalityCranky}
	kabukiSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesCat}}
	kabukiStyle         = villagersStyle{[2]VillagerStyle{villagerStyleSimple, villagerStyleSimple}}
	kabukiWallpaper     = villagersWallpaper{villagerWallpaperDojoWall}
)
