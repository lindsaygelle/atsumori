package main

// VillagerGoal is the personal aspiration of an Animal Crossing villager.
type VillagerGoal string

// villagerGoal is a composable struct.
type villagerGoal struct {

	// Goal is the personal aspiration of an Animal Crossing villager.
	Goal *VillagerGoal `json:"goal"`
}

func (v villagerGoal) SetGoal(x *VillagerGoal) {
	v.Goal = x
}

// const villagerActor
const villagerActor VillagerGoal = "Actor"

// const villagerAesthetician
const villagerAesthetician VillagerGoal = "Aesthetician"

// const villagerArchaeologist
const villagerArchaeologist VillagerGoal = "Archaeologist"

// const villagerArchitect
const villagerArchitect VillagerGoal = "Architect"

// const villagerArtist
const villagerArtist VillagerGoal = "Artist"

// const villagerAstronaut
const villagerAstronaut VillagerGoal = "Astronaut"

// const villagerBaseballPlayer
const villagerBaseballPlayer VillagerGoal = "Baseball Player"

// const villagerBasketballPlayer
const villagerBasketballPlayer VillagerGoal = "Basketball Player"

// const villagerBoxer
const villagerBoxer VillagerGoal = "Boxer"

// const villagerCEO
const villagerCEO VillagerGoal = "CEO"

// const villagerCarpenter
const villagerCarpenter VillagerGoal = "Carpenter"

// const villagerCartoonist
const villagerCartoonist VillagerGoal = "Cartoonist"

// const villagerChef
const villagerChef VillagerGoal = "Chef"

// const villagerDancer
const villagerDancer VillagerGoal = "Dancer"

// const villagerDentist
const villagerDentist VillagerGoal = "Dentist"

// const villagerDermatologist
const villagerDermatologist VillagerGoal = "Dermatologist"

// const villagerDesigner
const villagerDesigner VillagerGoal = "Designer"

// const villagerDetective
const villagerDetective VillagerGoal = "Detective"

// const villagerDoctor
const villagerDoctor VillagerGoal = "Doctor"

// const villagerDrummer
const villagerDrummer VillagerGoal = "Drummer"

// const villagerEsthetician
const villagerEsthetician VillagerGoal = "Esthetician"

// const villagerExplorer
const villagerExplorer VillagerGoal = "Explorer"

// const villagerFashionDesigner
const villagerFashionDesigner VillagerGoal = "Fashion Designer"

// const villagerFigureSkater
const villagerFigureSkater VillagerGoal = "Figure Skater"

// const villagerFirefighter
const villagerFirefighter VillagerGoal = "Firefighter"

// const villagerFisherman
const villagerFisherman VillagerGoal = "Fisherman"

// const villagerFlorist
const villagerFlorist VillagerGoal = "Florist"

// const villagerForestRanger
const villagerForestRanger VillagerGoal = "Forest Ranger"

// const villagerGuitarist
const villagerGuitarist VillagerGoal = "Guitarist"

// const villagerGymnast
const villagerGymnast VillagerGoal = "Gymnast"

// const villagerHairStylist
const villagerHairStylist VillagerGoal = "Hair Stylist"

// const villagerIllustrator
const villagerIllustrator VillagerGoal = "Illustrator"

// const villagerJournalist
const villagerJournalist VillagerGoal = "Journalist"

// const villagerJudge
const villagerJudge VillagerGoal = "Judge"

// const villagerLawyer
const villagerLawyer VillagerGoal = "Lawyer"

// const villagerMakeupArtist
const villagerMakeupArtist VillagerGoal = "Makeup Artist"

// const villagerMarathonRunner
const villagerMarathonRunner VillagerGoal = "Marathon Runner"

// const villagerMayor
const villagerMayor VillagerGoal = "Mayor"

// const villagerModel
const villagerModel VillagerGoal = "Model"

// const villagerMovieDirector
const villagerMovieDirector VillagerGoal = "Movie Director"

// const villagerMovieStar
const villagerMovieStar VillagerGoal = "Movie Star"

// const villagerMusician
const villagerMusician VillagerGoal = "Musician"

// const villagerPainter
const villagerPainter VillagerGoal = "Painter"

// const villagerPastryChef
const villagerPastryChef VillagerGoal = "Pastry Chef"

// const villagerPharmacist
const villagerPharmacist VillagerGoal = "Pharmacist"

// const villagerPhotographer
const villagerPhotographer VillagerGoal = "Photographer"

// const villagerPianist
const villagerPianist VillagerGoal = "Pianist"

// const villagerPilot
const villagerPilot VillagerGoal = "Pilot"

// const villagerPoet
const villagerPoet VillagerGoal = "Poet"

// const villagerPoliceOfficer
const villagerPoliceOfficer VillagerGoal = "Police Officer"

// const villagerPolitician
const villagerPolitician VillagerGoal = "Politician"

// const villagerPresident
const villagerPresident VillagerGoal = "President"

// const villagerProGolfer
const villagerProGolfer VillagerGoal = "Pro Golfer"

// const villagerProWrestler
const villagerProWrestler VillagerGoal = "Pro Wrestler"

// const villagerProfessor
const villagerProfessor VillagerGoal = "Professor"

// const villagerProgrammer
const villagerProgrammer VillagerGoal = "Programmer"

// const villagerRaceCarDriver
const villagerRaceCarDriver VillagerGoal = "Racecar driver"

// const villagerRugbyPlayer
const villagerRugbyPlayer VillagerGoal = "Rugby Player"

// const villagerSinger
const villagerSinger VillagerGoal = "Singer"

// const villagerSoccerPlayer
const villagerSoccerPlayer VillagerGoal = "Soccer Player"

// const villagerStylist
const villagerStylist VillagerGoal = "Stylist"

// const villagerSumoWrestler
const villagerSumoWrestler VillagerGoal = "Sumo Wrestler"

// const villagerSuperhero
const villagerSuperhero VillagerGoal = "Superhero"

// const villagerSwimmer
const villagerSwimmer VillagerGoal = "Swimmer"

// const villagerTeacher
const villagerTeacher VillagerGoal = "Teacher"

// const villagerTennisPlayer
const villagerTennisPlayer VillagerGoal = "Tennis Player"

// const villagerTourGuide
const villagerTourGuide VillagerGoal = "Tour Guide"

// const villagerTranslator
const villagerTranslator VillagerGoal = "Translator"

// const villagerTycoon
const villagerTycoon VillagerGoal = "Tycoon"

// const villagerViolinist
const villagerViolinist VillagerGoal = "Violinist"

// const villagerVolleyballPlayer
const villagerVolleyballPlayer VillagerGoal = "Volleyball Player"

// const villagerWriter
const villagerWriter VillagerGoal = "Writer"
