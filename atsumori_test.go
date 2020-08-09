package atsumori_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/lindsaygelle/atsumori"
)

func Test(t *testing.T) {
	for _, villager := range atsumori.Villagers {
		var b, _ = json.MarshalIndent(villager, "", strings.Repeat(" ", 4))
		fmt.Println(string(b))
	}
}
