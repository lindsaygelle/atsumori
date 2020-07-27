package main

type villagerAmiiboCard struct {
	villagerCard

	villagersBirthday
	villagersClothes
	villagersHome
	villagersHomeRequest
	villagersRoll
	villagersSign
}

type villagersAmiiboCard struct {
	AmiiboCard *villagerAmiiboCard
}
