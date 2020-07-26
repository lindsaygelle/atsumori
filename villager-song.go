package main

// VillagerSong is an enum of Animal Crossing songs performed by K.K. Slider.
type VillagerSong uint

func (v VillagerSong) String() string {
	return (villagerSongs[v])
}

// villagerSong is a composable struct.
type villagerSong struct {

	// Song is the Animal Crossing villagers favourite K.K Slider song.
	Song VillagerSong `json:"song"`
}

const (
	agentKK VillagerSong = iota
	alohaKK
	animalCity
	bubblegumKK
	cafeKK
	comradeKK
	dJKK
	forestLife
	goKKRider
	hypnoKK
	iLoveYou
	imperialKK
	kkAdventure
	kkAria
	kkBallad
	kkBazaar
	kkBirthday
	kkBlues
	kkBossa
	kkCalypso
	kkCasbah
	kkChorale
	kkCondor
	kkCountry
	kkCruisin
	kkDAndB
	kkDirge
	kkDisco
	kkDixie
	kkFaire
	kkFlamenco
	kkFolk
	kkFusion
	kkGroove
	kkGumbo
	kkHouse
	kkIsland
	kkJazz
	kkJongara
	kkLament
	kkLoveSong
	kkLullaby
	kkMambo
	kkMarathon
	kkMarch
	kkMariachi
	kkMetal
	kkMilonga
	kkMoody
	kkOasis
	kkParade
	kkRagtime
	kkRally
	kkReggae
	kkRock
	kkRockabilly
	kkSafari
	kkSalsa
	kkSamba
	kkSka
	kkSonata
	kkSong
	kkSoul
	kkSteppe
	kkStroll
	kkSwing
	kkSynth
	kkTango
	kkTechnopop
	kkWaltz
	kkWestern
	kkEtude
	kingKK
	luckyKK
	marineSong2001
	mountainSong
	mrKK
	myPlace
	neapolitan
	onlyMe
	pondering
	rockinKK
	soulfulKK
	spaceKK
	springBlossoms
	staleCupcakes
	steepHill
	surfinKK
	theKFunk
	totheEdge
	twoDaysAgo
	wandering
)

var villagerSongs = [...]string{
	"Agent K.K.",
	"Aloha K.K.",
	"Animal City",
	"Bubblegum K.K.",
	"Café K.K.",
	"Comrade K.K.",
	"DJ K.K. (song)",
	"Forest Life",
	"Go K.K. Rider",
	"Hypno K.K.",
	"I Love You",
	"Imperial K.K.",
	"K.K. Adventure",
	"K.K. Aria",
	"K.K. Ballad",
	"K.K. Bazaar",
	"K.K. Birthday",
	"K.K. Blues",
	"K.K. Bossa",
	"K.K. Calypso",
	"K.K. Casbah",
	"K.K. Chorale",
	"K.K. Condor",
	"K.K. Country",
	"K.K. Cruisin'",
	"K.K. D&B",
	"K.K. Dirge",
	"K.K. Disco",
	"K.K. Dixie",
	"K.K. Faire",
	"K.K. Flamenco",
	"K.K. Folk",
	"K.K. Fusion",
	"K.K. Groove",
	"K.K. Gumbo",
	"K.K. House",
	"K.K. Island",
	"K.K. Jazz",
	"K.K. Jongara",
	"K.K. Lament",
	"K.K. Love Song",
	"K.K. Lullaby",
	"K.K. Mambo",
	"K.K. Marathon",
	"K.K. March",
	"K.K. Mariachi",
	"K.K. Metal",
	"K.K. Milonga",
	"K.K. Moody",
	"K.K. Oasis",
	"K.K. Parade",
	"K.K. Ragtime",
	"K.K. Rally",
	"K.K. Reggae",
	"K.K. Rock",
	"K.K. Rockabilly",
	"K.K. Safari",
	"K.K. Salsa",
	"K.K. Samba",
	"K.K. Ska",
	"K.K. Sonata",
	"K.K. Song",
	"K.K. Soul",
	"K.K. Steppe",
	"K.K. Stroll",
	"K.K. Swing",
	"K.K. Synth",
	"K.K. Tango",
	"K.K. Technopop",
	"K.K. Waltz",
	"K.K. Western",
	"K.K. Étude",
	"King K.K.",
	"Lucky K.K.",
	"Marine Song 2001",
	"Mountain Song",
	"Mr. K.K.",
	"My Place",
	"Neapolitan",
	"Only Me",
	"Pondering",
	"Rockin' K.K.",
	"Soulful K.K.",
	"Space K.K.",
	"Spring Blossoms",
	"Stale Cupcakes",
	"Steep Hill",
	"Surfin' K.K.",
	"The K. Funk",
	"To the Edge",
	"Two Days Ago",
	"Wandering"}
