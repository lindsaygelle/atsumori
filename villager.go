package atsumori

var _ Villager = villager{}

// Villager is a representation of an Animal Crossing villager as code.
type Villager interface {
	villagerAstrology
	villagerBirthDay
	villagerBirthMonth
	villagerBubbleColor
	villagerCategory
	villagerClothes
	villagerClothesColor
	villagerFlooring
	villagerFurniture
	villagerGame
	villagerGender
	villagerInterest
	villagerNameColor
	villagerName
	villagerMusic
	villagerPersonality
	villagerSpecies
	villagerStyle
	villagerWallpaper
}

type villager struct {
	villagersAstrology
	villagersBirthDay
	villagersBirthMonth
	villagersBubbleColor
	villagersCategory
	villagersClothes
	villagersClothesColors
	villagersFlooring
	villagersFurniture
	villagersGames
	villagersGender
	villagersInterest
	villagersName
	villagersNameColor
	villagersMusic
	villagersPersonality
	villagersSpecies
	villagersStyle
	villagersWallpaper
}

var (
	// Villagers is a collection of Animal Crossing villagers.
	Villagers = [...]Villager{
		// A
		Ace,
		Admiral,
		AgentS,
		Agnes,
		Aisle,
		Al,
		Alfonso,
		Alice,
		Alli,
		Amelia,
		Anabelle,
		Anchovy,
		Angus,
		Anicotti,
		Ankha,
		Annalisa,
		Annalise,
		Antonio,
		Apollo,
		Apple,
		Astrid,
		Audie,
		Aurora,
		Ava,
		Avery,
		Axel,
		// B
		Baabara,
		Bam,
		Bangle,
		Barold,
		Bea,
		Beardo,
		Beau,
		Becky,
		Bella,
		Benedict,
		Benjamin,
		Bertha,
		Bettina,
		Bianca,
		Biff,
		BigTop,
		Bill,
		Billy,
		Biskit,
		Bitty,
		Blaire,
		Blanche,
		Bluebear,
		Bob,
		Bonbon,
		Bones,
		Boomer,
		Boone,
		Boots,
		Boris,
		Boyd,
		Bree,
		Broccolo,
		Broffina,
		Bruce,
		Bubbles,
		Buck,
		Bud,
		Bunnie,
		Butch,
		Buzz,
		// C
		Cally,
		Camofrog}
)
