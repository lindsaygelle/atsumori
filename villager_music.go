package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerMusic(0)

var _ json.Marshaler = VillagerMusic(0)

var _ villagerMusic = villagersMusic{}

// VillagerMusic is an Animal Crossing villagers favourite song.
type VillagerMusic uint8

func (v VillagerMusic) String() string { return villagerMusicAll[v] }

// MarshalJSON returns the encoding of VillagerMusic.
func (v VillagerMusic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerMusic interface {
	Music() string
}

type villagersMusic struct {
	VillagerMusic VillagerMusic `json:"music"`
}

func (v villagersMusic) Music() string { return v.VillagerMusic.String() }

const (
	_villagerMusic               string = _nil
	_villagerMusicAgentKK        string = "Agent K.K."
	_villagerMusicAlohaKK        string = "Aloha K.K."
	_villagerMusicAnimalCity     string = "Animal City"
	_villagerMusicBubblegumKK    string = "Bubblegum K.K."
	_villagerMusicCafeKK         string = "Café K.K."
	_villagerMusicComradeKK      string = "Comrade K.K."
	_villagerMusicDJKK           string = "DJ K.K."
	_villagerMusicDrivin         string = "Drivin'"
	_villagerMusicFarewell       string = "Farewell"
	_villagerMusicForestLife     string = "Forest Life"
	_villagerMusicGoKKRider      string = "Go K.K. Rider"
	_villagerMusicHypnoKK        string = "Hypno K.K."
	_villagerMusicILoveYou       string = "I Love You"
	_villagerMusicImperialKK     string = "Imperial K.K."
	_villagerMusicKKAdventure    string = "K.K. Adventure"
	_villagerMusicKKAria         string = "K.K. Aria"
	_villagerMusicKKBallad       string = "K.K. Ballad"
	_villagerMusicKKBazaar       string = "K.K. Bazaar"
	_villagerMusicKKBlues        string = "K.K. Blues"
	_villagerMusicKKBossa        string = "K.K. Bossa"
	_villagerMusicKKCalypso      string = "K.K. Calypso"
	_villagerMusicKKCasbah       string = "K.K. Casbah"
	_villagerMusicKKChorale      string = "K.K. Chorale"
	_villagerMusicKKCondor       string = "K.K. Condor"
	_villagerMusicKKCountry      string = "K.K. Country"
	_villagerMusicKKCruisin      string = "K.K. Cruisin'"
	_villagerMusicKKDB           string = "K.K. D&B"
	_villagerMusicKKDirge        string = "K.K. Dirge"
	_villagerMusicKKDisco        string = "K.K. Disco"
	_villagerMusicKKDixie        string = "K.K. Dixie"
	_villagerMusicKKFaire        string = "K.K. Faire"
	_villagerMusicKKFlamenco     string = "K.K. Flamenco"
	_villagerMusicKKFolk         string = "K.K. Folk"
	_villagerMusicKKFusion       string = "K.K. Fusion"
	_villagerMusicKKGroove       string = "K.K. Groove"
	_villagerMusicKKGumbo        string = "K.K. Gumbo"
	_villagerMusicKKHouse        string = "K.K. House"
	_villagerMusicKKIsland       string = "K.K. Island"
	_villagerMusicKKJazz         string = "K.K. Jazz"
	_villagerMusicKKJongara      string = "K.K. Jongara"
	_villagerMusicKKLament       string = "K.K. Lament"
	_villagerMusicKKLoveSong     string = "K.K. Love Song"
	_villagerMusicKKLullaby      string = "K.K. Lullaby"
	_villagerMusicKKMambo        string = "K.K. Mambo"
	_villagerMusicKKMarathon     string = "K.K. Marathon"
	_villagerMusicKKMarch        string = "K.K. March"
	_villagerMusicKKMetal        string = "K.K. Metal"
	_villagerMusicKKMilonga      string = "K.K. Milonga"
	_villagerMusicKKMoody        string = "K.K. Moody"
	_villagerMusicKKOasis        string = "K.K. Oasis"
	_villagerMusicKKParade       string = "K.K. Parade"
	_villagerMusicKKRagtime      string = "K.K. Ragtime"
	_villagerMusicKKRally        string = "K.K. Rally"
	_villagerMusicKKReggae       string = "K.K. Reggae"
	_villagerMusicKKRock         string = "K.K. Rock"
	_villagerMusicKKRockabilly   string = "K.K. Rockabilly"
	_villagerMusicKKSafari       string = "K.K. Safari"
	_villagerMusicKKSalsa        string = "K.K. Salsa"
	_villagerMusicKKSonata       string = "K.K. Sonata"
	_villagerMusicKKSong         string = "K.K. Song"
	_villagerMusicKKSoul         string = "K.K. Soul"
	_villagerMusicKKSteppe       string = "K.K. Steppe"
	_villagerMusicKKStroll       string = "K.K. Stroll"
	_villagerMusicKKSwing        string = "K.K. Swing"
	_villagerMusicKKSynth        string = "K.K. Synth"
	_villagerMusicKKTango        string = "K.K. Tango"
	_villagerMusicKKTechnopop    string = "K.K. Technopop"
	_villagerMusicKKWestern      string = "K.K. Western"
	_villagerMusicKKEtude        string = "K.K. Étude"
	_villagerMusicKingKK         string = "King K.K."
	_villagerMusicLuckyKK        string = "Lucky K.K."
	_villagerMusicMarineSong2001 string = "Marine Song 2001"
	_villagerMusicMountainSong   string = "Mountain Song"
	_villagerMusicMrKK           string = "Mr. K.K."
	_villagerMusicMyPlace        string = "My Place"
	_villagerMusicNeapolitan     string = "Neapolitan"
	_villagerMusicOnlyMe         string = "Only Me"
	_villagerMusicPondering      string = "Pondering"
	_villagerMusicRockinKK       string = "Rockin' K.K."
	_villagerMusicSoulfulKK      string = "Soulful K.K."
	_villagerMusicSpaceKK        string = "Space K.K."
	_villagerMusicSpringBlossoms string = "Spring Blossoms"
	_villagerMusicStaleCupcakes  string = "Stale Cupcakes"
	_villagerMusicSteepHill      string = "Steep Hill"
	_villagerMusicSurfinKK       string = "Surfin' K.K."
	_villagerMusicTheKFunk       string = "The K. Funk"
	_villagerMusicToTheEdge      string = "To The Edge"
	_villagerMusicTwoDaysAgo     string = "Two Days Ago"
	_villagerMusicWandering      string = "Wandering"
)

const (
	villagerMusicAgentKK VillagerMusic = iota + 1
	villagerMusicAlohaKK
	villagerMusicAnimalCity
	villagerMusicBubblegumKK
	villagerMusicCafeKK
	villagerMusicComradeKK
	villagerMusicDJKK
	villagerMusicDrivin
	villagerMusicFarewell
	villagerMusicForestLife
	villagerMusicGoKKRider
	villagerMusicHypnoKK
	villagerMusicILoveYou
	villagerMusicImperialKK
	villagerMusicKKAdventure
	villagerMusicKKAria
	villagerMusicKKBallad
	villagerMusicKKBazaar
	villagerMusicKKBlues
	villagerMusicKKBossa
	villagerMusicKKCalypso
	villagerMusicKKCasbah
	villagerMusicKKChorale
	villagerMusicKKCondor
	villagerMusicKKCountry
	villagerMusicKKCruisin
	villagerMusicKKDB
	villagerMusicKKDirge
	villagerMusicKKDisco
	villagerMusicKKDixie
	villagerMusicKKFaire
	villagerMusicKKFlamenco
	villagerMusicKKFolk
	villagerMusicKKFusion
	villagerMusicKKGroove
	villagerMusicKKGumbo
	villagerMusicKKHouse
	villagerMusicKKIsland
	villagerMusicKKJazz
	villagerMusicKKJongara
	villagerMusicKKLament
	villagerMusicKKLoveSong
	villagerMusicKKLullaby
	villagerMusicKKMambo
	villagerMusicKKMarathon
	villagerMusicKKMarch
	villagerMusicKKMetal
	villagerMusicKKMilonga
	villagerMusicKKMoody
	villagerMusicKKOasis
	villagerMusicKKParade
	villagerMusicKKRagtime
	villagerMusicKKRally
	villagerMusicKKReggae
	villagerMusicKKRock
	villagerMusicKKRockabilly
	villagerMusicKKSafari
	villagerMusicKKSalsa
	villagerMusicKKSonata
	villagerMusicKKSong
	villagerMusicKKSoul
	villagerMusicKKSteppe
	villagerMusicKKStroll
	villagerMusicKKSwing
	villagerMusicKKSynth
	villagerMusicKKTango
	villagerMusicKKTechnopop
	villagerMusicKKWestern
	villagerMusicKKEtude
	villagerMusicKingKK
	villagerMusicLuckyKK
	villagerMusicMarineSong2001
	villagerMusicMountainSong
	villagerMusicMrKK
	villagerMusicMyPlace
	villagerMusicNeapolitan
	villagerMusicOnlyMe
	villagerMusicPondering
	villagerMusicRockinKK
	villagerMusicSoulfulKK
	villagerMusicSpaceKK
	villagerMusicSpringBlossoms
	villagerMusicStaleCupcakes
	villagerMusicSteepHill
	villagerMusicSurfinKK
	villagerMusicTheKFunk
	villagerMusicToTheEdge
	villagerMusicTwoDaysAgo
	villagerMusicWandering
)

var (
	villagerMusicAll = [...]string{
		_villagerMusic,
		_villagerMusicAgentKK,
		_villagerMusicAlohaKK,
		_villagerMusicAnimalCity,
		_villagerMusicBubblegumKK,
		_villagerMusicCafeKK,
		_villagerMusicComradeKK,
		_villagerMusicDJKK,
		_villagerMusicDrivin,
		_villagerMusicFarewell,
		_villagerMusicForestLife,
		_villagerMusicGoKKRider,
		_villagerMusicHypnoKK,
		_villagerMusicILoveYou,
		_villagerMusicImperialKK,
		_villagerMusicKKAdventure,
		_villagerMusicKKAria,
		_villagerMusicKKBallad,
		_villagerMusicKKBazaar,
		_villagerMusicKKBlues,
		_villagerMusicKKBossa,
		_villagerMusicKKCalypso,
		_villagerMusicKKCasbah,
		_villagerMusicKKChorale,
		_villagerMusicKKCondor,
		_villagerMusicKKCountry,
		_villagerMusicKKCruisin,
		_villagerMusicKKDB,
		_villagerMusicKKDirge,
		_villagerMusicKKDisco,
		_villagerMusicKKDixie,
		_villagerMusicKKFaire,
		_villagerMusicKKFlamenco,
		_villagerMusicKKFolk,
		_villagerMusicKKFusion,
		_villagerMusicKKGroove,
		_villagerMusicKKGumbo,
		_villagerMusicKKHouse,
		_villagerMusicKKIsland,
		_villagerMusicKKJazz,
		_villagerMusicKKJongara,
		_villagerMusicKKLament,
		_villagerMusicKKLoveSong,
		_villagerMusicKKLullaby,
		_villagerMusicKKMambo,
		_villagerMusicKKMarathon,
		_villagerMusicKKMarch,
		_villagerMusicKKMetal,
		_villagerMusicKKMilonga,
		_villagerMusicKKMoody,
		_villagerMusicKKOasis,
		_villagerMusicKKParade,
		_villagerMusicKKRagtime,
		_villagerMusicKKRally,
		_villagerMusicKKReggae,
		_villagerMusicKKRock,
		_villagerMusicKKRockabilly,
		_villagerMusicKKSafari,
		_villagerMusicKKSalsa,
		_villagerMusicKKSonata,
		_villagerMusicKKSong,
		_villagerMusicKKSoul,
		_villagerMusicKKSteppe,
		_villagerMusicKKStroll,
		_villagerMusicKKSwing,
		_villagerMusicKKSynth,
		_villagerMusicKKTango,
		_villagerMusicKKTechnopop,
		_villagerMusicKKWestern,
		_villagerMusicKKEtude,
		_villagerMusicKingKK,
		_villagerMusicLuckyKK,
		_villagerMusicMarineSong2001,
		_villagerMusicMountainSong,
		_villagerMusicMrKK,
		_villagerMusicMyPlace,
		_villagerMusicNeapolitan,
		_villagerMusicOnlyMe,
		_villagerMusicPondering,
		_villagerMusicRockinKK,
		_villagerMusicSoulfulKK,
		_villagerMusicSpaceKK,
		_villagerMusicSpringBlossoms,
		_villagerMusicStaleCupcakes,
		_villagerMusicSteepHill,
		_villagerMusicSurfinKK,
		_villagerMusicTheKFunk,
		_villagerMusicToTheEdge,
		_villagerMusicTwoDaysAgo,
		_villagerMusicWandering}
)
