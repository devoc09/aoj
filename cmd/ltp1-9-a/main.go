package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    counter := 0

    var words []string

    for scanner.Scan() {
        line := scanner.Text()

        if line == "END_OF_TEXT" {
            break
        }

        for _, w := range strings.Split(line, " ") {
            words = append(words, strings.ToLower(w))
        }
    }

    for _, w := range words[1:] {
        if w == words[0] {
            counter++
        }
    }
    fmt.Printf("%d\n", counter)
}
