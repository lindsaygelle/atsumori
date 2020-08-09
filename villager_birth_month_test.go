package atsumori

import (
	"testing"
	"time"
)

func TestVillagerBirthMonth(t *testing.T) {
	if month := (villagersBirthMonth{time.January}).BirthMonth(); month != 1 {
		t.Fatalf("(villagersBirthMonth).BirthMonth() uint8 != 1; %d", month)
	}
}
