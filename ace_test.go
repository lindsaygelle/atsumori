package atsumori_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/lindsaygelle/atsumori"
)

func x(...string) {}

func TestAce(t *testing.T) {
	var b, _ = json.Marshal(&atsumori.Ace)
	fmt.Println(string(b))
}
