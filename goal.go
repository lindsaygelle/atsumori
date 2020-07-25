package main

// Goal is an enum of Animal Crossing villager goals.
type Goal uint

func (g Goal) String() string {
	return (goals[g])
}

// goal is a composable struct.
type goal struct {

	// Goal is the career ambition of an Animal Crossing villager.
	Goal Goal `json:"goal"`
}

const (
	actor Goal = iota
	aesthetician
	archaeologist
	architect
	artist
	astronaut
	baseballPlayer
	basketballPlayer
	boxer
	ceo
	carpenter
	cartoonist
	chef
	dancer
	dentist
	dermatologist
	designer
	detective
	doctor
	drummer
	esthetician
	explorer
	fashionDesigner
	figureSkater
	firefighter
	fisherman
	florist
	forestRanger
	guitarist
	gymnast
	hairStylist
	illustrator
	journalist
	judge
	lawyer
	makeupArtist
	marathonRunner
	mayor
	model
	movieDirector
	movieStar
	musician
	painter
	pastryChef
	pharmacist
	photographer
	pianist
	pilot
	poet
	policeOfficer
	politician
	president
	proGolfer
	proWrestler
	professor
	programmer
	raceCarDriver
	rugbyPlayer
	singer
	soccerPlayer
	stylist
	sumoWrestler
	superhero
	swimmer
	teacher
	tennisPlayer
	tourGuide
	translator
	tycoon
	unknown
	violinist
	volleyballPlayer
	writer
)

var goals = [...]string{
	"Actor",
	"Aesthetician",
	"Archaeologist",
	"Architect",
	"Artist",
	"Astronaut",
	"Baseball Player",
	"Basketball Player",
	"Boxer",
	"CEO",
	"Carpenter",
	"Cartoonist",
	"Chef",
	"Dancer",
	"Dentist",
	"Dermatologist",
	"Designer",
	"Detective",
	"Doctor",
	"Drummer",
	"Esthetician",
	"Explorer",
	"Fashion Designer",
	"Figure Skater",
	"Firefighter",
	"Fisherman",
	"Florist",
	"Forest Ranger",
	"Guitarist",
	"Gymnast",
	"Hair Stylist",
	"Illustrator",
	"Journalist",
	"Judge",
	"Lawyer",
	"Makeup Artist",
	"Marathon Runner",
	"Mayor",
	"Model",
	"Movie Director",
	"Movie Star",
	"Musician",
	"Painter",
	"Pastry Chef",
	"Pharmacist",
	"Photographer",
	"Pianist",
	"Pilot",
	"Poet",
	"Police Officer",
	"Politician",
	"President",
	"Pro Golfer",
	"Pro Wrestler",
	"Professor",
	"Programmer",
	"Race Car Driver",
	"Rugby Player",
	"Singer",
	"Soccer Player",
	"Stylist",
	"Sumo Wrestler",
	"Superhero",
	"Swimmer",
	"Teacher",
	"Tennis Player",
	"Tour Guide",
	"Translator",
	"Tycoon",
	"Unknown",
	"Violinist",
	"Volleyball Player",
	"Writer"}
