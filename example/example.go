package main

import (
	"fmt"

	"github.com/unknowns24/gominatim"
)

func main() {
	geocoder, err := gominatim.NewGominatim(gominatim.DefaultConfig())
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := geocoder.Search(gominatim.SearchParameters{Country: "Germany", City: "Hamburg"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
