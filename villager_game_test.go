package atsumori

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVillagerGame(t *testing.T) {
	var b, _ = json.Marshal(villagerGameAnimalCrossing)
	fmt.Println(string(b))
	fmt.Println(villagerGameAnimalCrossing)
}
