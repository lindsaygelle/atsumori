package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	var u = "https://animalcrossing.fandom.com/wiki/Villager_list_(New_Horizons)"

	var r, _ = http.Get(u)

	var d, _ = goquery.NewDocumentFromResponse(r)
	var wg sync.WaitGroup
	var roll = map[string]int{}
	var hand = map[string]int{}

	d.Find(".roundy.sortable tr").Each(func(_ int, s *goquery.Selection) {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			var r, _ = http.Get("https://animalcrossing.fandom.com" + u)
			if r.StatusCode != http.StatusOK {
				return
			}
			var d, _ = goquery.NewDocumentFromResponse(r)

			var s = d.Find("table.roundytop")
			fmt.Println(u, s.Length())
			if s.Length() == 0 {
				return
			}
			s.Find("table").Each(func(_ int, s *goquery.Selection) {
				s = s.Find("tr")
				if s.Length() == 0 {
					return
				}
				s.Each(func(_ int, s *goquery.Selection) {
					var k = s.Find("td:nth-child(1)").Text()
					var v = s.Find("td:nth-child(2)").Text()
					fmt.Println(s, k, v)
					if len(k) == 0 || len(v) == 0 {
						panic("!")
					}
					var m map[string]int
					switch strings.ToUpper(k) {
					case "ROLL VALUE":
						m = roll
					case "HAND SIGN":
						m = hand
					}
					if reflect.ValueOf(m).IsZero() {
						return
					}
					if _, ok := m[v]; !ok {
						m[v] = 0
					}
					m[v]++
					fmt.Println(k, v)
				})
			})
		}(s.Find("td a").First().AttrOr("href", ""))
	})
	wg.Wait()
	for _, v := range []map[string]int{roll, hand} {
		var b, _ = json.Marshal(v)
		fmt.Println(string(b))
		fmt.Println("")
		fmt.Println("")
	}
}
