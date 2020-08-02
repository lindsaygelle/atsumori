package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var b, _ = json.Marshal(villagerAce)
	fmt.Println(string(b))
}
