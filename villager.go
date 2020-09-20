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
		Camofrog,
		Canberra,
		Candi,
		CarmenMouse,
		CarmenRabbit,
		Caroline,
		Carrie,
		Cashmere,
		Celia,
		Cesar,
		Chadder,
		Charlise,
		Cheri,
		Cherry,
		Chester,
		Chevre,
		Chief,
		Chops,
		Chow,
		Chrissy,
		Claude,
		Claudia,
		Clay,
		Cleo,
		Clyde,
		Coach,
		Cobb,
		Coco,
		Cole,
		Colton,
		Cookie,
		Cousteau,
		Cranston,
		Croque,
		Cube,
		Curlos,
		Curly,
		Curt,
		Cyd,
		Cyrano,
		// D
		Daisy,
		Deena,
		Deirdre,
		Del,
		Deli,
		Derwin,
		Diana,
		Diva,
		Dizzy,
		Dobie,
		Doc,
		Dom,
		Dora,
		Dotty,
		Drago,
		Drake,
		Drift,
		// E
		Ed,
		Egbert,
		Elise,
		Ellie,
		Elmer,
		Eloise,
		Elvis,
		Erik,
		Eugene,
		Eunice,
		// F
		Fang,
		Fauna,
		Felicity,
		Filbert,
		Flip,
		Flo,
		Flora,
		Flurry,
		Francine,
		Frank,
		Freckles,
		Freya,
		Friga,
		Frita,
		Frobert,
		Fuchsia,
		// G
		Gabi,
		Gala,
		Gaston,
		Gayle,
		Genji,
		Gigi,
		Gladys,
		Gloria,
		Goldie,
		Gonzo,
		Goose,
		Graham,
		Greta,
		Grizzly,
		Groucho,
		Gruff,
		Gwen,
		// H
		Hamlet,
		Hamphrey,
		Hans,
		Harry,
		Hazel,
		Henry,
		Hippeux,
		Hopkins,
		Hopper,
		Hornsby,
		Huck,
		Hugh,
		// I
		Iggly,
		Ike,
		// J
		Jacob,
		Jacques,
		Jambette,
		Jay,
		Jeremiah,
		Jitters,
		Joey,
		Judy,
		Julia,
		Julian,
		June,
		// K
		Kabuki,
		Katt,
		Keaton,
		Ken,
		Ketchup,
		Kevin,
		KidCat,
		Kidd,
		Kiki,
		Kitt,
		Kitty,
		Klaus,
		Knox,
		Kody,
		Kyle,
		// L
		Leonardo,
		Leopold,
		Lily,
		Limberg,
		Lionel,
		Lobo,
		Lolly,
		Lopez,
		Louie,
		Lucha,
		Lucky,
		Lucy,
		Lyman,
		// M
		Mac,
		Maddie,
		Maelle,
		Maggie,
		Mallary,
		Maple,
		Marcel,
		Marcie,
		Margie,
		Marina,
		Marshal,
		Mathilda,
		Megan,
		Melba,
		Merengue,
		Merry,
		Midge,
		Mint,
		Mira,
		Miranda,
		Mitzi,
		Moe,
		Molly,
		Monique,
		Monty,
		Moose,
		Mott,
		Muffy,
		Murphy,
		// N
		Nan,
		Nana,
		Naomi,
		Nate,
		Nibbles,
		Norma,
		// O
		OHare,
		Octavian,
		Olaf,
		Olive,
		Olivia,
		Opal,
		Ozzie,
		// P
		Pancetti,
		Pango,
		Paolo,
		Papi,
		Pashmina,
		Pate,
		Patty,
		Paula,
		Peaches,
		Peanut,
		Pecan,
		Peck,
		Peewee,
		Peggy,
		Pekoe,
		Penelope,
		Phil,
		Phoebe,
		Pierce,
		Pietro,
		Pinky,
		Piper,
		Pippy,
		Plucky,
		Pompom,
		Poncho,
		Poppy,
		Portia,
		Prince,
		Puck,
		Puddles,
		Pudge,
		Punchy,
		Purrl,
		// Q
		Queenie,
		Quillson,
		// R
		Raddle,
		Rasher,
		Raymond,
		Renee,
		Reneigh,
		Rex,
		Rhonda,
		Ribbot,
		Ricky,
		Rizzo,
		Roald,
		Robin,
		Rocco,
		Rocket,
		Rod,
		Rodeo,
		Rodney,
		Rolf,
		Rooney,
		Rory,
		Roscoe,
		Rosie,
		Rowan}
)
