package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s := scanner.Text()

	for _, v := range s {
		if unicode.IsUpper(v) {
			fmt.Printf("%s", strings.ToLower(string(v)))
		} else {
			fmt.Printf("%s", strings.ToUpper(string(v)))
		}
	}
	fmt.Println("")
}
