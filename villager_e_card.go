package main

type villagerECard struct {
	villagerCard
	villagersDescription

	Letter   string
	Password []string
	Phrase   string
}

type villagersECard struct {
	ECard *villagerECard
}
