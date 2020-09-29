package atsumori

import "time"

// TBone is an Animal Crossing villager.
var TBone = villager{
	tBoneAstrology,
	tBoneBirthDay,
	tBoneBirthMonth,
	tBoneBubbleColor,
	tBoneCategory,
	tBoneClothes,
	tBoneClothesColors,
	tBoneFlooring,
	tBoneFurniture,
	tBoneGames,
	tBoneGender,
	tBoneInterest,
	tBoneName,
	tBoneNameColor,
	tBoneMusic,
	tBonePersonality,
	tBoneSpecies,
	tBoneStyle,
	tBoneWallpaper}

var (
	tBoneAstrology     = villagersAstrology{villagerAstrologyTaurus}
	tBoneBirthDay      = villagersBirthDay{20}
	tBoneBirthMonth    = villagersBirthMonth{time.May} // 5
	tBoneBubbleColor   = villagersBubbleColor{villagerBubbleColorACC8CF}
	tBoneCategory      = villagersCategory{villagerCategoryB}
	tBoneClothes       = villagersClothes{villagerClothesHantenJacket} // 3164
	tBoneClothesColors = villagersClothesColors{[2]VillagerClothesColor{villagerClothesColorBlue, villagerClothesColorBlack}}
	tBoneFlooring      = villagersFlooring{villagerFlooringArabesqueFlooring}
	tBoneFurniture     = villagersFurniture{[]VillagerFurniture{villagerFurnitureAntiqueBed, villagerFurnitureDoubleSofa, villagerFurnitureAntiqueBureau, villagerFurnitureHangingTerrarium, villagerFurnitureMonstera, villagerFurnitureFireplace, villagerFurnitureMatryoshka, villagerFurnitureAntiqueMiniTable, villagerFurniturePhonograph, villagerFurnitureCello}}
	tBoneGames         = villagersGames{[]VillagerGame{}} // TBD
	tBoneGender        = villagersGender{[2]VillagerGender{villagerGenderMale}}
	tBoneInterest      = villagersInterest{villagerInterestEducation}
	tBoneName          = villagersName{villagerNameTBone}
	tBoneNameColor     = villagersNameColor{villagerNameColor498992}
	tBoneMusic         = villagersMusic{villagerMusicKKSteppe}
	tBonePersonality   = villagersPersonality{villagerPersonalityCranky}
	tBoneSpecies       = villagersSpecies{[2]VillagerSpecies{villagerSpeciesBull}}
	tBoneStyle         = villagersStyle{[2]VillagerStyle{villagerStyleCool, villagerStyleSimple}}
	tBoneWallpaper     = villagersWallpaper{villagerWallpaperArchedWindowWall}
)
