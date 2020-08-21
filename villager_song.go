package atsumori

import (
	"encoding/json"
	"fmt"
)

var _ fmt.Stringer = VillagerSong(0)

var _ json.Marshaler = VillagerSong(0)

var _ villagerSong = villagersSong{}

// VillagerSong is an Animal Crossing villagers favourite song.
type VillagerSong uint8

func (v VillagerSong) String() string { return villagerSongAll[v] }

// MarshalJSON returns the encoding of VillagerSong.
func (v VillagerSong) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

type villagerSong interface {
	Song() string
}

type villagersSong struct {
	VillagerSong VillagerSong `json:"song"`
}

func (v villagersSong) Song() string { return v.VillagerSong.String() }

const (
	_villagerSong               string = _nil
	_villagerSongAgentKK        string = "Agent K.K."
	_villagerSongAlohaKK        string = "Aloha K.K."
	_villagerSongAnimalCity     string = "Animal City"
	_villagerSongBubblegumKK    string = "Bubblegum K.K."
	_villagerSongCafeKK         string = "Café K.K."
	_villagerSongComradeKK      string = "Comrade K.K."
	_villagerSongDJKK           string = "DJ K.K."
	_villagerSongDrivin         string = "Drivin'"
	_villagerSongFarewell       string = "Farewell"
	_villagerSongForestLife     string = "Forest Life"
	_villagerSongGoKKRider      string = "Go K.K. Rider"
	_villagerSongHypnoKK        string = "Hypno K.K."
	_villagerSongILoveYou       string = "I Love You"
	_villagerSongImperialKK     string = "Imperial K.K."
	_villagerSongKKAdventure    string = "K.K. Adventure"
	_villagerSongKKAria         string = "K.K. Aria"
	_villagerSongKKBallad       string = "K.K. Ballad"
	_villagerSongKKBazaar       string = "K.K. Bazaar"
	_villagerSongKKBlues        string = "K.K. Blues"
	_villagerSongKKBossa        string = "K.K. Bossa"
	_villagerSongKKCalypso      string = "K.K. Calypso"
	_villagerSongKKCasbah       string = "K.K. Casbah"
	_villagerSongKKChorale      string = "K.K. Chorale"
	_villagerSongKKCondor       string = "K.K. Condor"
	_villagerSongKKCountry      string = "K.K. Country"
	_villagerSongKKCruisin      string = "K.K. Cruisin'"
	_villagerSongKKDB           string = "K.K. D&B"
	_villagerSongKKDirge        string = "K.K. Dirge"
	_villagerSongKKDisco        string = "K.K. Disco"
	_villagerSongKKDixie        string = "K.K. Dixie"
	_villagerSongKKFaire        string = "K.K. Faire"
	_villagerSongKKFlamenco     string = "K.K. Flamenco"
	_villagerSongKKFolk         string = "K.K. Folk"
	_villagerSongKKFusion       string = "K.K. Fusion"
	_villagerSongKKGroove       string = "K.K. Groove"
	_villagerSongKKGumbo        string = "K.K. Gumbo"
	_villagerSongKKHouse        string = "K.K. House"
	_villagerSongKKIsland       string = "K.K. Island"
	_villagerSongKKJazz         string = "K.K. Jazz"
	_villagerSongKKJongara      string = "K.K. Jongara"
	_villagerSongKKLament       string = "K.K. Lament"
	_villagerSongKKLoveSong     string = "K.K. Love Song"
	_villagerSongKKLullaby      string = "K.K. Lullaby"
	_villagerSongKKMambo        string = "K.K. Mambo"
	_villagerSongKKMarathon     string = "K.K. Marathon"
	_villagerSongKKMarch        string = "K.K. March"
	_villagerSongKKMetal        string = "K.K. Metal"
	_villagerSongKKMilonga      string = "K.K. Milonga"
	_villagerSongKKMoody        string = "K.K. Moody"
	_villagerSongKKOasis        string = "K.K. Oasis"
	_villagerSongKKParade       string = "K.K. Parade"
	_villagerSongKKRagtime      string = "K.K. Ragtime"
	_villagerSongKKRally        string = "K.K. Rally"
	_villagerSongKKReggae       string = "K.K. Reggae"
	_villagerSongKKRock         string = "K.K. Rock"
	_villagerSongKKRockabilly   string = "K.K. Rockabilly"
	_villagerSongKKSafari       string = "K.K. Safari"
	_villagerSongKKSalsa        string = "K.K. Salsa"
	_villagerSongKKSonata       string = "K.K. Sonata"
	_villagerSongKKSong         string = "K.K. Song"
	_villagerSongKKSoul         string = "K.K. Soul"
	_villagerSongKKSteppe       string = "K.K. Steppe"
	_villagerSongKKStroll       string = "K.K. Stroll"
	_villagerSongKKSwing        string = "K.K. Swing"
	_villagerSongKKSynth        string = "K.K. Synth"
	_villagerSongKKTango        string = "K.K. Tango"
	_villagerSongKKTechnopop    string = "K.K. Technopop"
	_villagerSongKKWestern      string = "K.K. Western"
	_villagerSongKKEtude        string = "K.K. Étude"
	_villagerSongKingKK         string = "King K.K."
	_villagerSongLuckyKK        string = "Lucky K.K."
	_villagerSongMarineSong2001 string = "Marine Song 2001"
	_villagerSongMountainSong   string = "Mountain Song"
	_villagerSongMrKK           string = "Mr. K.K."
	_villagerSongMyPlace        string = "My Place"
	_villagerSongNeapolitan     string = "Neapolitan"
	_villagerSongOnlyMe         string = "Only Me"
	_villagerSongPondering      string = "Pondering"
	_villagerSongRockinKK       string = "Rockin' K.K."
	_villagerSongSoulfulKK      string = "Soulful K.K."
	_villagerSongSpaceKK        string = "Space K.K."
	_villagerSongSpringBlossoms string = "Spring Blossoms"
	_villagerSongStaleCupcakes  string = "Stale Cupcakes"
	_villagerSongSteepHill      string = "Steep Hill"
	_villagerSongSurfinKK       string = "Surfin' K.K."
	_villagerSongTheKFunk       string = "The K. Funk"
	_villagerSongToTheEdge      string = "To The Edge"
	_villagerSongTwoDaysAgo     string = "Two Days Ago"
	_villagerSongWandering      string = "Wandering"
)

const (
	villagerSongAgentKK VillagerSong = iota + 1
	villagerSongAlohaKK
	villagerSongAnimalCity
	villagerSongBubblegumKK
	villagerSongCafeKK
	villagerSongComradeKK
	villagerSongDJKK
	villagerSongDrivin
	villagerSongFarewell
	villagerSongForestLife
	villagerSongGoKKRider
	villagerSongHypnoKK
	villagerSongILoveYou
	villagerSongImperialKK
	villagerSongKKAdventure
	villagerSongKKAria
	villagerSongKKBallad
	villagerSongKKBazaar
	villagerSongKKBlues
	villagerSongKKBossa
	villagerSongKKCalypso
	villagerSongKKCasbah
	villagerSongKKChorale
	villagerSongKKCondor
	villagerSongKKCountry
	villagerSongKKCruisin
	villagerSongKKDB
	villagerSongKKDirge
	villagerSongKKDisco
	villagerSongKKDixie
	villagerSongKKFaire
	villagerSongKKFlamenco
	villagerSongKKFolk
	villagerSongKKFusion
	villagerSongKKGroove
	villagerSongKKGumbo
	villagerSongKKHouse
	villagerSongKKIsland
	villagerSongKKJazz
	villagerSongKKJongara
	villagerSongKKLament
	villagerSongKKLoveSong
	villagerSongKKLullaby
	villagerSongKKMambo
	villagerSongKKMarathon
	villagerSongKKMarch
	villagerSongKKMetal
	villagerSongKKMilonga
	villagerSongKKMoody
	villagerSongKKOasis
	villagerSongKKParade
	villagerSongKKRagtime
	villagerSongKKRally
	villagerSongKKReggae
	villagerSongKKRock
	villagerSongKKRockabilly
	villagerSongKKSafari
	villagerSongKKSalsa
	villagerSongKKSonata
	villagerSongKKSong
	villagerSongKKSoul
	villagerSongKKSteppe
	villagerSongKKStroll
	villagerSongKKSwing
	villagerSongKKSynth
	villagerSongKKTango
	villagerSongKKTechnopop
	villagerSongKKWestern
	villagerSongKKEtude
	villagerSongKingKK
	villagerSongLuckyKK
	villagerSongMarineSong2001
	villagerSongMountainSong
	villagerSongMrKK
	villagerSongMyPlace
	villagerSongNeapolitan
	villagerSongOnlyMe
	villagerSongPondering
	villagerSongRockinKK
	villagerSongSoulfulKK
	villagerSongSpaceKK
	villagerSongSpringBlossoms
	villagerSongStaleCupcakes
	villagerSongSteepHill
	villagerSongSurfinKK
	villagerSongTheKFunk
	villagerSongToTheEdge
	villagerSongTwoDaysAgo
	villagerSongWandering
)

var (
	villagerSongAll = [...]string{
		_villagerSong,
		_villagerSongAgentKK,
		_villagerSongAlohaKK,
		_villagerSongAnimalCity,
		_villagerSongBubblegumKK,
		_villagerSongCafeKK,
		_villagerSongComradeKK,
		_villagerSongDJKK,
		_villagerSongDrivin,
		_villagerSongFarewell,
		_villagerSongForestLife,
		_villagerSongGoKKRider,
		_villagerSongHypnoKK,
		_villagerSongILoveYou,
		_villagerSongImperialKK,
		_villagerSongKKAdventure,
		_villagerSongKKAria,
		_villagerSongKKBallad,
		_villagerSongKKBazaar,
		_villagerSongKKBlues,
		_villagerSongKKBossa,
		_villagerSongKKCalypso,
		_villagerSongKKCasbah,
		_villagerSongKKChorale,
		_villagerSongKKCondor,
		_villagerSongKKCountry,
		_villagerSongKKCruisin,
		_villagerSongKKDB,
		_villagerSongKKDirge,
		_villagerSongKKDisco,
		_villagerSongKKDixie,
		_villagerSongKKFaire,
		_villagerSongKKFlamenco,
		_villagerSongKKFolk,
		_villagerSongKKFusion,
		_villagerSongKKGroove,
		_villagerSongKKGumbo,
		_villagerSongKKHouse,
		_villagerSongKKIsland,
		_villagerSongKKJazz,
		_villagerSongKKJongara,
		_villagerSongKKLament,
		_villagerSongKKLoveSong,
		_villagerSongKKLullaby,
		_villagerSongKKMambo,
		_villagerSongKKMarathon,
		_villagerSongKKMarch,
		_villagerSongKKMetal,
		_villagerSongKKMilonga,
		_villagerSongKKMoody,
		_villagerSongKKOasis,
		_villagerSongKKParade,
		_villagerSongKKRagtime,
		_villagerSongKKRally,
		_villagerSongKKReggae,
		_villagerSongKKRock,
		_villagerSongKKRockabilly,
		_villagerSongKKSafari,
		_villagerSongKKSalsa,
		_villagerSongKKSonata,
		_villagerSongKKSong,
		_villagerSongKKSoul,
		_villagerSongKKSteppe,
		_villagerSongKKStroll,
		_villagerSongKKSwing,
		_villagerSongKKSynth,
		_villagerSongKKTango,
		_villagerSongKKTechnopop,
		_villagerSongKKWestern,
		_villagerSongKKEtude,
		_villagerSongKingKK,
		_villagerSongLuckyKK,
		_villagerSongMarineSong2001,
		_villagerSongMountainSong,
		_villagerSongMrKK,
		_villagerSongMyPlace,
		_villagerSongNeapolitan,
		_villagerSongOnlyMe,
		_villagerSongPondering,
		_villagerSongRockinKK,
		_villagerSongSoulfulKK,
		_villagerSongSpaceKK,
		_villagerSongSpringBlossoms,
		_villagerSongStaleCupcakes,
		_villagerSongSteepHill,
		_villagerSongSurfinKK,
		_villagerSongTheKFunk,
		_villagerSongToTheEdge,
		_villagerSongTwoDaysAgo,
		_villagerSongWandering}
)
