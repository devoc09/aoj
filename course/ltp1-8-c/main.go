package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	count := make(map[rune]int)
	for {
		data, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		t := strings.ToLower(string(data))

		for _, v := range t {
			for i := 0; i < 26; i++ {
				if v == rune('a'+i) {
					count[rune('a'+i)] += 1
				}
			}
		}
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", rune('a'+i), count[rune('a'+i)])
	}
}
