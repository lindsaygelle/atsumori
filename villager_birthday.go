package main

import "time"

type villagerBirthday struct {
	Day   uint8
	Month time.Month
}

type villagersBirthday struct {
	Birthday villagerBirthday
}
