package main

import (
	"fmt"
)

func main() {
	for {
		var cards string
		var m, index int

		fmt.Scan(&cards)
		if cards == "-" {
			break
		}

		fmt.Scan(&m)

		for i := 0; i < m; i++ {
			fmt.Scan(&index)
			cards = cards[index:] + cards[:index]
		}
		fmt.Printf("%s\n", cards)
	}
}
