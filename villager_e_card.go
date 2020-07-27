package main

type villagerECard struct {
	villagerCard

	Description string
	Letter      string
	Password    string
	Phrase      string
}

type villagersECard struct {
	ECard *villagerECard
}
