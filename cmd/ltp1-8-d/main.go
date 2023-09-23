package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _, _ := reader.ReadLine()
	p, _, _ := reader.ReadLine()

	ts := string(s) + string(s)
	if strings.Contains(ts, string(p)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
