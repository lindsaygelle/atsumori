package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var b, _ = json.Marshal(villager{})
	fmt.Println(string(b))
}
