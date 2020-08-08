package atsumori_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lindsaygelle/atsumori"
)

func Test(t *testing.T) {
	for _, villager := range atsumori.Villagers {
		var b, _ = json.Marshal(villager)
		fmt.Println(string(b))
	}
}
