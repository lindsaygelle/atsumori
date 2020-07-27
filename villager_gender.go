package main

type villagerGender uint8

type villagersGender struct {
	Gender *villagerGender
}

const (
	villagerGenderFemale villagerGender = iota
	villagerGenderMale
)
