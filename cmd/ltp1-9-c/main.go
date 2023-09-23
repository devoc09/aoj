package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var taro, hanako string
	var taroPoint, hanakoPoint int
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &taro, &hanako)

		if result := strings.Compare(taro, hanako); result == 0 {
			taroPoint++
			hanakoPoint++
		} else if result == 1 {
			taroPoint += 3
		} else {
			hanakoPoint += 3
		}
	}
	fmt.Printf("%d %d\n", taroPoint, hanakoPoint)
}
