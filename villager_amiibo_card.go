package main

type villagerAmiiboCard struct {
	villagerCard

	villagersBirthday
	villagersClothes
	villagersHome
	villagersHomeRequest
	villagersID
	villagersRoll
	villagersSign
}

type villagersAmiiboCard struct {
	AmiiboCard villagerAmiiboCard
}
